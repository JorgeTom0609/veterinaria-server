package stock_individual

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
	r.Get("/stocksIndividual", res.getStocksIndividual)
	r.Get("/stocksIndividual/<idStockIndividual>", res.getStockIndividualPorId)
	r.Post("/stocksIndividual", res.crearStockIndividual)
	r.Put("/stocksIndividual", res.actualizarStockIndividual)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getStocksIndividual(c *routing.Context) error {
	stocksIndividual, err := r.service.GetStocksIndividual(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(stocksIndividual)
}

func (r resource) crearStockIndividual(c *routing.Context) error {
	var input CreateStockIndividualRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	stockIndividual, err := r.service.CrearStockIndividual(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(stockIndividual, http.StatusCreated)
}

func (r resource) actualizarStockIndividual(c *routing.Context) error {
	var input UpdateStockIndividualRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	stockIndividual, err := r.service.ActualizarStockIndividual(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(stockIndividual, http.StatusCreated)
}

func (r resource) getStockIndividualPorId(c *routing.Context) error {
	idStockIndividual, _ := strconv.Atoi(c.Param("idStockIndividual"))
	stockIndividual, err := r.service.GetStockIndividualPorId(c.Request.Context(), idStockIndividual)
	if err != nil {
		return err
	}
	return c.Write(stockIndividual)
}
