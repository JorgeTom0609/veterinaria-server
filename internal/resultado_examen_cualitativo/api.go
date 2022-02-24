package resultado_examen_cualitativo

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
	r.Get("/resultadoDetalleCualitativo/<idResultadoDetalleCualitativo>", res.getResultadoDetalleCualitativoPorId)
	r.Post("/resultadoDetalleCualitativo", res.crearResultadoDetalleCualitativo)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) crearResultadoDetalleCualitativo(c *routing.Context) error {
	var input CreateResultadoDetalleCualitativoRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	resultadoDetalleCualitativo, err := r.service.CrearResultadoDetalleCualitativo(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(resultadoDetalleCualitativo, http.StatusCreated)
}

func (r resource) getResultadoDetalleCualitativoPorId(c *routing.Context) error {
	idResultadoDetalleCualitativo, _ := strconv.Atoi(c.Param("idResultadoDetalleCualitativo"))
	resultadoDetalleCualitativo, err := r.service.GetResultadoDetalleCualitativoPorId(c.Request.Context(), idResultadoDetalleCualitativo)
	if err != nil {
		return err
	}

	return c.Write(resultadoDetalleCualitativo)
}
