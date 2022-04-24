package lote

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
	r.Get("/lotes", res.getLotes)
	r.Get("/lotes/<idLote>", res.getLotePorId)
	r.Post("/lotes", res.crearLote)
	r.Put("/lotes", res.actualizarLote)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getLotes(c *routing.Context) error {
	lotes, err := r.service.GetLotes(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(lotes)
}

func (r resource) crearLote(c *routing.Context) error {
	var input CreateLoteRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	lote, err := r.service.CrearLote(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(lote, http.StatusCreated)
}

func (r resource) actualizarLote(c *routing.Context) error {
	var input UpdateLoteRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	lote, err := r.service.ActualizarLote(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(lote, http.StatusCreated)
}

func (r resource) getLotePorId(c *routing.Context) error {
	idLote, _ := strconv.Atoi(c.Param("idLote"))
	lote, err := r.service.GetLotePorId(c.Request.Context(), idLote)
	if err != nil {
		return err
	}
	return c.Write(lote)
}
