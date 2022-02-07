package tipo_examen

import (
	"net/http"
	"strconv"
	"veterinaria-server/internal/errors"
	"veterinaria-server/pkg/log"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}
	r.Use(authHandler)
	r.Get("/tipo_examen", res.getTipoExamenes)
	r.Get("/tipo_examen/<idTipoExamen>", res.getTipoExamenPorId)
	r.Post("/tipo_examen", res.crearTipoExamen)
	r.Put("/tipo_examen", res.actualizarTipoExamen)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getTipoExamenes(c *routing.Context) error {
	tipoExamenes, err := r.service.GetTipoExamenes(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(tipoExamenes)
}

func (r resource) crearTipoExamen(c *routing.Context) error {
	var input CreateTipoExamenRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	tipoExamen, err := r.service.CrearTipoExamen(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(tipoExamen, http.StatusCreated)
}

func (r resource) actualizarTipoExamen(c *routing.Context) error {
	var input UpdateTipoExamenRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	tipoExamen, err := r.service.ActualizarTipoExamen(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(tipoExamen, http.StatusCreated)
}

func (r resource) getTipoExamenPorId(c *routing.Context) error {
	idTipoExamen, _ := strconv.Atoi(c.Param("idTipoExamen"))
	tipoExamen, err := r.service.GetTipoExamenPorId(c.Request.Context(), idTipoExamen)
	if err != nil {
		return err
	}

	return c.Write(tipoExamen)
}
