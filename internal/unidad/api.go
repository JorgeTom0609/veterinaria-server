package unidad

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
	r.Get("/unidades", res.getUnidades)
	r.Get("/unidades/<idUnidad>", res.getUnidadPorId)
	r.Post("/unidades", res.crearUnidad)
	r.Put("/unidades", res.actualizarUnidad)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getUnidades(c *routing.Context) error {
	unidades, err := r.service.GetUnidades(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(unidades)
}

func (r resource) crearUnidad(c *routing.Context) error {
	var input CreateUnidadRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	unidad, err := r.service.CrearUnidad(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(unidad, http.StatusCreated)
}

func (r resource) actualizarUnidad(c *routing.Context) error {
	var input UpdateUnidadRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	unidad, err := r.service.ActualizarUnidad(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(unidad, http.StatusCreated)
}

func (r resource) getUnidadPorId(c *routing.Context) error {
	idUnidad, _ := strconv.Atoi(c.Param("idUnidad"))
	unidad, err := r.service.GetUnidadPorId(c.Request.Context(), idUnidad)
	if err != nil {
		return err
	}
	return c.Write(unidad)
}
