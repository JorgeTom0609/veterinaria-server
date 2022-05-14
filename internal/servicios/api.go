package servicios

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
	r.Get("/servicios", res.getServicios)
	r.Get("/servicios/conNumProductosUso", res.getServiciosConProductos)
	r.Get("/servicios/<idServicio>", res.getServicioPorId)
	r.Post("/servicios", res.crearServicio)
	r.Put("/servicios", res.actualizarServicio)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getServicios(c *routing.Context) error {
	servicios, err := r.service.GetServicios(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(servicios)
}

func (r resource) getServiciosConProductos(c *routing.Context) error {
	servicios, err := r.service.GetServiciosConProductos(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(servicios)
}

func (r resource) crearServicio(c *routing.Context) error {
	var input CreateServicioRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	servicio, err := r.service.CrearServicio(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(servicio, http.StatusCreated)
}

func (r resource) actualizarServicio(c *routing.Context) error {
	var input UpdateServicioRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	servicio, err := r.service.ActualizarServicio(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(servicio, http.StatusCreated)
}

func (r resource) getServicioPorId(c *routing.Context) error {
	idServicio, _ := strconv.Atoi(c.Param("idServicio"))
	servicio, err := r.service.GetServicioPorId(c.Request.Context(), idServicio)
	if err != nil {
		return err
	}
	return c.Write(servicio)
}
