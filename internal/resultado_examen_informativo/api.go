package resultado_examen_informativo

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
	r.Get("/resultadoDetalleInformativo/<idResultadoDetalleInformativo>", res.getResultadoDetalleInformativoPorId)
	r.Post("/resultadoDetalleInformativo", res.crearResultadoDetalleInformativo)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) crearResultadoDetalleInformativo(c *routing.Context) error {
	var input CreateResultadoDetalleInformativoRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	resultadoDetalleInformativo, err := r.service.CrearResultadoDetalleInformativo(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(resultadoDetalleInformativo, http.StatusCreated)
}

func (r resource) getResultadoDetalleInformativoPorId(c *routing.Context) error {
	idResultadoDetalleInformativo, _ := strconv.Atoi(c.Param("idResultadoDetalleInformativo"))
	resultadoDetalleInformativo, err := r.service.GetResultadoDetalleInformativoPorId(c.Request.Context(), idResultadoDetalleInformativo)
	if err != nil {
		return err
	}

	return c.Write(resultadoDetalleInformativo)
}
