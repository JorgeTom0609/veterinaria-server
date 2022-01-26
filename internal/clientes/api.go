package clientes

import (
	"net/http"
	"veterinaria-server/internal/errors"
	"veterinaria-server/pkg/log"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}
	r.Use(authHandler)
	// the following endpoints require a valid JWT
	r.Get("/clientes", res.getClientes)
	r.Post("/clientes", res.crearCliente)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getClientes(c *routing.Context) error {
	clientes, err := r.service.GetClientes(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(clientes)
}

func (r resource) crearCliente(c *routing.Context) error {
	var input CreateClienteRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	cliente, err := r.service.CrearCliente(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(cliente, http.StatusCreated)
}
