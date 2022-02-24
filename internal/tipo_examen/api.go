package tipo_examen

import (
	"net/http"
	"strconv"
	"veterinaria-server/internal/detalle_examen_cualitativo"
	"veterinaria-server/internal/detalle_examen_cuantitativo"
	"veterinaria-server/internal/detalle_examen_informativo"
	"veterinaria-server/internal/errors"
	"veterinaria-server/internal/examen_mascota"
	"veterinaria-server/internal/resultado_examen_cualitativo"
	"veterinaria-server/internal/resultado_examen_cuantitativo"
	"veterinaria-server/internal/resultado_examen_informativo"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger, db *dbcontext.DB) {
	res := resource{service, logger, db}
	r.Use(authHandler)
	r.Get("/tipo_examen", res.getTipoExamenes)
	r.Get("/tipo_examen/<idTipoExamen>", res.getTipoExamenPorId)
	r.Get("/tipo_examen/por_especie/<idEspecie>", res.getTipoExamenPorEspecie)
	r.Get("/tipo_examen/detalles/<idTipoExamen>", res.getDetallesExamenPorTipoExamen)
	r.Post("/tipo_examen", res.crearTipoExamen)
	r.Post("/tipo_examen/con_resultados", res.guardarResultados)
	r.Put("/tipo_examen", res.actualizarTipoExamen)
	r.Put("/tipo_examen/con_detalles", res.actualizarTipoExamenConDetalles)
}

type resource struct {
	service Service
	logger  log.Logger
	db      *dbcontext.DB
}

func (r resource) getTipoExamenes(c *routing.Context) error {
	tipoExamenes, err := r.service.GetTipoExamenes(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(tipoExamenes)
}

func (r resource) getTipoExamenPorEspecie(c *routing.Context) error {
	idEspecie, _ := strconv.Atoi(c.Param("idEspecie"))
	tipoExamenes, err := r.service.GetTipoExamenPorEspecie(c.Request.Context(), idEspecie)
	if err != nil {
		return err
	}
	return c.Write(tipoExamenes)
}

func (r resource) getDetallesExamenPorTipoExamen(c *routing.Context) error {
	idTipoExamen, _ := strconv.Atoi(c.Param("idTipoExamen"))
	detallesExamen, err := r.service.GetDetallesExamenPorTipoExamen(c.Request.Context(), idTipoExamen)
	if err != nil {
		return err
	}
	return c.Write(detallesExamen)
}

func (r resource) crearTipoExamen(c *routing.Context) error {
	var input CreateTipoExamenRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	tipoExamen, err := r.service.CrearTipoExamen(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(tipoExamen, http.StatusCreated)
}

func (r resource) actualizarTipoExamen(c *routing.Context) error {
	var input UpdateTipoExamenRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	tipoExamen, err := r.service.ActualizarTipoExamen(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(tipoExamen, http.StatusCreated)
}

func (r resource) actualizarTipoExamenConDetalles(c *routing.Context) error {
	var input UpdateTipoExamenConDetallesRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	tipoExamen, err := r.service.ActualizarTipoExamen(c.Request.Context(), input.TipoExamen)
	if err != nil {
		return err
	}
	//Guardar detalles cualitativos
	cualitativosG := []detalle_examen_cualitativo.DetallesExamenCualitativo{}
	for i := 0; i < len(input.Cualitativos); i++ {
		input.Cualitativos[i].IdTipoExamen = tipoExamen.IdTipoExamen
		s := detalle_examen_cualitativo.NewService(detalle_examen_cualitativo.NewRepository(r.db, r.logger), r.logger)
		detalleExamenCualitativo, err := s.ActualizarDetalleExamenCualitativo(c.Request.Context(), input.Cualitativos[i])
		if err != nil {
			return err
		}
		cualitativosG = append(cualitativosG, detalleExamenCualitativo)
	}
	//Guardar detalles cuantitativos
	cuantitativosG := []detalle_examen_cuantitativo.DetallesExamenCuantitativo{}
	for i := 0; i < len(input.Cuantitativos); i++ {
		input.Cuantitativos[i].IdTipoExamen = tipoExamen.IdTipoExamen
		s := detalle_examen_cuantitativo.NewService(detalle_examen_cuantitativo.NewRepository(r.db, r.logger), r.logger)
		detalleExamenCuantitativo, err := s.ActualizarDetalleExamenCuantitativo(c.Request.Context(), input.Cuantitativos[i])
		if err != nil {
			return err
		}
		cuantitativosG = append(cuantitativosG, detalleExamenCuantitativo)
	}

	//Guardar detalles informativos
	informativosG := []detalle_examen_informativo.DetallesExamenInformativo{}
	for i := 0; i < len(input.Informativos); i++ {
		input.Informativos[i].IdTipoExamen = tipoExamen.IdTipoExamen
		s := detalle_examen_informativo.NewService(detalle_examen_informativo.NewRepository(r.db, r.logger), r.logger)
		detalleExamenInformativo, err := s.ActualizarDetalleExamenInformativo(c.Request.Context(), input.Informativos[i])
		if err != nil {
			return err
		}
		informativosG = append(informativosG, detalleExamenInformativo)
	}

	var result = struct {
		TipoDeExamen  TipoExamen
		Cualitativos  []detalle_examen_cualitativo.DetallesExamenCualitativo
		Cuantitativos []detalle_examen_cuantitativo.DetallesExamenCuantitativo
		Informativos  []detalle_examen_informativo.DetallesExamenInformativo
	}{tipoExamen, cualitativosG, cuantitativosG, informativosG}

	return c.WriteWithStatus(result, http.StatusCreated)
}

func (r resource) getTipoExamenPorId(c *routing.Context) error {
	idTipoExamen, _ := strconv.Atoi(c.Param("idTipoExamen"))
	tipoExamen, err := r.service.GetTipoExamenPorId(c.Request.Context(), idTipoExamen)
	if err != nil {
		return err
	}

	return c.Write(tipoExamen)
}

func (r resource) guardarResultados(c *routing.Context) error {
	var input CreateResultadosRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	//Guardar resultados cualitativos
	resultadosCualitativosG := []resultado_examen_cualitativo.ResultadoDetalleCualitativo{}
	for i := 0; i < len(input.Cualitativos); i++ {
		s := resultado_examen_cualitativo.NewService(resultado_examen_cualitativo.NewRepository(r.db, r.logger), r.logger)
		resultadoCualitativo, err := s.CrearResultadoDetalleCualitativo(c.Request.Context(), input.Cualitativos[i])
		if err != nil {
			return err
		}
		resultadosCualitativosG = append(resultadosCualitativosG, resultadoCualitativo)
	}

	//Guardar resultados cuantitativos
	resultadosCuantitativosG := []resultado_examen_cuantitativo.ResultadoDetalleCuantitativo{}
	for i := 0; i < len(input.Cuantitativos); i++ {
		s := resultado_examen_cuantitativo.NewService(resultado_examen_cuantitativo.NewRepository(r.db, r.logger), r.logger)
		resultadoCuantitativo, err := s.CrearResultadoDetalleCuantitativo(c.Request.Context(), input.Cuantitativos[i])
		if err != nil {
			return err
		}
		resultadosCuantitativosG = append(resultadosCuantitativosG, resultadoCuantitativo)
	}

	//Guardar resultados informativos
	resultadosInformativosG := []resultado_examen_informativo.ResultadoDetalleInformativo{}
	for i := 0; i < len(input.Informativos); i++ {
		s := resultado_examen_informativo.NewService(resultado_examen_informativo.NewRepository(r.db, r.logger), r.logger)
		resultadoInformativo, err := s.CrearResultadoDetalleInformativo(c.Request.Context(), input.Informativos[i])
		if err != nil {
			return err
		}
		resultadosInformativosG = append(resultadosInformativosG, resultadoInformativo)
	}
	_, err := examen_mascota.ActualizarEstadoExamenMascota(c.Request.Context(), input.IdExamenMascota, r.db)
	if err != nil {
		return err
	}

	var result = struct {
		Cualitativos  []resultado_examen_cualitativo.ResultadoDetalleCualitativo
		Cuantitativos []resultado_examen_cuantitativo.ResultadoDetalleCuantitativo
		Informativos  []resultado_examen_informativo.ResultadoDetalleInformativo
	}{resultadosCualitativosG, resultadosCuantitativosG, resultadosInformativosG}
	return c.WriteWithStatus(result, http.StatusCreated)
}
