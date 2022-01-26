package accesos

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
	r.Get("/accesos/<idUsuario>", res.getAccesosPorIdUsuario)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getAccesosPorIdUsuario(c *routing.Context) error {
	id, _ := strconv.Atoi(c.Param("idUsuario"))
	accesos, idRol, err := r.service.GetAccesosPorIdUsuario(c.Request.Context(), id)
	if err != nil {
		return err
	}

	return c.Write(struct {
		IdRol   int       `json:"id_rol"`
		Accesos []Accesos `json:"accesos"`
	}{IdRol: idRol, Accesos: accesos})
}
