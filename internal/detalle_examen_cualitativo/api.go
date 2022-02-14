package detalle_examen_cualitativo

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
	r.Get("/detallesExamenCualitativo", res.getDetallesExamenCualitativo)
	r.Get("/detallesExamenCualitativo/<idDetalleExamenCualitativo>", res.getDetalleExamenCualitativoPorId)
	r.Get("/detallesExamenCualitativo/tipoExamen/<idTipoDeExamen>", res.getDetallesExamenCualitativoPorTipoExamen)
	r.Post("/detallesExamenCualitativo", res.crearDetalleExamenCualitativo)
	r.Put("/detallesExamenCualitativo", res.actualizarDetalleExamenCualitativo)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getDetallesExamenCualitativo(c *routing.Context) error {
	detallesExamenCualitativo, err := r.service.GetDetallesExamenCualitativo(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(detallesExamenCualitativo)
}

func (r resource) crearDetalleExamenCualitativo(c *routing.Context) error {
	var input CreateDetalleExamenCualitativoRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleExamenCualitativo, err := r.service.CrearDetalleExamenCualitativo(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleExamenCualitativo, http.StatusCreated)
}

func (r resource) actualizarDetalleExamenCualitativo(c *routing.Context) error {
	var input UpdateDetalleExamenCualitativoRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleExamenCualitativo, err := r.service.ActualizarDetalleExamenCualitativo(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleExamenCualitativo, http.StatusCreated)
}

func (r resource) getDetalleExamenCualitativoPorId(c *routing.Context) error {
	idDetalleExamenCualitativo, _ := strconv.Atoi(c.Param("idDetalleExamenCualitativo"))
	detalleExamenCualitativo, err := r.service.GetDetalleExamenCualitativoPorId(c.Request.Context(), idDetalleExamenCualitativo)
	if err != nil {
		return err
	}
	return c.Write(detalleExamenCualitativo)
}

func (r resource) getDetallesExamenCualitativoPorTipoExamen(c *routing.Context) error {
	idTipoDeExamen, _ := strconv.Atoi(c.Param("idTipoDeExamen"))
	detallesExamenCualitativo, err := r.service.GetDetallesExamenCualitativoPorTipoExamen(c.Request.Context(), idTipoDeExamen)
	if err != nil {
		return err
	}
	return c.Write(detallesExamenCualitativo)
}
