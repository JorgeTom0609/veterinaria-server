package resultado_examen_cuantitativo

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
	r.Get("/resultadoDetalleCuantitativo/<idResultadoDetalleCuantitativo>", res.getResultadoDetalleCuantitativoPorId)
	r.Post("/resultadoDetalleCuantitativo", res.crearResultadoDetalleCuantitativo)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) crearResultadoDetalleCuantitativo(c *routing.Context) error {
	var input CreateResultadoDetalleCuantitativoRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	resultadoDetalleCuantitativo, err := r.service.CrearResultadoDetalleCuantitativo(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(resultadoDetalleCuantitativo, http.StatusCreated)
}

func (r resource) getResultadoDetalleCuantitativoPorId(c *routing.Context) error {
	idResultadoDetalleCuantitativo, _ := strconv.Atoi(c.Param("idResultadoDetalleCuantitativo"))
	resultadoDetalleCuantitativo, err := r.service.GetResultadoDetalleCuantitativoPorId(c.Request.Context(), idResultadoDetalleCuantitativo)
	if err != nil {
		return err
	}

	return c.Write(resultadoDetalleCuantitativo)
}
