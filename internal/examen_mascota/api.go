package examen_mascota

import (
	"fmt"
	"net/http"
	"strconv"
	"veterinaria-server/internal/errors"
	"veterinaria-server/pkg/log"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/nguyenthenguyen/docx"
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
	r.Post("/examenesMascota/archivo", res.archivo)
	r.Post("/examenesMascota/autorizacion", res.autorizacion)
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

func (r resource) archivo(c *routing.Context) error {
	var input ResultadosRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	ss, err := excelize.OpenFile("./plantillas/Resultados.xlsx")
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

	fileName := fmt.Sprintf("Resultado-%s-%s.xlsx", input.Datos.Paciente, input.Datos.FechaLlenado.Format("2006-01-02"))
	ss.SaveAs("./resources" + "/" + fileName)
	return c.Write(fileName)
}

func (r resource) autorizacion(c *routing.Context) error {
	var input DatosMascotaDueñoRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	var nombreDoc string
	// Read from docx file
	switch input.NumAutorizacion {
	case 1:
		nombreDoc = "Anestesia"
	case 2:
		nombreDoc = "Eutanasia"
	case 3:
		nombreDoc = "Hospitalizacion"
	case 4:
		nombreDoc = "PlanSanitario"
	}
	rd, err := docx.ReadDocxFile("./plantillas/" + nombreDoc + ".docx")
	if err != nil {
		return err
	}
	docx1 := rd.Editable()
	docx1.Replace("cnombredueño", input.Propietario, -1)
	docx1.Replace("cnacionalidaddueño", input.Nacionalidad, -1)
	docx1.Replace("cdomiciliodueño", input.Direccion, -1)
	docx1.Replace("cnombremascota", input.Paciente, -1)
	docx1.Replace("cceduladueño", input.Cedula, -1)
	docx1.Replace("ccedula", input.Cedula, -1)
	docx1.Replace("csexomascota", input.Sexo, -1)
	docx1.Replace("cedadmascota", input.Edad, -1)
	docx1.Replace("crazamascota", input.Raza, -1)
	docx1.Replace("cenfermedadmascota", input.Enfermedad, -1)
	docx1.Replace("cintervención", input.Intervencion, -1)
	docx1.Replace("caño", strconv.Itoa(input.Fecha.Year()), -1)
	docx1.Replace("cabono", fmt.Sprintf("%f", input.Abono), -1)
	docx1.Replace("cprofesional", input.Profesional, -1)
	var mes string
	switch input.Fecha.Month() {
	case 1:
		mes = "Enero"
	case 2:
		mes = "Febrero"
	case 3:
		mes = "Marzo"
	case 4:
		mes = "Abril"
	case 5:
		mes = "Mayo"
	case 6:
		mes = "Junio"
	case 7:
		mes = "Julio"
	case 8:
		mes = "Agosto"
	case 9:
		mes = "Septiembre"
	case 10:
		mes = "Octubre"
	case 11:
		mes = "Noviembre"
	case 12:
		mes = "Diciembre"
	}
	docx1.Replace("cmes", mes, -1)
	docx1.Replace("cdías", strconv.Itoa(input.Fecha.Day()), -1)
	docx1.Replace("cdía", strconv.Itoa(input.Fecha.Day()), -1)
	fileName := fmt.Sprintf("%s-%s-%s.docx", nombreDoc, input.Paciente, input.Fecha.Format("2006-01-02"))
	docx1.WriteToFile("./resources" + "/" + fileName)
	rd.Close()
	return c.Write(fileName)
}