package detalle_examen_informativo

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for detallesExamenInformativo.
type Service interface {
	GetDetallesExamenInformativo(ctx context.Context) ([]DetallesExamenInformativo, error)
	GetDetalleExamenInformativoPorId(ctx context.Context, idDetalleExamenInformativo int) (DetallesExamenInformativo, error)
	CrearDetalleExamenInformativo(ctx context.Context, input CreateDetalleExamenInformativoRequest) (DetallesExamenInformativo, error)
	ActualizarDetalleExamenInformativo(ctx context.Context, input UpdateDetalleExamenInformativoRequest) (DetallesExamenInformativo, error)
}

// DetallesExamenInformativo represents the data about an detallesExamenInformativo.
type DetallesExamenInformativo struct {
	entity.DetallesExamenInformativo
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new detallesExamenInformativo service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list detallesExamenInformativo.
func (s service) GetDetallesExamenInformativo(ctx context.Context) ([]DetallesExamenInformativo, error) {
	detallesExamenInformativo, err := s.repo.GetDetallesExamenInformativo(ctx)
	if err != nil {
		return nil, err
	}
	result := []DetallesExamenInformativo{}
	for _, item := range detallesExamenInformativo {
		result = append(result, DetallesExamenInformativo{item})
	}
	return result, nil
}

// CreateDetalleExamenInformativoRequest represents an detalleExamenInformativo creation request.
type CreateDetalleExamenInformativoRequest struct {
	IdTipoExamen int    `json:"id_tipo_examen"`
	Parametro    string `json:"parametro"`
}

type UpdateDetalleExamenInformativoRequest struct {
	IdDetalleExamenInformativo int    `json:"id_detalle_examen_informativo"`
	IdTipoExamen               int    `json:"id_tipo_examen"`
	Parametro                  string `json:"parametro"`
}

// Validate validates the UpdateDetalleExamenInformativoRequest fields.
func (m UpdateDetalleExamenInformativoRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdTipoExamen, validation.Required),
		validation.Field(&m.Parametro, validation.Required, validation.Length(0, 128)),
	)
}

// Validate validates the CreateDetalleExamenInformativoRequest fields.
func (m CreateDetalleExamenInformativoRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdTipoExamen, validation.Required),
		validation.Field(&m.Parametro, validation.Required, validation.Length(0, 128)),
	)
}

// CrearDetalleExamenInformativo creates a new detalleExamenInformativo.
func (s service) CrearDetalleExamenInformativo(ctx context.Context, req CreateDetalleExamenInformativoRequest) (DetallesExamenInformativo, error) {
	if err := req.Validate(); err != nil {
		return DetallesExamenInformativo{}, err
	}
	detalleExamenInformativoG, err := s.repo.CrearDetalleExamenInformativo(ctx, entity.DetallesExamenInformativo{
		IdTipoExamen: req.IdTipoExamen,
		Parametro:    req.Parametro,
	})
	if err != nil {
		return DetallesExamenInformativo{}, err
	}
	return DetallesExamenInformativo{detalleExamenInformativoG}, nil
}

// ActualizarDetalleExamenInformativo creates a new detalleExamenInformativo.
func (s service) ActualizarDetalleExamenInformativo(ctx context.Context, req UpdateDetalleExamenInformativoRequest) (DetallesExamenInformativo, error) {
	if err := req.ValidateUpdate(); err != nil {
		return DetallesExamenInformativo{}, err
	}
	detalleExamenInformativoG, err := s.repo.ActualizarDetalleExamenInformativo(ctx, entity.DetallesExamenInformativo{
		IdDetalleExamenInformativo: req.IdDetalleExamenInformativo,
		IdTipoExamen:               req.IdTipoExamen,
		Parametro:                  req.Parametro,
	})
	if err != nil {
		return DetallesExamenInformativo{}, err
	}
	return DetallesExamenInformativo{detalleExamenInformativoG}, nil
}

// GetDetalleExamenInformativoPorId returns the detalleExamenInformativo with the specified the detalleExamenInformativo ID.
func (s service) GetDetalleExamenInformativoPorId(ctx context.Context, idDetalleExamenInformativo int) (DetallesExamenInformativo, error) {
	detalleExamenInformativo, err := s.repo.GetDetalleExamenInformativoPorId(ctx, idDetalleExamenInformativo)
	if err != nil {
		return DetallesExamenInformativo{}, err
	}
	return DetallesExamenInformativo{detalleExamenInformativo}, nil
}
