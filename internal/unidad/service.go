package unidad

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for unidades.
type Service interface {
	GetUnidades(ctx context.Context) ([]Unidad, error)
	GetUnidadPorId(ctx context.Context, idUnidad int) (Unidad, error)
	CrearUnidad(ctx context.Context, input CreateUnidadRequest) (Unidad, error)
	ActualizarUnidad(ctx context.Context, input UpdateUnidadRequest) (Unidad, error)
}

// Unidades represents the data about an unidades.
type Unidad struct {
	entity.Unidad
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new unidades service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list unidades.
func (s service) GetUnidades(ctx context.Context) ([]Unidad, error) {
	unidades, err := s.repo.GetUnidades(ctx)
	if err != nil {
		return nil, err
	}
	result := []Unidad{}
	for _, item := range unidades {
		result = append(result, Unidad{item})
	}
	return result, nil
}

// CreateUnidadRequest represents an unidad creation request.
type CreateUnidadRequest struct {
	IdMedida    int    `json:"id_medida"`
	Descripcion string `json:"descripcion"`
}

type UpdateUnidadRequest struct {
	IdUnidad    int    `json:"id_unidad"`
	IdMedida    int    `json:"id_medida"`
	Descripcion string `json:"descripcion"`
}

// Validate validates the UpdateUnidadRequest fields.
func (m UpdateUnidadRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.IdMedida, validation.Required),
	)
}

// Validate validates the CreateUnidadRequest fields.
func (m CreateUnidadRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.IdMedida, validation.Required),
	)
}

// CrearUnidad creates a new unidad.
func (s service) CrearUnidad(ctx context.Context, req CreateUnidadRequest) (Unidad, error) {
	if err := req.Validate(); err != nil {
		return Unidad{}, err
	}
	unidadG, err := s.repo.CrearUnidad(ctx, entity.Unidad{
		IdMedida:    req.IdMedida,
		Descripcion: req.Descripcion,
	})
	if err != nil {
		return Unidad{}, err
	}
	return Unidad{unidadG}, nil
}

// ActualizarUnidad creates a new unidad.
func (s service) ActualizarUnidad(ctx context.Context, req UpdateUnidadRequest) (Unidad, error) {
	if err := req.ValidateUpdate(); err != nil {
		return Unidad{}, err
	}
	unidadG, err := s.repo.ActualizarUnidad(ctx, entity.Unidad{
		IdUnidad:    req.IdUnidad,
		IdMedida:    req.IdMedida,
		Descripcion: req.Descripcion,
	})
	if err != nil {
		return Unidad{}, err
	}
	return Unidad{unidadG}, nil
}

// GetUnidadPorId returns the unidad with the specified the unidad ID.
func (s service) GetUnidadPorId(ctx context.Context, idUnidad int) (Unidad, error) {
	unidad, err := s.repo.GetUnidadPorId(ctx, idUnidad)
	if err != nil {
		return Unidad{}, err
	}
	return Unidad{unidad}, nil
}
