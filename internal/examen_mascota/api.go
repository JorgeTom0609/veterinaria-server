package examen_mascota

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
	r.Get("/examenesMascota", res.getExamenesMascota)
	r.Get("/examenesMascota/<idExamenMascota>", res.getExamenMascotaPorId)
	r.Get("/examenesMascota/examenes/<idMascota>/<estado>", res.getExamenesMascotaPorMascotayEstado)
	r.Post("/examenesMascota", res.crearExamenMascota)
	r.Put("/examenesMascota", res.actualizarExamenMascota)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getExamenesMascota(c *routing.Context) error {
	examenesMascota, err := r.service.GetExamenesMascota(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(examenesMascota)
}

func (r resource) getExamenesMascotaPorMascotayEstado(c *routing.Context) error {
	idExamenMascota, _ := strconv.Atoi(c.Param("idMascota"))
	estado := c.Param("estado")
	examenesMascota, err := r.service.GetExamenesMascotaPorMascotayEstado(c.Request.Context(), idExamenMascota, estado)
	if err != nil {
		return err
	}
	return c.Write(examenesMascota)
}

func (r resource) crearExamenMascota(c *routing.Context) error {
	var input CreateExamenMascotaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	examenesMascota, err := r.service.CrearExamenMascota(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(examenesMascota, http.StatusCreated)
}

func (r resource) actualizarExamenMascota(c *routing.Context) error {
	var input UpdateExamenMascotaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	examenesMascota, err := r.service.ActualizarExamenMascota(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(examenesMascota, http.StatusCreated)
}

func (r resource) getExamenMascotaPorId(c *routing.Context) error {
	idExamenMascota, _ := strconv.Atoi(c.Param("idExamenMascota"))
	examenesMascota, err := r.service.GetExamenMascotaPorId(c.Request.Context(), idExamenMascota)
	if err != nil {
		return err
	}

	return c.Write(examenesMascota)
}
