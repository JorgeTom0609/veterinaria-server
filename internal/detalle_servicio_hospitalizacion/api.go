package detalle_servicio_hospitalizacion

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
	r.Get("/detallesServicioHospitalizacion", res.getDetallesServicioHospitalizacion)
	r.Get("/detallesServicioHospitalizacion/<idDetalleServicioHospitalizacion>", res.getDetalleServicioHospitalizacionPorId)
	r.Post("/detallesServicioHospitalizacion", res.crearDetalleServicioHospitalizacion)
	r.Put("/detallesServicioHospitalizacion", res.actualizarDetalleServicioHospitalizacion)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getDetallesServicioHospitalizacion(c *routing.Context) error {
	detallesServicioHospitalizacion, err := r.service.GetDetallesServicioHospitalizacion(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(detallesServicioHospitalizacion)
}

func (r resource) crearDetalleServicioHospitalizacion(c *routing.Context) error {
	var input CreateDetalleServicioHospitalizacionRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleServicioHospitalizacion, err := r.service.CrearDetalleServicioHospitalizacion(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleServicioHospitalizacion, http.StatusCreated)
}

func (r resource) actualizarDetalleServicioHospitalizacion(c *routing.Context) error {
	var input UpdateDetalleServicioHospitalizacionRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleServicioHospitalizacion, err := r.service.ActualizarDetalleServicioHospitalizacion(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleServicioHospitalizacion, http.StatusCreated)
}

func (r resource) getDetalleServicioHospitalizacionPorId(c *routing.Context) error {
	idDetalleServicioHospitalizacion, _ := strconv.Atoi(c.Param("idDetalleServicioHospitalizacion"))
	detalleServicioHospitalizacion, err := r.service.GetDetalleServicioHospitalizacionPorId(c.Request.Context(), idDetalleServicioHospitalizacion)
	if err != nil {
		return err
	}
	return c.Write(detalleServicioHospitalizacion)
}
