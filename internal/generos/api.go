package generos

import (
	"strconv"
	"veterinaria-server/pkg/log"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}
	r.Use(authHandler)
	// the following endpoints require a valid JWT
	r.Get("/generos", res.getGeneros)
	r.Get("/generos/<idGenero>", res.getGeneroPorID)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getGeneros(c *routing.Context) error {
	generos, err := r.service.GetGeneros(c.Request.Context())
	if err != nil {
		return err
	}

	return c.Write(generos)
}

func (r resource) getGeneroPorID(c *routing.Context) error {
	idGenero, _ := strconv.Atoi(c.Param("idGenero"))
	genero, err := r.service.GetGeneroPorID(c.Request.Context(), idGenero)
	if err != nil {
		return err
	}
	return c.Write(genero)
}
