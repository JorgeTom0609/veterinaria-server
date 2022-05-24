package receta

import (
	"fmt"
	"math"
	"net/http"
	"runtime"
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
	r.Get("/recetas", res.getRecetas)
	r.Get("/recetas/<idReceta>", res.getRecetaPorId)
	r.Get("/recetas/porConsulta/<idConsulta>", res.getRecetaPorConsulta)
	r.Post("/recetas", res.crearReceta)
	r.Post("/recetas/archivo", res.archivo)
	r.Put("/recetas", res.actualizarReceta)
	r.Put("/recetas/toda", res.actualizarRecetaToda)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getRecetas(c *routing.Context) error {
	recetas, err := r.service.GetRecetas(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(recetas)
}

func (r resource) crearReceta(c *routing.Context) error {
	var input CreateRecetaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	receta, err := r.service.CrearReceta(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(receta, http.StatusCreated)
}

func (r resource) actualizarReceta(c *routing.Context) error {
	var input UpdateRecetaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	receta, err := r.service.ActualizarReceta(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(receta, http.StatusCreated)
}

func (r resource) actualizarRecetaToda(c *routing.Context) error {
	var input []UpdateRecetaRequest
	var recetas []Receta = []Receta{}
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	for i := 0; i < len(input); i++ {
		receta, err := r.service.ActualizarReceta(c.Request.Context(), input[i])
		if err != nil {
			return err
		}
		recetas = append(recetas, receta)
	}
	return c.WriteWithStatus(recetas, http.StatusCreated)
}

func (r resource) getRecetaPorId(c *routing.Context) error {
	idReceta, _ := strconv.Atoi(c.Param("idReceta"))
	receta, err := r.service.GetRecetaPorId(c.Request.Context(), idReceta)
	if err != nil {
		return err
	}
	return c.Write(receta)
}

func (r resource) getRecetaPorConsulta(c *routing.Context) error {
	idConsulta, _ := strconv.Atoi(c.Param("idConsulta"))
	receta, err := r.service.GetRecetaPorConsulta(c.Request.Context(), idConsulta)
	if err != nil {
		return err
	}
	return c.Write(receta)
}

func (r resource) archivo(c *routing.Context) error {
	var input RecetaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	var err error
	var ss *excelize.File
	if runtime.GOOS == "windows" {
		ss, err = excelize.OpenFile("./plantillas/Receta.xlsx")
	} else {
		ss, err = excelize.OpenFile("/root/go/src/github.com/JorgeTom0609/veterinaria-server/plantillas/Receta.xlsx")
	}

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
	ss.SetCellValue(sheet, "B8", input.Datos.FechaLlenado.Format("2006-01-02 15:04:05"))

	filaInicial := 11

	for i := 0; i < len(input.Prescripciones); i++ {
		celdaA := "A" + strconv.Itoa((11 + i))
		celdaD := "D" + strconv.Itoa((11 + i))
		celdaE := "E" + strconv.Itoa((11 + i))
		celdaH := "H" + strconv.Itoa((11 + i))
		ss.SetCellValue(sheet, celdaA, input.Prescripciones[i].Producto)
		ss.MergeCell(sheet, celdaA, celdaD)
		ss.SetCellValue(sheet, celdaE, input.Prescripciones[i].Prescripcion)
		ss.MergeCell(sheet, celdaE, celdaH)
		var posibleAltura int = 1
		var letras int = 0
		letras = len(input.Prescripciones[i].Prescripcion)
		var f float64 = float64(letras) / 50.00
		posibleAltura = int(math.Round(f))
		ss.SetRowHeight(sheet, filaInicial, float64(15*posibleAltura))
		filaInicial++
	}

	fileName := fmt.Sprintf("Receta-%s-%s.xlsx", input.Datos.Paciente, input.Datos.FechaLlenado.Format("2006-01-02"))
	if runtime.GOOS == "windows" {
		ss.SaveAs("./resources" + "/" + fileName)
	} else {
		ss.SaveAs("/root/go/src/github.com/JorgeTom0609/veterinaria-server/resources" + "/" + fileName)
	}
	return c.Write(fileName)
}
