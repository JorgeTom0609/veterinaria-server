package usuarios

import (
	"net/http"
	"strconv"
	"veterinaria-server/internal/errors"
	"veterinaria-server/pkg/log"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"golang.org/x/crypto/bcrypt"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}
	r.Use(authHandler)
	// the following endpoints require a valid JWT
	r.Get("/users", res.getUsers)
	r.Get("/users/<idUser>", res.getUserPorId)
	r.Post("/users", res.crearUser)
	r.Put("/users", res.actualizarUser)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getUsers(c *routing.Context) error {
	users, err := r.service.GetUsers(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(users)
}

func (r resource) crearUser(c *routing.Context) error {
	var input CreateUserRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	user, err := r.service.CrearUser(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(user, http.StatusCreated)
}

func (r resource) actualizarUser(c *routing.Context) error {
	var input UpdateUserRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	if input.CambioClave == 1 {
		// para encriptar la pass
		hash, _ := bcrypt.GenerateFromPassword([]byte(input.Clave), bcrypt.MinCost)
		input.Clave = string(hash)
	}

	user, err := r.service.ActualizarUser(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(user, http.StatusCreated)
}

func (r resource) getUserPorId(c *routing.Context) error {
	idUser, _ := strconv.Atoi(c.Param("idUser"))
	user, err := r.service.GetUserPorId(c.Request.Context(), idUser)
	if err != nil {
		return err
	}
	return c.Write(user)
}
