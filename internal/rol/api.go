package rol

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
	r.Get("/roles", res.getRoles)
	r.Get("/roles/<idRol>", res.getRolPorId)
	r.Post("/roles", res.crearRol)
	r.Put("/roles", res.actualizarRol)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getRoles(c *routing.Context) error {
	roles, err := r.service.GetRoles(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(roles)
}

func (r resource) crearRol(c *routing.Context) error {
	var input CreateRolRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	rol, err := r.service.CrearRol(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(rol, http.StatusCreated)
}

func (r resource) actualizarRol(c *routing.Context) error {
	var input UpdateRolRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	rol, err := r.service.ActualizarRol(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(rol, http.StatusCreated)
}

func (r resource) getRolPorId(c *routing.Context) error {
	idRol, _ := strconv.Atoi(c.Param("idRol"))
	rol, err := r.service.GetRolPorId(c.Request.Context(), idRol)
	if err != nil {
		return err
	}
	return c.Write(rol)
}
