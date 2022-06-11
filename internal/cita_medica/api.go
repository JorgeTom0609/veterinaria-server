package cita_medica

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
	r.Get("/citasMedica", res.getCitasMedica)
	r.Get("/citasMedica/pendientes", res.getCitasMedicaPendientes)
	r.Get("/citasMedica/sinNotificar", res.getCitasMedicaSinNotificar)
	r.Get("/citasMedica/<idCitaMedica>", res.getCitaMedicaPorId)
	r.Post("/citasMedica", res.crearCitaMedica)
	r.Put("/citasMedica", res.actualizarCitaMedica)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getCitasMedica(c *routing.Context) error {
	citasMedica, err := r.service.GetCitasMedica(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(citasMedica)
}

func (r resource) getCitasMedicaPendientes(c *routing.Context) error {
	citasMedica, err := r.service.GetCitasMedicaPendientes(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(citasMedica)
}

func (r resource) getCitasMedicaSinNotificar(c *routing.Context) error {
	citasMedica, err := r.service.GetCitasMedicaSinNotificar(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(citasMedica)
}

func (r resource) crearCitaMedica(c *routing.Context) error {
	var input CreateCitaMedicaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	citaMedica, err := r.service.CrearCitaMedica(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(citaMedica, http.StatusCreated)
}

func (r resource) actualizarCitaMedica(c *routing.Context) error {
	var input UpdateCitaMedicaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	citaMedica, err := r.service.ActualizarCitaMedica(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(citaMedica, http.StatusCreated)
}

func (r resource) getCitaMedicaPorId(c *routing.Context) error {
	idCitaMedica, _ := strconv.Atoi(c.Param("idCitaMedica"))
	citaMedica, err := r.service.GetCitaMedicaPorId(c.Request.Context(), idCitaMedica)
	if err != nil {
		return err
	}
	return c.Write(citaMedica)
}
