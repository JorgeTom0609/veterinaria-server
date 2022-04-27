package detalle_compra

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
	r.Get("/detallesCompra", res.getDetallesCompra)
	r.Get("/detallesCompra/<idDetalleCompra>", res.getDetalleCompraPorId)
	r.Get("/detallesCompra/porCompra/<idCompra>", res.getDetalleCompraPorIdCompra)
	r.Post("/detallesCompra", res.crearDetalleCompra)
	r.Put("/detallesCompra", res.actualizarDetalleCompra)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getDetallesCompra(c *routing.Context) error {
	detallesCompra, err := r.service.GetDetallesCompra(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(detallesCompra)
}

func (r resource) crearDetalleCompra(c *routing.Context) error {
	var input CreateDetalleCompraRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleCompra, err := r.service.CrearDetalleCompra(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleCompra, http.StatusCreated)
}

func (r resource) actualizarDetalleCompra(c *routing.Context) error {
	var input UpdateDetalleCompraRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleCompra, err := r.service.ActualizarDetalleCompra(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleCompra, http.StatusCreated)
}

func (r resource) getDetalleCompraPorId(c *routing.Context) error {
	idDetalleCompra, _ := strconv.Atoi(c.Param("idDetalleCompra"))
	detalleCompra, err := r.service.GetDetalleCompraPorId(c.Request.Context(), idDetalleCompra)
	if err != nil {
		return err
	}

	return c.Write(detalleCompra)
}

func (r resource) getDetalleCompraPorIdCompra(c *routing.Context) error {
	idFactura, _ := strconv.Atoi(c.Param("idCompra"))
	detallesFactura, err := r.service.GetDetalleCompraPorIdCompra(c.Request.Context(), idFactura)
	if err != nil {
		return err
	}

	return c.Write(detallesFactura)
}
