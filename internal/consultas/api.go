package consultas

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
	r.Get("/consultas", res.getConsultas)
	r.Get("/consultas/porMesYAnio", res.getConsultaPorMesYAnio)
	r.Get("/consultas/porMascota/<idMascota>", res.getConsultaPorMascota)
	r.Get("/consultas/<idConsulta>", res.getConsultaPorId)
	r.Get("/consultas/activa/<idUsuario>", res.getConsultaActiva)
	r.Post("/consultas", res.crearConsulta)
	r.Put("/consultas", res.actualizarConsulta)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getConsultas(c *routing.Context) error {
	consultas, err := r.service.GetConsultas(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(consultas)
}

func (r resource) crearConsulta(c *routing.Context) error {
	var input CreateConsultaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	consulta, err := r.service.CrearConsulta(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(consulta, http.StatusCreated)
}

func (r resource) actualizarConsulta(c *routing.Context) error {
	var input UpdateConsultaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	consulta, err := r.service.ActualizarConsulta(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(consulta, http.StatusCreated)
}

func (r resource) getConsultaPorId(c *routing.Context) error {
	idConsulta, _ := strconv.Atoi(c.Param("idConsulta"))
	consulta, err := r.service.GetConsultaPorId(c.Request.Context(), idConsulta)
	if err != nil {
		return err
	}
	return c.Write(consulta)
}

func (r resource) getConsultaActiva(c *routing.Context) error {
	idUsuario, _ := strconv.Atoi(c.Param("idUsuario"))
	consulta, err := r.service.GetConsultaActiva(c.Request.Context(), idUsuario)
	if err != nil {
		return err
	}
	return c.Write(consulta)
}

func (r resource) getConsultaPorMesYAnio(c *routing.Context) error {
	consulta, err := r.service.GetConsultaPorMesYAnio(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(consulta)
}

func (r resource) getConsultaPorMascota(c *routing.Context) error {
	idMascota, _ := strconv.Atoi(c.Param("idMascota"))
	consulta, err := r.service.GetConsultaPorMascota(c.Request.Context(), idMascota)
	if err != nil {
		return err
	}
	return c.Write(consulta)
}
