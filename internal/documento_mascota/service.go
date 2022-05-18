package documento_mascota

import (
	"context"
	"time"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for documentosMascota.
type Service interface {
	GetDocumentosMascota(ctx context.Context) ([]DocumentoMascota, error)
	GetDocumentoMascotaPorId(ctx context.Context, idDocumentoMascota int) (DocumentoMascota, error)
	GetDocumentoMascotaPorMascota(ctx context.Context, idMascota int) ([]DocumentoMascota, error)
	CrearDocumentoMascota(ctx context.Context, input CreateDocumentoMascotaRequest) (DocumentoMascota, error)
	ActualizarDocumentoMascota(ctx context.Context, input UpdateDocumentoMascotaRequest) (DocumentoMascota, error)
}

// DocumentosMascota represents the data about an documentosMascota.
type DocumentoMascota struct {
	entity.DocumentoMascota
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new documentosMascota service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list documentosMascota.
func (s service) GetDocumentosMascota(ctx context.Context) ([]DocumentoMascota, error) {
	documentosMascota, err := s.repo.GetDocumentosMascota(ctx)
	if err != nil {
		return nil, err
	}
	result := []DocumentoMascota{}
	for _, item := range documentosMascota {
		result = append(result, DocumentoMascota{item})
	}
	return result, nil
}

func (s service) GetDocumentoMascotaPorMascota(ctx context.Context, idMascota int) ([]DocumentoMascota, error) {
	documentosMascota, err := s.repo.GetDocumentoMascotaPorMascota(ctx, idMascota)
	if err != nil {
		return nil, err
	}
	result := []DocumentoMascota{}
	for _, item := range documentosMascota {
		result = append(result, DocumentoMascota{item})
	}
	return result, nil
}

// CreateDocumentoMascotaRequest represents an documentoMascota creation request.
type CreateDocumentoMascotaRequest struct {
	IdMascota   int       `json:"id_mascota"`
	IdUsuario   int       `json:"id_usuario"`
	Nombre      string    `json:"nombre"`
	Extension   string    `json:"extension"`
	Ruta        string    `json:"ruta"`
	Descripcion string    `json:"descripcion"`
	Fecha       time.Time `json:"fecha"`
	Base64      string    `json:"base64"`
}

type UpdateDocumentoMascotaRequest struct {
	IdDocumentoMascota int       `json:"id_documento_mascota"`
	IdMascota          int       `json:"id_mascota"`
	IdUsuario          int       `json:"id_usuario"`
	Nombre             string    `json:"nombre"`
	Extension          string    `json:"extension"`
	Ruta               string    `json:"ruta"`
	Descripcion        string    `json:"descripcion"`
	Fecha              time.Time `json:"fecha"`
	Base64             string    `json:"base64"`
}

// Validate validates the UpdateDocumentoMascotaRequest fields.
func (m UpdateDocumentoMascotaRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdMascota, validation.Required),
		validation.Field(&m.IdUsuario, validation.Required),
		validation.Field(&m.Base64, validation.Required),
		validation.Field(&m.Extension, validation.Required),
		validation.Field(&m.Nombre, validation.Required),
	)
}

// Validate validates the CreateDocumentoMascotaRequest fields.
func (m CreateDocumentoMascotaRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdMascota, validation.Required),
		validation.Field(&m.IdUsuario, validation.Required),
		validation.Field(&m.Base64, validation.Required),
		validation.Field(&m.Extension, validation.Required),
		validation.Field(&m.Nombre, validation.Required),
	)
}

// CrearDocumentoMascota creates a new documentoMascota.
func (s service) CrearDocumentoMascota(ctx context.Context, req CreateDocumentoMascotaRequest) (DocumentoMascota, error) {
	if err := req.Validate(); err != nil {
		return DocumentoMascota{}, err
	}
	documentoMascotaG, err := s.repo.CrearDocumentoMascota(ctx, entity.DocumentoMascota{
		IdMascota:   req.IdMascota,
		IdUsuario:   req.IdUsuario,
		Nombre:      req.Nombre,
		Extension:   req.Extension,
		Ruta:        req.Ruta,
		Descripcion: req.Descripcion,
		Fecha:       req.Fecha,
	})
	if err != nil {
		return DocumentoMascota{}, err
	}
	return DocumentoMascota{documentoMascotaG}, nil
}

// ActualizarDocumentoMascota creates a new documentoMascota.
func (s service) ActualizarDocumentoMascota(ctx context.Context, req UpdateDocumentoMascotaRequest) (DocumentoMascota, error) {
	if err := req.ValidateUpdate(); err != nil {
		return DocumentoMascota{}, err
	}
	documentoMascotaG, err := s.repo.ActualizarDocumentoMascota(ctx, entity.DocumentoMascota{
		IdDocumentoMascota: req.IdDocumentoMascota,
		IdMascota:          req.IdMascota,
		IdUsuario:          req.IdUsuario,
		Nombre:             req.Nombre,
		Extension:          req.Extension,
		Ruta:               req.Ruta,
		Descripcion:        req.Descripcion,
		Fecha:              req.Fecha,
	})
	if err != nil {
		return DocumentoMascota{}, err
	}
	return DocumentoMascota{documentoMascotaG}, nil
}

// GetDocumentoMascotaPorId returns the documentoMascota with the specified the documentoMascota ID.
func (s service) GetDocumentoMascotaPorId(ctx context.Context, idDocumentoMascota int) (DocumentoMascota, error) {
	documentoMascota, err := s.repo.GetDocumentoMascotaPorId(ctx, idDocumentoMascota)
	if err != nil {
		return DocumentoMascota{}, err
	}
	return DocumentoMascota{documentoMascota}, nil
}
