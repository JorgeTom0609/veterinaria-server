package medida

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
	r.Get("/medidas", res.getMedidas)
	r.Get("/medidas/<idMedida>", res.getMedidaPorId)
	r.Post("/medidas", res.crearMedida)
	r.Put("/medidas", res.actualizarMedida)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getMedidas(c *routing.Context) error {
	medidas, err := r.service.GetMedidas(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(medidas)
}

func (r resource) crearMedida(c *routing.Context) error {
	var input CreateMedidaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	medida, err := r.service.CrearMedida(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(medida, http.StatusCreated)
}

func (r resource) actualizarMedida(c *routing.Context) error {
	var input UpdateMedidaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	medida, err := r.service.ActualizarMedida(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(medida, http.StatusCreated)
}

func (r resource) getMedidaPorId(c *routing.Context) error {
	idMedida, _ := strconv.Atoi(c.Param("idMedida"))
	medida, err := r.service.GetMedidaPorId(c.Request.Context(), idMedida)
	if err != nil {
		return err
	}
	return c.Write(medida)
}
