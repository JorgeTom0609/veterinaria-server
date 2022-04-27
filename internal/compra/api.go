package compra

import (
	"net/http"
	"strconv"
	"veterinaria-server/internal/detalle_compra"
	"veterinaria-server/internal/errors"
	"veterinaria-server/internal/lote"
	"veterinaria-server/internal/stock_individual"
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
	detallesComprasConLoteG := []DetallesComprasConLote{}
	//Guardar detalles compra
	for i := 0; i < len(input.DetallesCompras); i++ {
		s := lote.NewService(lote.NewRepository(r.db, r.logger), r.logger)
		loteG, err := s.CrearLote(c.Request.Context(), input.DetallesCompras[i].Lote)
		if err != nil {
			return err
		}

		stockIndiidualesG := []stock_individual.StockIndividual{}
		for j := 0; j < len(input.DetallesCompras[i].StocksIndividuales); j++ {
			input.DetallesCompras[i].StocksIndividuales[j].IdLote = loteG.IdLote
			s2 := stock_individual.NewService(stock_individual.NewRepository(r.db, r.logger), r.logger)
			stockIndividualG, err := s2.CrearStockIndividual(c.Request.Context(), input.DetallesCompras[i].StocksIndividuales[j])
			if err != nil {
				return err
			}
			stockIndiidualesG = append(stockIndiidualesG, stockIndividualG)
		}
		input.DetallesCompras[i].DetalleCompra.IdCompra = compraG.IdCompra
		input.DetallesCompras[i].DetalleCompra.IdLote = loteG.IdLote
		s3 := detalle_compra.NewService(detalle_compra.NewRepository(r.db, r.logger), r.logger)
		detalleCompraG, err := s3.CrearDetalleCompra(c.Request.Context(), input.DetallesCompras[i].DetalleCompra)
		if err != nil {
			return err
		}

		detallesComprasConLoteG = append(detallesComprasConLoteG, DetallesComprasConLote{
			Lote:               loteG.Lote,
			DetalleCompra:      detalleCompraG.DetalleCompra,
			StocksIndividuales: stockIndiidualesG,
		})
	}

	var result = struct {
		Compra          Compras
		DetallesCompras []DetallesComprasConLote
	}{compraG, detallesComprasConLoteG}

	return c.WriteWithStatus(result, http.StatusCreated)
}
