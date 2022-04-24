package proveedor_producto

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
	r.Get("/proveedoresProducto", res.getProveedoresProducto)
	r.Get("/proveedoresProducto/<idProveedorProducto>", res.getProveedorProductoPorId)
	r.Post("/proveedoresProducto", res.crearProveedorProducto)
	r.Put("/proveedoresProducto", res.actualizarProveedorProducto)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getProveedoresProducto(c *routing.Context) error {
	proveedoresProducto, err := r.service.GetProveedoresProducto(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(proveedoresProducto)
}

func (r resource) crearProveedorProducto(c *routing.Context) error {
	var input CreateProveedorProductoRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	proveedorProducto, err := r.service.CrearProveedorProducto(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(proveedorProducto, http.StatusCreated)
}

func (r resource) actualizarProveedorProducto(c *routing.Context) error {
	var input UpdateProveedorProductoRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	proveedorProducto, err := r.service.ActualizarProveedorProducto(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(proveedorProducto, http.StatusCreated)
}

func (r resource) getProveedorProductoPorId(c *routing.Context) error {
	idProveedorProducto, _ := strconv.Atoi(c.Param("idProveedorProducto"))
	proveedorProducto, err := r.service.GetProveedorProductoPorId(c.Request.Context(), idProveedorProducto)
	if err != nil {
		return err
	}
	return c.Write(proveedorProducto)
}
