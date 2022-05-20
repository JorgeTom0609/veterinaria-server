package hospitalizacion

import (
	"net/http"
	"strconv"
	"time"
	"veterinaria-server/internal/detalle_hospitalizacion"
	"veterinaria-server/internal/errors"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger, db *dbcontext.DB) {
	res := resource{service, logger, db}
	r.Use(authHandler)
	// the following endpoints require a valid JWT
	r.Get("/hospitalizaciones", res.getHospitalizaciones)
	r.Get("/hospitalizaciones/activas", res.getHospitalizacionesActivas)
	r.Get("/hospitalizaciones/<idHospitalizacion>", res.getHospitalizacionPorId)
	r.Post("/hospitalizaciones", res.crearHospitalizacion)
	r.Put("/hospitalizaciones", res.actualizarHospitalizacion)
}

type resource struct {
	service Service
	logger  log.Logger
	db      *dbcontext.DB
}

func (r resource) getHospitalizaciones(c *routing.Context) error {
	hospitalizaciones, err := r.service.GetHospitalizaciones(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(hospitalizaciones)
}

func (r resource) getHospitalizacionesActivas(c *routing.Context) error {
	hospitalizaciones, err := r.service.GetHospitalizacionesActivas(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(hospitalizaciones)
}

func (r resource) crearHospitalizacion(c *routing.Context) error {
	var input CreateHospitalizacionRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	hospitalizacion, err := r.service.CrearHospitalizacion(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(hospitalizacion, http.StatusCreated)
}

func (r resource) actualizarHospitalizacion(c *routing.Context) error {
	var input UpdateHospitalizacionRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	input.FechaIngreso = time.Now()
	hospitalizacion, err := r.service.ActualizarHospitalizacion(c.Request.Context(), input)
	if err != nil {
		return err
	}
	if input.IdHospitalizacion == 0 {
		s := detalle_hospitalizacion.NewService(detalle_hospitalizacion.NewRepository(r.db, r.logger), r.logger)
		var detalleHospitalizacion detalle_hospitalizacion.UpdateDetalleHospitalizacionRequest = detalle_hospitalizacion.UpdateDetalleHospitalizacionRequest{
			IdDetalleHospitalizacion: 0,
			IdHospitalizacion:        hospitalizacion.IdHospitalizacion,
			IdUsuario:                input.IdUsuario,
			Descripcion:              "Ingreso al área de hospitalización",
			Fecha:                    time.Now(),
		}
		_, err = s.ActualizarDetalleHospitalizacion(c.Request.Context(), detalleHospitalizacion)
		if err != nil {
			return err
		}
	}
	return c.WriteWithStatus(hospitalizacion, http.StatusCreated)
}

func (r resource) getHospitalizacionPorId(c *routing.Context) error {
	idHospitalizacion, _ := strconv.Atoi(c.Param("idHospitalizacion"))
	hospitalizacion, err := r.service.GetHospitalizacionPorId(c.Request.Context(), idHospitalizacion)
	if err != nil {
		return err
	}
	return c.Write(hospitalizacion)
}
