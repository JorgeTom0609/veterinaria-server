package detalle_servicio_hospitalizacion

import (
	"net/http"
	"strconv"
	"veterinaria-server/internal/detalle_hospitalizacion"
	"veterinaria-server/internal/detalle_uso_servicio"
	"veterinaria-server/internal/errors"
	"veterinaria-server/internal/hospitalizacion"
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
	r.Get("/detallesServicioHospitalizacion", res.getDetallesServicioHospitalizacion)
	r.Get("/detallesServicioHospitalizacion/<idDetalleServicioHospitalizacion>", res.getDetalleServicioHospitalizacionPorId)
	r.Post("/detallesServicioHospitalizacion", res.crearDetalleServicioHospitalizacion)
	r.Put("/detallesServicioHospitalizacion/conDetalle", res.crearDetalleServicioHospitalizacionConDetalles)
	r.Put("/detallesServicioHospitalizacion", res.actualizarDetalleServicioHospitalizacion)
}

type resource struct {
	service Service
	logger  log.Logger
	db      *dbcontext.DB
}

func (r resource) getDetallesServicioHospitalizacion(c *routing.Context) error {
	detallesServicioHospitalizacion, err := r.service.GetDetallesServicioHospitalizacion(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(detallesServicioHospitalizacion)
}

func (r resource) crearDetalleServicioHospitalizacion(c *routing.Context) error {
	var input CreateDetalleServicioHospitalizacionRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleServicioHospitalizacion, err := r.service.CrearDetalleServicioHospitalizacion(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleServicioHospitalizacion, http.StatusCreated)
}

func (r resource) actualizarDetalleServicioHospitalizacion(c *routing.Context) error {
	var input UpdateDetalleServicioHospitalizacionRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleServicioHospitalizacion, err := r.service.ActualizarDetalleServicioHospitalizacion(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleServicioHospitalizacion, http.StatusCreated)
}

func (r resource) getDetalleServicioHospitalizacionPorId(c *routing.Context) error {
	idDetalleServicioHospitalizacion, _ := strconv.Atoi(c.Param("idDetalleServicioHospitalizacion"))
	detalleServicioHospitalizacion, err := r.service.GetDetalleServicioHospitalizacionPorId(c.Request.Context(), idDetalleServicioHospitalizacion)
	if err != nil {
		return err
	}
	return c.Write(detalleServicioHospitalizacion)
}

func (r resource) crearDetalleServicioHospitalizacionConDetalles(c *routing.Context) error {
	var input CreateDetalleServicioHospitalizacionConDetallesRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	detalleServicioHospitalizacionG, err := r.service.CrearDetalleServicioHospitalizacion(c.Request.Context(), input.DetalleServicioHospitalizacion)
	if err != nil {
		return err
	}

	sdh := detalle_hospitalizacion.NewService(detalle_hospitalizacion.NewRepository(r.db, r.logger), r.logger)
	var detalleHospitalizacion detalle_hospitalizacion.CreateDetalleHospitalizacionRequest = detalle_hospitalizacion.CreateDetalleHospitalizacionRequest{
		IdHospitalizacion: detalleServicioHospitalizacionG.IdHospitalizacion,
		IdUsuario:         detalleServicioHospitalizacionG.IdUsuario,
		Descripcion:       "Se aplicó el siguiente servicio: " + input.DetalleServicioHospitalizacion.Servicio,
		Fecha:             detalleServicioHospitalizacionG.Fecha,
	}
	detalleHospitalizacionG, err := sdh.CrearDetalleHospitalizacion(c.Request.Context(), detalleHospitalizacion)

	//Guardar detalles
	detallesUsoServicioG := []detalle_uso_servicio.DetalleUsoServicio{}
	for i := 0; i < len(input.Productos); i++ {
		input.Productos[i].IdDetalleServicioHospitalizacion = detalleServicioHospitalizacionG.IdDetalleServicioHospitalizacion
		s := detalle_uso_servicio.NewService(detalle_uso_servicio.NewRepository(r.db, r.logger), r.logger)
		detalleUsoServicioG, err := s.CrearDetalleUsoServicio(c.Request.Context(), input.Productos[i])
		if err != nil {
			return err
		}
		if detalleUsoServicioG.Tabla == "lote" {
			s2 := lote.NewService(lote.NewRepository(r.db, r.logger), r.logger)
			productoVenderBD, err2 := s2.GetLotePorId(c.Request.Context(), detalleUsoServicioG.IdReferencia)
			if err2 != nil {
				return err2
			}
			_, err3 := s2.ActualizarLote(c.Request.Context(),
				lote.UpdateLoteRequest{
					IdLote:              productoVenderBD.IdLote,
					Descripcion:         productoVenderBD.Descripcion,
					IdProveedorProducto: productoVenderBD.IdProveedorProducto,
					FechaCaducidad:      productoVenderBD.FechaCaducidad,
					Stock:               productoVenderBD.Stock - int(detalleUsoServicioG.Cantidad),
				})
			if err3 != nil {
				return err3
			}
		} else {
			s2 := stock_individual.NewService(stock_individual.NewRepository(r.db, r.logger), r.logger)
			productoVenderBD, err2 := s2.GetStockIndividualPorId(c.Request.Context(), detalleUsoServicioG.IdReferencia)
			if err2 != nil {
				return err2
			}
			_, err3 := s2.ActualizarStockIndividual(c.Request.Context(),
				stock_individual.UpdateStockIndividualRequest{
					IdLote:            productoVenderBD.IdLote,
					Descripcion:       productoVenderBD.Descripcion,
					IdStockIndividual: productoVenderBD.IdStockIndividual,
					CantidadInicial:   productoVenderBD.CantidadInicial,
					Cantidad:          productoVenderBD.Cantidad - detalleUsoServicioG.Cantidad,
				})
			if err3 != nil {
				return err3
			}
			if (productoVenderBD.Cantidad - detalleUsoServicioG.Cantidad) == float32(0) {
				s2 := lote.NewService(lote.NewRepository(r.db, r.logger), r.logger)
				productoVenderBD, err2 := s2.GetLotePorId(c.Request.Context(), productoVenderBD.IdLote)
				if err2 != nil {
					return err2
				}
				_, err3 := s2.ActualizarLote(c.Request.Context(),
					lote.UpdateLoteRequest{
						IdLote:              productoVenderBD.IdLote,
						Descripcion:         productoVenderBD.Descripcion,
						IdProveedorProducto: productoVenderBD.IdProveedorProducto,
						FechaCaducidad:      productoVenderBD.FechaCaducidad,
						Stock:               productoVenderBD.Stock - 1,
					})
				if err3 != nil {
					return err3
				}
			}
		}

		detallesUsoServicioG = append(detallesUsoServicioG, detalleUsoServicioG)
	}

	//Aumenta el valor de la hospitalizacion
	h := hospitalizacion.NewService(hospitalizacion.NewRepository(r.db, r.logger), r.logger)
	hospitalizacionBD, errR := h.GetHospitalizacionPorId(c.Request.Context(), detalleServicioHospitalizacionG.IdHospitalizacion)
	if errR != nil {
		return errR
	}

	hospitalizacionG, errG := h.ActualizarHospitalizacion(c.Request.Context(), hospitalizacion.UpdateHospitalizacionRequest{
		IdHospitalizacion:     hospitalizacionBD.IdHospitalizacion,
		Valor:                 hospitalizacionBD.Valor + detalleServicioHospitalizacionG.Valor,
		IdConsulta:            hospitalizacionBD.IdConsulta,
		Motivo:                hospitalizacionBD.Motivo,
		FechaIngreso:          hospitalizacionBD.FechaIngreso,
		FechaSalida:           hospitalizacionBD.FechaSalida,
		Abono:                 hospitalizacionBD.Abono,
		AutorizaExamenes:      hospitalizacionBD.AuorizaExamenes,
		EstadoHospitalizacion: hospitalizacionBD.EstadoHospitalizacion,
	})
	if errG != nil {
		return errG
	}

	var result = struct {
		DetalleServicioHospitalizacion DetalleServicioHospitalizacion
		Productos                      []detalle_uso_servicio.DetalleUsoServicio
		DetalleHospitalizacion         detalle_hospitalizacion.DetalleHospitalizacion
		Hospitalizacion                hospitalizacion.Hospitalizacion
	}{detalleServicioHospitalizacionG, detallesUsoServicioG, detalleHospitalizacionG, hospitalizacionG}

	return c.WriteWithStatus(result, http.StatusCreated)
}
