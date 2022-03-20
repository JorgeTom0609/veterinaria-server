package examen_mascota

import (
	"fmt"
	"net/http"
	"strconv"
	"veterinaria-server/internal/errors"
	"veterinaria-server/pkg/log"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/xuri/excelize/v2"
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

	ss, err := excelize.OpenFile("./Resultados.xlsx")
	if err != nil {
		return err
	}

	sheet := ss.GetSheetName(0)
	if err != nil {
		return err
	}
	ss.SetCellValue(sheet, "C4", input.Datos.Paciente)
	ss.SetCellValue(sheet, "G4", input.Datos.Especie)
	ss.SetCellValue(sheet, "C5", input.Datos.Propietario)
	ss.SetCellValue(sheet, "G5", input.Datos.Genero)
	ss.SetCellValue(sheet, "C6", input.Datos.Medico)
	ss.SetCellValue(sheet, "G6", input.Datos.Raza)
	ss.SetCellValue(sheet, "C7", input.Datos.Muestra)
	ss.SetCellValue(sheet, "B8", input.Datos.FechaLlenado.Format("2006-01-02 15:04:05"))

	for i := 0; i < len(input.Resultados); i++ {
		celdaA := "A" + strconv.Itoa((11 + i))
		celdaB := "B" + strconv.Itoa((11 + i))
		celdaC := "C" + strconv.Itoa((11 + i))
		celdaD := "D" + strconv.Itoa((11 + i))
		celdaE := "E" + strconv.Itoa((11 + i))
		celdaF := "F" + strconv.Itoa((11 + i))
		ss.SetCellValue(sheet, celdaA, input.Resultados[i].Parametro)
		ss.MergeCell(sheet, celdaA, celdaB)
		ss.SetCellValue(sheet, celdaC, input.Resultados[i].Resultado)
		ss.MergeCell(sheet, celdaC, celdaD)
		ss.SetCellValue(sheet, celdaE, input.Resultados[i].Alerta)
		ss.MergeCell(sheet, celdaE, celdaF)
	}

	fileName := fmt.Sprintf("resultado_%s_%s.xlsx", input.Datos.Paciente, input.Datos.FechaLlenado.Format("2006-01-02"))
	ss.SaveAs("./resources" + "/" + fileName)
	return c.Write(fileName)

}
