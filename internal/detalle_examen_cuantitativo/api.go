package detalle_examen_cuantitativo

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
	r.Get("/detallesExamenCuantitativo", res.getDetallesExamenCuantitativo)
	r.Get("/detallesExamenCuantitativo/<idDetalleExamenCuantitativo>", res.getDetalleExamenCuantitativoPorId)
	r.Post("/detallesExamenCuantitativo", res.crearDetalleExamenCuantitativo)
	r.Put("/detallesExamenCuantitativo", res.actualizarDetalleExamenCuantitativo)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getDetallesExamenCuantitativo(c *routing.Context) error {
	detallesExamenCuantitativo, err := r.service.GetDetallesExamenCuantitativo(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(detallesExamenCuantitativo)
}

func (r resource) crearDetalleExamenCuantitativo(c *routing.Context) error {
	var input CreateDetalleExamenCuantitativoRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleExamenCuantitativo, err := r.service.CrearDetalleExamenCuantitativo(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleExamenCuantitativo, http.StatusCreated)
}

func (r resource) actualizarDetalleExamenCuantitativo(c *routing.Context) error {
	var input UpdateDetalleExamenCuantitativoRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleExamenCuantitativo, err := r.service.ActualizarDetalleExamenCuantitativo(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleExamenCuantitativo, http.StatusCreated)
}

func (r resource) getDetalleExamenCuantitativoPorId(c *routing.Context) error {
	idDetalleExamenCuantitativo, _ := strconv.Atoi(c.Param("idDetalleExamenCuantitativo"))
	detalleExamenCuantitativo, err := r.service.GetDetalleExamenCuantitativoPorId(c.Request.Context(), idDetalleExamenCuantitativo)
	if err != nil {
		return err
	}

	return c.Write(detalleExamenCuantitativo)
}
