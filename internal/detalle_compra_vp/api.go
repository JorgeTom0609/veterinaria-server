package detalle_compra_vp

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
	r.Get("/detallesCompraVP", res.getDetallesCompraVP)
	r.Get("/detallesCompraVP/<idDetalleCompraVP>", res.getDetalleCompraVPPorId)
	r.Get("/detallesCompraVP/porCompra/<idCompra>", res.getDetalleCompraVPPorIdCompra)
	r.Post("/detallesCompraVP", res.crearDetalleCompraVP)
	r.Put("/detallesCompraVP", res.actualizarDetalleCompraVP)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getDetallesCompraVP(c *routing.Context) error {
	detallesCompraVP, err := r.service.GetDetallesCompraVP(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(detallesCompraVP)
}

func (r resource) crearDetalleCompraVP(c *routing.Context) error {
	var input CreateDetalleCompraVPRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleCompraVP, err := r.service.CrearDetalleCompraVP(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleCompraVP, http.StatusCreated)
}

func (r resource) actualizarDetalleCompraVP(c *routing.Context) error {
	var input UpdateDetalleCompraVPRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleCompraVP, err := r.service.ActualizarDetalleCompraVP(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleCompraVP, http.StatusCreated)
}

func (r resource) getDetalleCompraVPPorId(c *routing.Context) error {
	idDetalleCompraVP, _ := strconv.Atoi(c.Param("idDetalleCompraVP"))
	detalleCompraVP, err := r.service.GetDetalleCompraVPPorId(c.Request.Context(), idDetalleCompraVP)
	if err != nil {
		return err
	}

	return c.Write(detalleCompraVP)
}

func (r resource) getDetalleCompraVPPorIdCompra(c *routing.Context) error {
	idFactura, _ := strconv.Atoi(c.Param("idCompra"))
	detallesFactura, err := r.service.GetDetalleCompraVPPorIdCompra(c.Request.Context(), idFactura)
	if err != nil {
		return err
	}

	return c.Write(detallesFactura)
}
