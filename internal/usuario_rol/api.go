package usuario_rol

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
	r.Get("/usuarioRoles", res.getUsuarioRoles)
	r.Get("/usuarioRoles/<idUsuarioRol>", res.getUsuarioRolPorId)
	r.Get("/usuarioRoles/porCedula/<cedula>", res.getUsuarioRolPorCedula)
	r.Post("/usuarioRoles", res.crearUsuarioRol)
	r.Put("/usuarioRoles", res.actualizarUsuarioRol)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getUsuarioRoles(c *routing.Context) error {
	usuarioRoles, err := r.service.GetUsuarioRoles(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(usuarioRoles)
}

func (r resource) crearUsuarioRol(c *routing.Context) error {
	var input CreateUsuarioRolRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	usuarioRol, err := r.service.CrearUsuarioRol(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(usuarioRol, http.StatusCreated)
}

func (r resource) actualizarUsuarioRol(c *routing.Context) error {
	var input UpdateUsuarioRolRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	usuarioRol, err := r.service.ActualizarUsuarioRol(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(usuarioRol, http.StatusCreated)
}

func (r resource) getUsuarioRolPorId(c *routing.Context) error {
	idUsuarioRol, _ := strconv.Atoi(c.Param("idUsuarioRol"))
	usuarioRol, err := r.service.GetUsuarioRolPorId(c.Request.Context(), idUsuarioRol)
	if err != nil {
		return err
	}
	return c.Write(usuarioRol)
}

func (r resource) getUsuarioRolPorCedula(c *routing.Context) error {
	cedula := c.Param("cedula")
	usuarioRol, err := r.service.GetUsuarioRolPorCedula(c.Request.Context(), cedula)
	if err != nil {
		return err
	}
	return c.Write(usuarioRol)
}
