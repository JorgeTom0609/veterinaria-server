package detalle_uso_servicio

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
	r.Get("/detallesUsoServicio", res.getDetallesUsoServicio)
	r.Get("/detallesUsoServicio/<idDetalleUsoServicio>", res.getDetalleUsoServicioPorId)
	r.Post("/detallesUsoServicio", res.crearDetalleUsoServicio)
	r.Put("/detallesUsoServicio", res.actualizarDetalleUsoServicio)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getDetallesUsoServicio(c *routing.Context) error {
	detallesUsoServicio, err := r.service.GetDetallesUsoServicio(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(detallesUsoServicio)
}

func (r resource) crearDetalleUsoServicio(c *routing.Context) error {
	var input CreateDetalleUsoServicioRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleUsoServicio, err := r.service.CrearDetalleUsoServicio(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleUsoServicio, http.StatusCreated)
}

func (r resource) actualizarDetalleUsoServicio(c *routing.Context) error {
	var input UpdateDetalleUsoServicioRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleUsoServicio, err := r.service.ActualizarDetalleUsoServicio(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleUsoServicio, http.StatusCreated)
}

func (r resource) getDetalleUsoServicioPorId(c *routing.Context) error {
	idDetalleUsoServicio, _ := strconv.Atoi(c.Param("idDetalleUsoServicio"))
	detalleUsoServicio, err := r.service.GetDetalleUsoServicioPorId(c.Request.Context(), idDetalleUsoServicio)
	if err != nil {
		return err
	}
	return c.Write(detalleUsoServicio)
}
