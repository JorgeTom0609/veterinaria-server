package detalle_factura

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
	r.Get("/detallesFactura", res.getDetallesFactura)
	r.Get("/detallesFactura/<idDetalleFactura>", res.getDetalleFacturaPorId)
	r.Post("/detallesFactura", res.crearDetalleFactura)
	r.Put("/detallesFactura", res.actualizarDetalleFactura)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getDetallesFactura(c *routing.Context) error {
	detallesFactura, err := r.service.GetDetallesFactura(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(detallesFactura)
}

func (r resource) crearDetalleFactura(c *routing.Context) error {
	var input CreateDetalleFacturaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleFactura, err := r.service.CrearDetalleFactura(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleFactura, http.StatusCreated)
}

func (r resource) actualizarDetalleFactura(c *routing.Context) error {
	var input UpdateDetalleFacturaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleFactura, err := r.service.ActualizarDetalleFactura(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleFactura, http.StatusCreated)
}

func (r resource) getDetalleFacturaPorId(c *routing.Context) error {
	idDetalleFactura, _ := strconv.Atoi(c.Param("idDetalleFactura"))
	detalleFactura, err := r.service.GetDetalleFacturaPorId(c.Request.Context(), idDetalleFactura)
	if err != nil {
		return err
	}

	return c.Write(detalleFactura)
}
