package detalle_examen_informativo

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
	r.Get("/detallesExamenInformativo", res.getDetallesExamenInformativo)
	r.Get("/detallesExamenInformativo/<idDetalleExamenInformativo>", res.getDetalleExamenInformativoPorId)
	r.Get("/detallesExamenInformativo/<idTipoDeExamen>", res.getDetallesExamenInformativoPorTipoExamen)
	r.Post("/detallesExamenInformativo", res.crearDetalleExamenInformativo)
	r.Put("/detallesExamenInformativo", res.actualizarDetalleExamenInformativo)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getDetallesExamenInformativo(c *routing.Context) error {
	detallesExamenInformativo, err := r.service.GetDetallesExamenInformativo(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(detallesExamenInformativo)
}

func (r resource) crearDetalleExamenInformativo(c *routing.Context) error {
	var input CreateDetalleExamenInformativoRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleExamenInformativo, err := r.service.CrearDetalleExamenInformativo(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleExamenInformativo, http.StatusCreated)
}

func (r resource) actualizarDetalleExamenInformativo(c *routing.Context) error {
	var input UpdateDetalleExamenInformativoRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleExamenInformativo, err := r.service.ActualizarDetalleExamenInformativo(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleExamenInformativo, http.StatusCreated)
}

func (r resource) getDetalleExamenInformativoPorId(c *routing.Context) error {
	idDetalleExamenInformativo, _ := strconv.Atoi(c.Param("idDetalleExamenInformativo"))
	detalleExamenInformativo, err := r.service.GetDetalleExamenInformativoPorId(c.Request.Context(), idDetalleExamenInformativo)
	if err != nil {
		return err
	}

	return c.Write(detalleExamenInformativo)
}

func (r resource) getDetallesExamenInformativoPorTipoExamen(c *routing.Context) error {
	idTipoDeExamen, _ := strconv.Atoi(c.Param("idTipoDeExamen"))
	detallesExamenCualitativo, err := r.service.GetDetallesExamenInformativoPorTipoExamen(c.Request.Context(), idTipoDeExamen)
	if err != nil {
		return err
	}
	return c.Write(detallesExamenCualitativo)
}
