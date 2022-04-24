package proveedor

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
	r.Get("/proveedores", res.getProveedores)
	r.Get("/proveedores/<idProveedor>", res.getProveedorPorId)
	r.Post("/proveedores", res.crearProveedor)
	r.Put("/proveedores", res.actualizarProveedor)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getProveedores(c *routing.Context) error {
	proveedores, err := r.service.GetProveedores(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(proveedores)
}

func (r resource) crearProveedor(c *routing.Context) error {
	var input CreateProveedorRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	proveedor, err := r.service.CrearProveedor(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(proveedor, http.StatusCreated)
}

func (r resource) actualizarProveedor(c *routing.Context) error {
	var input UpdateProveedorRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	proveedor, err := r.service.ActualizarProveedor(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(proveedor, http.StatusCreated)
}

func (r resource) getProveedorPorId(c *routing.Context) error {
	idProveedor, _ := strconv.Atoi(c.Param("idProveedor"))
	proveedor, err := r.service.GetProveedorPorId(c.Request.Context(), idProveedor)
	if err != nil {
		return err
	}
	return c.Write(proveedor)
}
