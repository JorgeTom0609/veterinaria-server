package factura

import (
	"net/http"
	"strconv"
	"veterinaria-server/internal/errors"
	"veterinaria-server/pkg/log"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}
	r.Use(authHandler)
	// the following endpoints require a valid JWT
	r.Get("/facturas", res.getFacturas)
	r.Get("/facturas/<idFactura>", res.getFacturaPorId)
	r.Post("/facturas", res.crearFactura)
	r.Put("/facturas", res.actualizarFactura)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getFacturas(c *routing.Context) error {
	facturas, err := r.service.GetFacturas(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(facturas)
}

func (r resource) crearFactura(c *routing.Context) error {
	var input CreateFacturaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	factura, err := r.service.CrearFactura(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(factura, http.StatusCreated)
}

func (r resource) actualizarFactura(c *routing.Context) error {
	var input UpdateFacturaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	factura, err := r.service.ActualizarFactura(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(factura, http.StatusCreated)
}

func (r resource) getFacturaPorId(c *routing.Context) error {
	idFactura, _ := strconv.Atoi(c.Param("idFactura"))
	factura, err := r.service.GetFacturaPorId(c.Request.Context(), idFactura)
	if err != nil {
		return err
	}

	return c.Write(factura)
}
