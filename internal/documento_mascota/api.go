package documento_mascota

import (
	"encoding/base64"
	"net/http"
	"os"
	"strconv"
	"veterinaria-server/internal/errors"
	"veterinaria-server/pkg/log"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}
	r.Use(authHandler)
	// the following endpoints require a valid JWT
	r.Get("/documentosMascota", res.getDocumentosMascota)
	r.Get("/documentosMascota/<idDocumentoMascota>", res.getDocumentoMascotaPorId)
	r.Get("/documentosMascota/porMascota/<idMascota>", res.getDocumentoMascotaPorMascota)
	r.Post("/documentosMascota", res.crearDocumentoMascota)
	r.Put("/documentosMascota", res.actualizarDocumentoMascota)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) getDocumentosMascota(c *routing.Context) error {
	documentosMascota, err := r.service.GetDocumentosMascota(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(documentosMascota)
}

func (r resource) getDocumentoMascotaPorMascota(c *routing.Context) error {
	idMascota, _ := strconv.Atoi(c.Param("idMascota"))
	documentosMascota, err := r.service.GetDocumentoMascotaPorMascota(c.Request.Context(), idMascota)
	if err != nil {
		return err
	}
	return c.Write(documentosMascota)
}

func (r resource) crearDocumentoMascota(c *routing.Context) error {
	var input CreateDocumentoMascotaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	documentoMascota, err := r.service.CrearDocumentoMascota(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(documentoMascota, http.StatusCreated)
}

func (r resource) actualizarDocumentoMascota(c *routing.Context) error {
	var input UpdateDocumentoMascotaRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	dec, err := base64.StdEncoding.DecodeString(input.Base64)
	if err != nil {
		panic(err)
	}

	input.Nombre = input.Nombre + " - " + input.Fecha.Format("2006-01-02")
	input.Ruta = "./documentos-mascota/"

	f, err := os.Create("./documentos-mascota/" + input.Nombre + "." + input.Extension)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}

	documentoMascota, err := r.service.ActualizarDocumentoMascota(c.Request.Context(), input)
	if err != nil {
		return err
	}

	/*
		f.Seek(0, 0)
		io.Copy(os.Stdout, f)
	*/

	return c.WriteWithStatus(documentoMascota, http.StatusCreated)
}

func (r resource) getDocumentoMascotaPorId(c *routing.Context) error {
	idDocumentoMascota, _ := strconv.Atoi(c.Param("idDocumentoMascota"))
	documentoMascota, err := r.service.GetDocumentoMascotaPorId(c.Request.Context(), idDocumentoMascota)
	if err != nil {
		return err
	}
	return c.Write(documentoMascota)
}
