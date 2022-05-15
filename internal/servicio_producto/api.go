package servicio_producto

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
	r.Get("/servicioProductos", res.getServicioProductos)
	r.Get("/servicioProductos/conDatos", res.getServicioProductosConDatos)
	r.Get("/servicioProductos/<idServicioProducto>", res.getServicioProductoPorId)
	r.Get("/servicioProductos/porServicio/<idServicio>", res.getServicioProductoPorServicio)
	r.Post("/servicioProductos", res.crearServicioProducto)
	r.Put("/servicioProductos", res.actualizarServicioProducto)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getServicioProductos(c *routing.Context) error {
	servicioProductos, err := r.service.GetServicioProductos(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(servicioProductos)
}

func (r resource) getServicioProductosConDatos(c *routing.Context) error {
	servicioProductos, err := r.service.GetServicioProductosConDatos(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(servicioProductos)
}

func (r resource) crearServicioProducto(c *routing.Context) error {
	var input CreateServicioProductoRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	servicioProducto, err := r.service.CrearServicioProducto(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(servicioProducto, http.StatusCreated)
}

func (r resource) actualizarServicioProducto(c *routing.Context) error {
	var input UpdateServicioProductoRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	servicioProducto, err := r.service.ActualizarServicioProducto(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(servicioProducto, http.StatusCreated)
}

func (r resource) getServicioProductoPorId(c *routing.Context) error {
	idServicioProducto, _ := strconv.Atoi(c.Param("idServicioProducto"))
	servicioProducto, err := r.service.GetServicioProductoPorId(c.Request.Context(), idServicioProducto)
	if err != nil {
		return err
	}
	return c.Write(servicioProducto)
}

func (r resource) getServicioProductoPorServicio(c *routing.Context) error {
	idServicio, _ := strconv.Atoi(c.Param("idServicio"))
	servicioProducto, err := r.service.GetServicioProductoPorServicio(c.Request.Context(), idServicio)
	if err != nil {
		return err
	}
	return c.Write(servicioProducto)
}
