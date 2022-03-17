package examen_mascota

import (
	"fmt"
	"net/http"
	"strconv"
	"veterinaria-server/internal/errors"
	"veterinaria-server/pkg/log"

	"baliance.com/gooxml/spreadsheet"
	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}
	r.Use(authHandler)
	// the following endpoints require a valid JWT
	r.Get("/examenesMascota", res.getExamenesMascota)
	r.Get("/examenesMascota/<idExamenMascota>", res.getExamenMascotaPorId)
	r.Get("/examenesMascota/examenes/<idMascota>/<estado>", res.getExamenesMascotaPorMascotayEstado)
	r.Get("/examenesMascota/examenesPorEstado/<estado>", res.getExamenesMascotaPorEstado)
	r.Get("/examenesMascota/resultados/<idExamenMascota>", res.obtenerResultadosPorExamen)
	r.Post("/examenesMascota", res.crearExamenMascota)
	r.Post("/examenesMascota/archivo", res.create)
	r.Put("/examenesMascota", res.actualizarExamenMascota)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getExamenesMascota(c *routing.Context) error {
	examenesMascota, err := r.service.GetExamenesMascota(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(examenesMascota)
}

func (r resource) getExamenesMascotaPorMascotayEstado(c *routing.Context) error {
	idExamenMascota, _ := strconv.Atoi(c.Param("idMascota"))
	estado := c.Param("estado")
	examenesMascota, err := r.service.GetExamenesMascotaPorMascotayEstado(c.Request.Context(), idExamenMascota, estado)
	if err != nil {
		return err
	}
	return c.Write(examenesMascota)
}

func (r resource) getExamenesMascotaPorEstado(c *routing.Context) error {
	estado := c.Param("estado")
	examenesMascota, err := r.service.GetExamenesMascotaPorEstado(c.Request.Context(), estado)
	if err != nil {
		return err
	}
	return c.Write(examenesMascota)
}

func (r resource) crearExamenMascota(c *routing.Context) error {
	var input CreateExamenMascotaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	examenesMascota, err := r.service.CrearExamenMascota(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(examenesMascota, http.StatusCreated)
}

func (r resource) actualizarExamenMascota(c *routing.Context) error {
	var input UpdateExamenMascotaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	examenesMascota, err := r.service.ActualizarExamenMascota(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(examenesMascota, http.StatusCreated)
}

func (r resource) getExamenMascotaPorId(c *routing.Context) error {
	idExamenMascota, _ := strconv.Atoi(c.Param("idExamenMascota"))
	examenesMascota, err := r.service.GetExamenMascotaPorId(c.Request.Context(), idExamenMascota)
	if err != nil {
		return err
	}
	return c.Write(examenesMascota)
}

func (r resource) obtenerResultadosPorExamen(c *routing.Context) error {
	idExamenMascota, _ := strconv.Atoi(c.Param("idExamenMascota"))
	resultados, err := r.service.ObtenerResultadosPorExamen(c.Request.Context(), idExamenMascota)
	if err != nil {
		return err
	}
	return c.Write(resultados)
}

func (r resource) create(c *routing.Context) error {
	var input ResultadosRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	ss, err := spreadsheet.Open("./Resultados.xlsx")
	if err != nil {
		return err
	}

	sheet, err := ss.GetSheet("Hoja1")
	if err != nil {
		return err
	}

	sheet.Cell("C4").Clear()
	sheet.Cell("C4").SetString(input.Datos.Paciente)
	sheet.Cell("G4").Clear()
	sheet.Cell("G4").SetString(input.Datos.Especie)

	sheet.Cell("C5").Clear()
	sheet.Cell("C5").SetString(input.Datos.Propietario)
	sheet.Cell("G5").Clear()
	sheet.Cell("G5").SetString(input.Datos.Genero)

	sheet.Cell("C6").Clear()
	sheet.Cell("C6").SetString(input.Datos.Medico)
	sheet.Cell("G6").Clear()
	sheet.Cell("G6").SetString(input.Datos.Raza)

	sheet.Cell("C7").Clear()
	sheet.Cell("C7").SetString(input.Datos.Muestra)

	sheet.Cell("B8").Clear()
	sheet.Cell("B8").SetString(input.Datos.FechaLlenado.Format("2006-01-02 15:04:05"))

	for i := 0; i < len(input.Resultados); i++ {
		celdaA := "A" + strconv.Itoa((11 + i))
		celdaB := "B" + strconv.Itoa((11 + i))
		celdaC := "C" + strconv.Itoa((11 + i))
		celdaD := "D" + strconv.Itoa((11 + i))
		celdaE := "E" + strconv.Itoa((11 + i))
		celdaF := "F" + strconv.Itoa((11 + i))
		sheet.Cell(celdaA).Clear()
		sheet.Cell(celdaA).SetString(input.Resultados[i].Parametro)
		sheet.AddMergedCells(celdaA, celdaB)
		sheet.Cell(celdaC).Clear()
		sheet.Cell(celdaC).SetString(input.Resultados[i].Resultado)
		sheet.AddMergedCells(celdaC, celdaD)
		sheet.Cell(celdaE).Clear()
		sheet.Cell(celdaE).SetString(input.Resultados[i].Alerta)
		sheet.AddMergedCells(celdaE, celdaF)
	}
	fileName := fmt.Sprintf("resultado_%s_%s.xlsx", input.Datos.Paciente, input.Datos.FechaLlenado.Format("2006-01-02"))
	ss.SaveToFile("./resources" + "/" + fileName)
	return c.Write(fileName)

}
