package detalle_servicio_consulta

import (
	"net/http"
	"strconv"
	"veterinaria-server/internal/consultas"
	"veterinaria-server/internal/detalle_uso_servicio_consulta"
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
	r.Get("/detallesServicioConsulta", res.getDetallesServicioConsulta)
	r.Get("/detallesServicioConsulta/<idDetalleServicioConsulta>", res.getDetalleServicioConsultaPorId)
	r.Get("/detallesServicioConsulta/porConsulta/<idConsulta>", res.getDetalleServicioConsultaPorConsulta)
	r.Post("/detallesServicioConsulta", res.crearDetalleServicioConsulta)
	r.Put("/detallesServicioConsulta/conDetalle", res.crearDetalleServicioConsultaConDetalles)
	r.Put("/detallesServicioConsulta", res.actualizarDetalleServicioConsulta)
}

type resource struct {
	service Service
	logger  log.Logger
	db      *dbcontext.DB
}

func (r resource) getDetallesServicioConsulta(c *routing.Context) error {
	detallesServicioConsulta, err := r.service.GetDetallesServicioConsulta(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(detallesServicioConsulta)
}

func (r resource) crearDetalleServicioConsulta(c *routing.Context) error {
	var input CreateDetalleServicioConsultaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleServicioConsulta, err := r.service.CrearDetalleServicioConsulta(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleServicioConsulta, http.StatusCreated)
}

func (r resource) actualizarDetalleServicioConsulta(c *routing.Context) error {
	var input UpdateDetalleServicioConsultaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	detalleServicioConsulta, err := r.service.ActualizarDetalleServicioConsulta(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(detalleServicioConsulta, http.StatusCreated)
}

func (r resource) getDetalleServicioConsultaPorId(c *routing.Context) error {
	idDetalleServicioConsulta, _ := strconv.Atoi(c.Param("idDetalleServicioConsulta"))
	detalleServicioConsulta, err := r.service.GetDetalleServicioConsultaPorId(c.Request.Context(), idDetalleServicioConsulta)
	if err != nil {
		return err
	}
	return c.Write(detalleServicioConsulta)
}

func (r resource) getDetalleServicioConsultaPorConsulta(c *routing.Context) error {
	idConsulta, _ := strconv.Atoi(c.Param("idConsulta"))
	detalleServicioConsulta, err := r.service.GetDetalleServicioConsultaPorConsulta(c.Request.Context(), idConsulta)
	if err != nil {
		return err
	}
	return c.Write(detalleServicioConsulta)
}

func (r resource) crearDetalleServicioConsultaConDetalles(c *routing.Context) error {
	var input CreateDetalleServicioConsultaConDetallesRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	detalleServicioConsultaG, err := r.service.CrearDetalleServicioConsulta(c.Request.Context(), input.DetalleServicioConsulta)
	if err != nil {
		return err
	}

	//Guardar detalles
	detallesUsoServicioConsultaG := []detalle_uso_servicio_consulta.DetalleUsoServicioConsulta{}
	for i := 0; i < len(input.Productos); i++ {
		input.Productos[i].IdDetalleServicioConsulta = detalleServicioConsultaG.IdDetalleServicioConsulta
		s := detalle_uso_servicio_consulta.NewService(detalle_uso_servicio_consulta.NewRepository(r.db, r.logger), r.logger)
		detalleUsoServicioConsultaG, err := s.CrearDetalleUsoServicioConsulta(c.Request.Context(), input.Productos[i])
		if err != nil {
			return err
		}
		if detalleUsoServicioConsultaG.Tabla == "lote" {
			s2 := lote.NewService(lote.NewRepository(r.db, r.logger), r.logger)
			productoVenderBD, err2 := s2.GetLotePorId(c.Request.Context(), detalleUsoServicioConsultaG.IdReferencia)
			if err2 != nil {
				return err2
			}
			_, err3 := s2.ActualizarLote(c.Request.Context(),
				lote.UpdateLoteRequest{
					IdLote:              productoVenderBD.IdLote,
					Descripcion:         productoVenderBD.Descripcion,
					IdProveedorProducto: productoVenderBD.IdProveedorProducto,
					FechaCaducidad:      productoVenderBD.FechaCaducidad,
					Stock:               productoVenderBD.Stock - int(detalleUsoServicioConsultaG.Cantidad),
				})
			if err3 != nil {
				return err3
			}
		} else {
			s2 := stock_individual.NewService(stock_individual.NewRepository(r.db, r.logger), r.logger)
			productoVenderBD, err2 := s2.GetStockIndividualPorId(c.Request.Context(), detalleUsoServicioConsultaG.IdReferencia)
			if err2 != nil {
				return err2
			}
			_, err3 := s2.ActualizarStockIndividual(c.Request.Context(),
				stock_individual.UpdateStockIndividualRequest{
					IdLote:            productoVenderBD.IdLote,
					Descripcion:       productoVenderBD.Descripcion,
					IdStockIndividual: productoVenderBD.IdStockIndividual,
					CantidadInicial:   productoVenderBD.CantidadInicial,
					Cantidad:          productoVenderBD.Cantidad - detalleUsoServicioConsultaG.Cantidad,
				})
			if err3 != nil {
				return err3
			}
			if (productoVenderBD.Cantidad - detalleUsoServicioConsultaG.Cantidad) == float32(0) {
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

		detallesUsoServicioConsultaG = append(detallesUsoServicioConsultaG, detalleUsoServicioConsultaG)
	}

	//Aumenta el valor de la consulta
	consul := consultas.NewService(consultas.NewRepository(r.db, r.logger), r.logger)
	consultaBD, errR := consul.GetConsultaPorId(c.Request.Context(), detalleServicioConsultaG.IdConsulta)
	if errR != nil {
		return errR
	}

	consultaG, errG := consul.ActualizarConsulta(c.Request.Context(), consultas.UpdateConsultaRequest{
		IdConsulta:             consultaBD.IdConsulta,
		IdMascota:              consultaBD.IdMascota,
		IdUsuario:              consultaBD.IdUsuario,
		Fecha:                  consultaBD.Fecha,
		Valor:                  consultaBD.Valor + detalleServicioConsultaG.Valor,
		Motivo:                 consultaBD.Motivo,
		Temperatura:            consultaBD.Temperatura,
		Peso:                   consultaBD.Peso,
		Tamaño:                 consultaBD.Tamaño,
		CondicionCorporal:      consultaBD.CondicionCorporal,
		NivelesDeshidratacion:  consultaBD.NivelesDeshidratacion,
		Diagnostico:            consultaBD.Diagnostico,
		Edad:                   consultaBD.Edad,
		TiempoLlenadoCapilar:   consultaBD.TiempoLlenadoCapilar,
		FrecuenciaCardiaca:     consultaBD.FrecuenciaCardiaca,
		FrecuenciaRespiratoria: consultaBD.FrecuenciaRespiratoria,
		EstadoConsulta:         consultaBD.EstadoConsulta,
	})
	if errG != nil {
		return errG
	}

	var result = struct {
		DetalleServicioConsulta DetalleServicioConsulta
		Productos               []detalle_uso_servicio_consulta.DetalleUsoServicioConsulta
		Consulta                consultas.Consulta
	}{detalleServicioConsultaG, detallesUsoServicioConsultaG, consultaG}

	return c.WriteWithStatus(result, http.StatusCreated)
}
