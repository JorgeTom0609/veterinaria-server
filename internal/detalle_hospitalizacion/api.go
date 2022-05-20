package detalle_hospitalizacion

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
	r.Get("/detallesHospitalizacion", res.getDetallesHospitalizacion)
	r.Get("/detallesHospitalizacion/<idDetalleHospitalizacion>", res.getDetalleHospitalizacionPorId)
	r.Post("/detallesHospitalizacion", res.crearDetalleHospitalizacion)
	r.Put("/detallesHospitalizacion", res.actualizarDetalleHospitalizacion)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getDetallesHospitalizacion(c *routing.Context) error {
	detallesHospitalizacion, err := r.service.GetDetallesHospitalizacion(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(detallesHospitalizacion)
}

func (r resource) crearDetalleHospitalizacion(c *routing.Context) error {
	var input CreateDetalleHospitalizacionRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleHospitalizacion, err := r.service.CrearDetalleHospitalizacion(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleHospitalizacion, http.StatusCreated)
}

func (r resource) actualizarDetalleHospitalizacion(c *routing.Context) error {
	var input UpdateDetalleHospitalizacionRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleHospitalizacion, err := r.service.ActualizarDetalleHospitalizacion(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleHospitalizacion, http.StatusCreated)
}

func (r resource) getDetalleHospitalizacionPorId(c *routing.Context) error {
	idDetalleHospitalizacion, _ := strconv.Atoi(c.Param("idDetalleHospitalizacion"))
	detalleHospitalizacion, err := r.service.GetDetalleHospitalizacionPorId(c.Request.Context(), idDetalleHospitalizacion)
	if err != nil {
		return err
	}
	return c.Write(detalleHospitalizacion)
}
