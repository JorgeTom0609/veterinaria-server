package receta

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
	r.Get("/recetas", res.getRecetas)
	r.Get("/recetas/<idReceta>", res.getRecetaPorId)
	r.Get("/recetas/porConsulta/<idConsulta>", res.getRecetaPorConsulta)
	r.Post("/recetas", res.crearReceta)
	r.Put("/recetas", res.actualizarReceta)
	r.Put("/recetas/toda", res.actualizarRecetaToda)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getRecetas(c *routing.Context) error {
	recetas, err := r.service.GetRecetas(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(recetas)
}

func (r resource) crearReceta(c *routing.Context) error {
	var input CreateRecetaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	receta, err := r.service.CrearReceta(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(receta, http.StatusCreated)
}

func (r resource) actualizarReceta(c *routing.Context) error {
	var input UpdateRecetaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	receta, err := r.service.ActualizarReceta(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(receta, http.StatusCreated)
}

func (r resource) actualizarRecetaToda(c *routing.Context) error {
	var input []UpdateRecetaRequest
	var recetas []Receta = []Receta{}
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	for i := 0; i < len(input); i++ {
		receta, err := r.service.ActualizarReceta(c.Request.Context(), input[i])
		if err != nil {
			return err
		}
		recetas = append(recetas, receta)
	}
	return c.WriteWithStatus(recetas, http.StatusCreated)
}

func (r resource) getRecetaPorId(c *routing.Context) error {
	idReceta, _ := strconv.Atoi(c.Param("idReceta"))
	receta, err := r.service.GetRecetaPorId(c.Request.Context(), idReceta)
	if err != nil {
		return err
	}
	return c.Write(receta)
}

func (r resource) getRecetaPorConsulta(c *routing.Context) error {
	idConsulta, _ := strconv.Atoi(c.Param("idConsulta"))
	receta, err := r.service.GetRecetaPorConsulta(c.Request.Context(), idConsulta)
	if err != nil {
		return err
	}
	return c.Write(receta)
}
