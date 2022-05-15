package productos

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
	r.Get("/productos", res.getProductos)
	r.Get("/productos/sinAsignar/<idProveedor>", res.getProductosSinAsignarAProveedor)
	r.Get("/productos/conStock", res.getProductosConStock)
	r.Get("/productos/usoInterno", res.getProductosUsoInterno)
	r.Get("/productos/<idProducto>", res.getProductoPorId)
	r.Get("/productos/comparacion/<idProveedor1>/<idProveedor2>", res.getProductosAComparar)
	r.Post("/productos", res.crearProducto)
	r.Put("/productos", res.actualizarProducto)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getProductos(c *routing.Context) error {
	productos, err := r.service.GetProductos(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(productos)
}

func (r resource) getProductosSinAsignarAProveedor(c *routing.Context) error {
	idProveedor, _ := strconv.Atoi(c.Param("idProveedor"))
	productos, err := r.service.GetProductosSinAsignarAProveedor(c.Request.Context(), idProveedor)
	if err != nil {
		return err
	}
	return c.Write(productos)
}

func (r resource) getProductosAComparar(c *routing.Context) error {
	idProveedor1, _ := strconv.Atoi(c.Param("idProveedor1"))
	idProveedor2, _ := strconv.Atoi(c.Param("idProveedor2"))
	productos, err := r.service.GetProductosAComparar(c.Request.Context(), idProveedor1, idProveedor2)
	if err != nil {
		return err
	}
	return c.Write(productos)
}

func (r resource) getProductosConStock(c *routing.Context) error {
	productos, err := r.service.GetProductosConStock(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(productos)
}

func (r resource) getProductosUsoInterno(c *routing.Context) error {
	productos, err := r.service.GetProductosUsoInterno(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(productos)
}

func (r resource) crearProducto(c *routing.Context) error {
	var input CreateProductoRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	producto, err := r.service.CrearProducto(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(producto, http.StatusCreated)
}

func (r resource) actualizarProducto(c *routing.Context) error {
	var input UpdateProductoRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	producto, err := r.service.ActualizarProducto(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(producto, http.StatusCreated)
}

func (r resource) getProductoPorId(c *routing.Context) error {
	idProducto, _ := strconv.Atoi(c.Param("idProducto"))
	producto, err := r.service.GetProductoPorId(c.Request.Context(), idProducto)
	if err != nil {
		return err
	}
	return c.Write(producto)
}
