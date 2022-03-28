package producto_vp

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
	r.Get("/productosVP", res.getProductosVP)
	r.Get("/productosVP/conStock", res.getProductosVPConStock)
	r.Get("/productosVP/pocoStock", res.getProductosVPPocoStock)
	r.Get("/productosVP/<idProductoVP>", res.getProductoVPPorId)
	r.Post("/productosVP", res.crearProductoVP)
	r.Put("/productosVP", res.actualizarProductoVP)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getProductosVP(c *routing.Context) error {
	productosVP, err := r.service.GetProductosVP(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(productosVP)
}

func (r resource) getProductosVPConStock(c *routing.Context) error {
	productosVP, err := r.service.GetProductosVPConStock(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(productosVP)
}

func (r resource) getProductosVPPocoStock(c *routing.Context) error {
	productosVP, err := r.service.GetProductosVPPocoStock(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(productosVP)
}

func (r resource) crearProductoVP(c *routing.Context) error {
	var input CreateProductoVPRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	productoVP, err := r.service.CrearProductoVP(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(productoVP, http.StatusCreated)
}

func (r resource) actualizarProductoVP(c *routing.Context) error {
	var input UpdateProductoVPRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	productoVP, err := r.service.ActualizarProductoVP(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(productoVP, http.StatusCreated)
}

func (r resource) getProductoVPPorId(c *routing.Context) error {
	idProductoVP, _ := strconv.Atoi(c.Param("idProductoVP"))
	productoVP, err := r.service.GetProductoVPPorId(c.Request.Context(), idProductoVP)
	if err != nil {
		return err
	}

	return c.Write(productoVP)
}
