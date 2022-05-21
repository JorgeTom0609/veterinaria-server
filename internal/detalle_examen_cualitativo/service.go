package detalle_examen_cualitativo

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for detallesExamenCualitativo.
type Service interface {
	GetDetallesExamenCualitativo(ctx context.Context) ([]DetallesExamenCualitativo, error)
	GetDetalleExamenCualitativoPorId(ctx context.Context, idDetalleExamenCualitativo int) (DetallesExamenCualitativo, error)
	GetDetallesExamenCualitativoPorTipoExamen(ctx context.Context, idTipoDeExamen int) ([]DetallesExamenCualitativo, error)
	CrearDetalleExamenCualitativo(ctx context.Context, input CreateDetalleExamenCualitativoRequest) (DetallesExamenCualitativo, error)
	ActualizarDetalleExamenCualitativo(ctx context.Context, input UpdateDetalleExamenCualitativoRequest) (DetallesExamenCualitativo, error)
}

// detallesExamenCualitativo represents the data about an detallesExamenCualitativo.
type DetallesExamenCualitativo struct {
	entity.DetallesExamenCualitativo
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new detallesExamenCualitativo service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list detallesExamenCualitativo.
func (s service) GetDetallesExamenCualitativo(ctx context.Context) ([]DetallesExamenCualitativo, error) {
	detallesExamenCualitativo, err := s.repo.GetDetallesExamenCualitativo(ctx)
	if err != nil {
		return nil, err
	}
	result := []DetallesExamenCualitativo{}
	for _, item := range detallesExamenCualitativo {
		result = append(result, DetallesExamenCualitativo{item})
	}
	return result, nil
}

// CreateDetalleExamenCualitativoRequest represents an detalleExamenCualitativo creation request.
type CreateDetalleExamenCualitativoRequest struct {
	IdTipoExamen int    `json:"id_tipo_examen"`
	Parametro    string `json:"parametro"`
}

type UpdateDetalleExamenCualitativoRequest struct {
	IdDetalleExamenCualitativo int    `json:"id_detalle_examen_cualitativo"`
	IdTipoExamen               int    `json:"id_tipo_examen"`
	Parametro                  string `json:"parametro"`
}

// Validate validates the UpdateDetalleExamenCualitativoRequest fields.
func (m UpdateDetalleExamenCualitativoRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdTipoExamen, validation.Required),
		validation.Field(&m.Parametro, validation.Required, validation.Length(0, 1000)),
	)
}

// Validate validates the CreateDetalleExamenCualitativoRequest fields.
func (m CreateDetalleExamenCualitativoRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdTipoExamen, validation.Required),
		validation.Field(&m.Parametro, validation.Required, validation.Length(0, 1000)),
	)
}

// CrearDetalleExamenCualitativo creates a new detalleExamenCualitativo.
func (s service) CrearDetalleExamenCualitativo(ctx context.Context, req CreateDetalleExamenCualitativoRequest) (DetallesExamenCualitativo, error) {
	if err := req.Validate(); err != nil {
		return DetallesExamenCualitativo{}, err
	}
	detalleExamenCualitativoG, err := s.repo.CrearDetalleExamenCualitativo(ctx, entity.DetallesExamenCualitativo{
		IdTipoExamen: req.IdTipoExamen,
		Parametro:    req.Parametro,
	})
	if err != nil {
		return DetallesExamenCualitativo{}, err
	}
	return DetallesExamenCualitativo{detalleExamenCualitativoG}, nil
}

// ActualizarDetalleExamenCualitativo creates a new detalleExamenCualitativo.
func (s service) ActualizarDetalleExamenCualitativo(ctx context.Context, req UpdateDetalleExamenCualitativoRequest) (DetallesExamenCualitativo, error) {
	if err := req.ValidateUpdate(); err != nil {
		return DetallesExamenCualitativo{}, err
	}
	detalleExamenCualitativoG, err := s.repo.ActualizarDetalleExamenCualitativo(ctx, entity.DetallesExamenCualitativo{
		IdDetalleExamenCualitativo: req.IdDetalleExamenCualitativo,
		IdTipoExamen:               req.IdTipoExamen,
		Parametro:                  req.Parametro,
	})
	if err != nil {
		return DetallesExamenCualitativo{}, err
	}
	return DetallesExamenCualitativo{detalleExamenCualitativoG}, nil
}

// GetDetalleExamenCualitativoPorId returns the detalleExamenCualitativo with the specified the detalleExamenCualitativo ID.
func (s service) GetDetalleExamenCualitativoPorId(ctx context.Context, idDetalleExamenCualitativo int) (DetallesExamenCualitativo, error) {
	detalleExamenCualitativo, err := s.repo.GetDetalleExamenCualitativoPorId(ctx, idDetalleExamenCualitativo)
	if err != nil {
		return DetallesExamenCualitativo{}, err
	}
	return DetallesExamenCualitativo{detalleExamenCualitativo}, nil
}

func (s service) GetDetallesExamenCualitativoPorTipoExamen(ctx context.Context, idTipoDeExamen int) ([]DetallesExamenCualitativo, error) {
	detallesExamenCualitativo, err := s.repo.GetDetallesExamenCualitativoPorTipoExamen(ctx, idTipoDeExamen)
	if err != nil {
		return nil, err
	}
	result := []DetallesExamenCualitativo{}
	for _, item := range detallesExamenCualitativo {
		result = append(result, DetallesExamenCualitativo{item})
	}
	return result, nil
}
