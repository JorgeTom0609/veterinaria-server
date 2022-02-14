package mascotas

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
	r.Get("/mascotas/cliente/<idCliente>", res.getMascotasPorCliente)
	r.Get("/mascotas/<idMascota>", res.getMascotaPorId)
	r.Post("/mascotas", res.crearMascota)
	r.Put("/mascotas/grupo", res.actualizarMascotaPorGrupo)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getMascotasPorCliente(c *routing.Context) error {
	idCliente, _ := strconv.Atoi(c.Param("idCliente"))
	mascotas, err := r.service.GetMascotasPorCliente(c.Request.Context(), idCliente)
	if err != nil {
		return err
	}
	return c.Write(mascotas)
}

func (r resource) crearMascota(c *routing.Context) error {
	var input CreateMascotaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	mascota, err := r.service.CrearMascota(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(mascota, http.StatusCreated)
}

func (r resource) actualizarMascotaPorGrupo(c *routing.Context) error {
	var input CreateMascotaPorGrupoRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	mascotas, err := r.service.ActualizarMascotaPorGrupo(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(mascotas, http.StatusCreated)
}

func (r resource) getMascotaPorId(c *routing.Context) error {
	idMascota, _ := strconv.Atoi(c.Param("idMascota"))
	mascota, err := r.service.GetMascotaPorId(c.Request.Context(), idMascota)
	if err != nil {
		return err
	}

	return c.Write(mascota)
}
