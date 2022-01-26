package especies

import (
	"veterinaria-server/pkg/log"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}
	r.Use(authHandler)
	// the following endpoints require a valid JWT
	r.Get("/especies", res.getEspecies)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getEspecies(c *routing.Context) error {
	especies, err := r.service.GetEspecies(c.Request.Context())
	if err != nil {
		return err
	}

	return c.Write(especies)
}
