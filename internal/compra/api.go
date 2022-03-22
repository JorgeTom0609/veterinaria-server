package compra

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
	r.Get("/compras", res.getCompras)
	r.Get("/compras/<idCompra>", res.getCompraPorId)
	r.Post("/compras", res.crearCompra)
	r.Put("/compras", res.actualizarCompra)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getCompras(c *routing.Context) error {
	compras, err := r.service.GetCompras(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(compras)
}

func (r resource) crearCompra(c *routing.Context) error {
	var input CreateCompraRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	compra, err := r.service.CrearCompra(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(compra, http.StatusCreated)
}

func (r resource) actualizarCompra(c *routing.Context) error {
	var input UpdateCompraRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	compra, err := r.service.ActualizarCompra(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(compra, http.StatusCreated)
}

func (r resource) getCompraPorId(c *routing.Context) error {
	idCompra, _ := strconv.Atoi(c.Param("idCompra"))
	compra, err := r.service.GetCompraPorId(c.Request.Context(), idCompra)
	if err != nil {
		return err
	}

	return c.Write(compra)
}
