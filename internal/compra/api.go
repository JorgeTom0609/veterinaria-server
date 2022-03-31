package compra

import (
	"net/http"
	"strconv"
	"veterinaria-server/internal/detalle_compra_vp"
	"veterinaria-server/internal/errors"
	"veterinaria-server/internal/producto_vp"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger, db *dbcontext.DB) {
	res := resource{service, logger, db}
	r.Use(authHandler)
	// the following endpoints require a valid JWT
	r.Get("/compras", res.getCompras)
	r.Get("/compras/conDatos", res.getComprasConDatos)
	r.Get("/compras/<idCompra>", res.getCompraPorId)
	r.Post("/compras", res.crearCompra)
	r.Post("/compras/conDetalle", res.crearCompraConDetalles)
	r.Put("/compras", res.actualizarCompra)
}

type resource struct {
	service Service
	logger  log.Logger
	db      *dbcontext.DB
}

func (r resource) getCompras(c *routing.Context) error {
	compras, err := r.service.GetCompras(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(compras)
}

func (r resource) getComprasConDatos(c *routing.Context) error {
	compras, err := r.service.GetComprasConDatos(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(compras)
}

func (r resource) crearCompra(c *routing.Context) error {
	var input CreateCompraRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	compra, err := r.service.CrearCompra(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(compra, http.StatusCreated)
}

func (r resource) actualizarCompra(c *routing.Context) error {
	var input UpdateCompraRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	compra, err := r.service.ActualizarCompra(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(compra, http.StatusCreated)
}

func (r resource) getCompraPorId(c *routing.Context) error {
	idCompra, _ := strconv.Atoi(c.Param("idCompra"))
	compra, err := r.service.GetCompraPorId(c.Request.Context(), idCompra)
	if err != nil {
		return err
	}
	return c.Write(compra)
}

func (r resource) crearCompraConDetalles(c *routing.Context) error {
	var input CreateCompraConDetallesRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	compraG, err := r.service.CrearCompra(c.Request.Context(), input.Compra)
	if err != nil {
		return err
	}
	//Guardar detalles compra
	detallesCompraG := []detalle_compra_vp.DetalleCompraVP{}
	for i := 0; i < len(input.DetallesComprasVP); i++ {
		input.DetallesComprasVP[i].IdCompra = compraG.IdCompra
		s := detalle_compra_vp.NewService(detalle_compra_vp.NewRepository(r.db, r.logger), r.logger)
		detalleCompraG, err := s.CrearDetalleCompraVP(c.Request.Context(), input.DetallesComprasVP[i])
		if err != nil {
			return err
		}
		s2 := producto_vp.NewService(producto_vp.NewRepository(r.db, r.logger), r.logger)
		productoCompraBD, err2 := s2.GetProductoVPPorId(c.Request.Context(), detalleCompraG.IdProductoVp)
		if err2 != nil {
			return err2
		}
		_, err3 := s2.ActualizarProductoVP(c.Request.Context(),
			producto_vp.UpdateProductoVPRequest{
				IdProductoVP: productoCompraBD.IdProductoVP,
				Descripcion:  productoCompraBD.Descripcion,
				PrecioCompra: productoCompraBD.PrecioCompra,
				PrecioVenta:  productoCompraBD.PrecioVenta,
				Stock:        productoCompraBD.Stock + detalleCompraG.Cantidad,
				StockMinimo:  productoCompraBD.StockMinimo,
			})
		if err3 != nil {
			return err3
		}
		detallesCompraG = append(detallesCompraG, detalleCompraG)
	}

	var result = struct {
		Compra            Compras
		DetallesComprasVP []detalle_compra_vp.DetalleCompraVP
	}{compraG, detallesCompraG}

	return c.WriteWithStatus(result, http.StatusCreated)
}
