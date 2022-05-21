package medida

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for medidas.
type Service interface {
	GetMedidas(ctx context.Context) ([]Medida, error)
	GetMedidaPorId(ctx context.Context, idMedida int) (Medida, error)
	CrearMedida(ctx context.Context, input CreateMedidaRequest) (Medida, error)
	ActualizarMedida(ctx context.Context, input UpdateMedidaRequest) (Medida, error)
}

// Medidas represents the data about an medidas.
type Medida struct {
	entity.Medida
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new medidas service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list medidas.
func (s service) GetMedidas(ctx context.Context) ([]Medida, error) {
	medidas, err := s.repo.GetMedidas(ctx)
	if err != nil {
		return nil, err
	}
	result := []Medida{}
	for _, item := range medidas {
		result = append(result, Medida{item})
	}
	return result, nil
}

// CreateMedidaRequest represents an medida creation request.
type CreateMedidaRequest struct {
	Descripcion string `json:"descripcion"`
}

type UpdateMedidaRequest struct {
	IdMedida    int    `json:"id_medida"`
	Descripcion string `json:"descripcion"`
}

// Validate validates the UpdateMedidaRequest fields.
func (m UpdateMedidaRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 1000)),
	)
}

// Validate validates the CreateMedidaRequest fields.
func (m CreateMedidaRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 1000)),
	)
}

// CrearMedida creates a new medida.
func (s service) CrearMedida(ctx context.Context, req CreateMedidaRequest) (Medida, error) {
	if err := req.Validate(); err != nil {
		return Medida{}, err
	}
	medidaG, err := s.repo.CrearMedida(ctx, entity.Medida{
		Descripcion: req.Descripcion,
	})
	if err != nil {
		return Medida{}, err
	}
	return Medida{medidaG}, nil
}

// ActualizarMedida creates a new medida.
func (s service) ActualizarMedida(ctx context.Context, req UpdateMedidaRequest) (Medida, error) {
	if err := req.ValidateUpdate(); err != nil {
		return Medida{}, err
	}
	medidaG, err := s.repo.ActualizarMedida(ctx, entity.Medida{
		IdMedida:    req.IdMedida,
		Descripcion: req.Descripcion,
	})
	if err != nil {
		return Medida{}, err
	}
	return Medida{medidaG}, nil
}

// GetMedidaPorId returns the medida with the specified the medida ID.
func (s service) GetMedidaPorId(ctx context.Context, idMedida int) (Medida, error) {
	medida, err := s.repo.GetMedidaPorId(ctx, idMedida)
	if err != nil {
		return Medida{}, err
	}
	return Medida{medida}, nil
}
