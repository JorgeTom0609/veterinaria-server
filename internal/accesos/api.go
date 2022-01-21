package accesos

import (
	"veterinaria-server/pkg/log"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}
	r.Get("/acceso/<idUsuario>", res.getAccesosPorIdUsuario)
	r.Use(authHandler)
	// the following endpoints require a valid JWT
	r.Get("/accesos/<idUsuario>", res.getAccesosPorIdUsuario)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getAccesosPorIdUsuario(c *routing.Context) error {
	//c.Param("id")
	accesos, idRol, err := r.service.GetAccesosPorIdUsuario(c.Request.Context(), 1)
	if err != nil {
		return err
	}

	return c.Write(struct {
		IdRol   int       `json:"id_rol"`
		Accesos []Accesos `json:"accesos"`
	}{IdRol: idRol, Accesos: accesos})
}
