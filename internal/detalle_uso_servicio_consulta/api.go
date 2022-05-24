package detalle_uso_servicio_consulta

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
	r.Get("/detallesUsoServicioConsulta", res.getDetallesUsoServicioConsulta)
	r.Get("/detallesUsoServicioConsulta/<idDetalleUsoServicioConsulta>", res.getDetalleUsoServicioConsultaPorId)
	r.Post("/detallesUsoServicioConsulta", res.crearDetalleUsoServicioConsulta)
	r.Put("/detallesUsoServicioConsulta", res.actualizarDetalleUsoServicioConsulta)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getDetallesUsoServicioConsulta(c *routing.Context) error {
	detallesUsoServicioConsulta, err := r.service.GetDetallesUsoServicioConsulta(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(detallesUsoServicioConsulta)
}

func (r resource) crearDetalleUsoServicioConsulta(c *routing.Context) error {
	var input CreateDetalleUsoServicioConsultaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleUsoServicioConsulta, err := r.service.CrearDetalleUsoServicioConsulta(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleUsoServicioConsulta, http.StatusCreated)
}

func (r resource) actualizarDetalleUsoServicioConsulta(c *routing.Context) error {
	var input UpdateDetalleUsoServicioConsultaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleUsoServicioConsulta, err := r.service.ActualizarDetalleUsoServicioConsulta(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleUsoServicioConsulta, http.StatusCreated)
}

func (r resource) getDetalleUsoServicioConsultaPorId(c *routing.Context) error {
	idDetalleUsoServicioConsulta, _ := strconv.Atoi(c.Param("idDetalleUsoServicioConsulta"))
	detalleUsoServicioConsulta, err := r.service.GetDetalleUsoServicioConsultaPorId(c.Request.Context(), idDetalleUsoServicioConsulta)
	if err != nil {
		return err
	}
	return c.Write(detalleUsoServicioConsulta)
}
