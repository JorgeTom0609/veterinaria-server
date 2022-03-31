package factura

import (
	"net/http"
	"strconv"
	"veterinaria-server/internal/clientes"
	"veterinaria-server/internal/detalle_factura"
	"veterinaria-server/internal/entity"
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
	r.Get("/facturas", res.getFacturas)
	r.Get("/facturas/conDatos", res.getFacturasConDatos)
	r.Get("/facturas/<idFactura>", res.getFacturaPorId)
	r.Post("/facturas", res.crearFactura)
	r.Post("/facturas/conDetalle", res.crearFacturaConDetalles)
	r.Put("/facturas", res.actualizarFactura)
}

type resource struct {
	service Service
	logger  log.Logger
	db      *dbcontext.DB
}

func (r resource) getFacturas(c *routing.Context) error {
	facturas, err := r.service.GetFacturas(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(facturas)
}

func (r resource) getFacturasConDatos(c *routing.Context) error {
	facturas, err := r.service.GetFacturasConDatos(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(facturas)
}

func (r resource) crearFactura(c *routing.Context) error {
	var input CreateFacturaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	factura, err := r.service.CrearFactura(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(factura, http.StatusCreated)
}

func (r resource) actualizarFactura(c *routing.Context) error {
	var input UpdateFacturaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	factura, err := r.service.ActualizarFactura(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(factura, http.StatusCreated)
}

func (r resource) getFacturaPorId(c *routing.Context) error {
	idFactura, _ := strconv.Atoi(c.Param("idFactura"))
	factura, err := r.service.GetFacturaPorId(c.Request.Context(), idFactura)
	if err != nil {
		return err
	}
	return c.Write(factura)
}

func (r resource) crearFacturaConDetalles(c *routing.Context) error {
	var input CreateFacturaConDetallesRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	var clienteG clientes.Cliente

	if input.Factura.IdCliente == 0 {
		var err error
		s1 := clientes.NewService(clientes.NewRepository(r.db, r.logger), r.logger)
		clienteG, err = s1.CrearCliente(c.Request.Context(), input.Cliente)
		if err != nil {
			return err
		}
		input.Factura.IdCliente = clienteG.IdCliente
	} else {
		clienteG = clientes.Cliente{
			Cliente: entity.Cliente{
				IdCliente:    input.Factura.IdCliente,
				Nombres:      input.Cliente.Nombres,
				Apellidos:    input.Cliente.Apellidos,
				Cedula:       input.Cliente.Cedula,
				Correo:       input.Cliente.Correo,
				Telefono:     input.Cliente.Telefono,
				Direccion:    input.Cliente.Direccion,
				Nacionalidad: input.Cliente.Nacionalidad,
			},
		}
	}

	facturaG, err := r.service.CrearFactura(c.Request.Context(), input.Factura)
	if err != nil {
		return err
	}
	//Guardar detalles factura
	detallesFacturaG := []detalle_factura.DetalleFactura{}
	for i := 0; i < len(input.DetallesFactura); i++ {
		input.DetallesFactura[i].IdFactura = facturaG.IdFactura
		s := detalle_factura.NewService(detalle_factura.NewRepository(r.db, r.logger), r.logger)
		detalleFacturaG, err := s.CrearDetalleFactura(c.Request.Context(), input.DetallesFactura[i])
		if err != nil {
			return err
		}
		s2 := producto_vp.NewService(producto_vp.NewRepository(r.db, r.logger), r.logger)
		productoVenderBD, err2 := s2.GetProductoVPPorId(c.Request.Context(), detalleFacturaG.IdProductoVp)
		if err2 != nil {
			return err2
		}
		_, err3 := s2.ActualizarProductoVP(c.Request.Context(),
			producto_vp.UpdateProductoVPRequest{
				IdProductoVP: productoVenderBD.IdProductoVP,
				Descripcion:  productoVenderBD.Descripcion,
				PrecioCompra: productoVenderBD.PrecioCompra,
				PrecioVenta:  productoVenderBD.PrecioVenta,
				Stock:        productoVenderBD.Stock - detalleFacturaG.Cantidad,
				StockMinimo:  productoVenderBD.StockMinimo,
			})
		if err3 != nil {
			return err3
		}
		detallesFacturaG = append(detallesFacturaG, detalleFacturaG)
	}

	var result = struct {
		Cliente         clientes.Cliente
		Factura         Factura
		DetallesFactura []detalle_factura.DetalleFactura
	}{clienteG, facturaG, detallesFacturaG}

	return c.WriteWithStatus(result, http.StatusCreated)
}
