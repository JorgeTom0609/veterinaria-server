package detalle_examen_cuantitativo

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for detallesExamenCuantitativo.
type Service interface {
	GetDetallesExamenCuantitativo(ctx context.Context) ([]DetallesExamenCuantitativo, error)
	GetDetallesExamenCuantitativoPorTipoExamen(ctx context.Context, idTipoDeExamen int) ([]DetallesExamenCuantitativo, error)
	GetDetalleExamenCuantitativoPorId(ctx context.Context, idDetalleExamenCuantitativo int) (DetallesExamenCuantitativo, error)
	CrearDetalleExamenCuantitativo(ctx context.Context, input CreateDetalleExamenCuantitativoRequest) (DetallesExamenCuantitativo, error)
	ActualizarDetalleExamenCuantitativo(ctx context.Context, input UpdateDetalleExamenCuantitativoRequest) (DetallesExamenCuantitativo, error)
}

// DetallesExamenCuantitativo represents the data about an detallesExamenCuantitativo.
type DetallesExamenCuantitativo struct {
	entity.DetallesExamenCuantitativo
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new detallesExamenCuantitativo service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list detallesExamenCuantitativo.
func (s service) GetDetallesExamenCuantitativo(ctx context.Context) ([]DetallesExamenCuantitativo, error) {
	detallesExamenCuantitativo, err := s.repo.GetDetallesExamenCuantitativo(ctx)
	if err != nil {
		return nil, err
	}
	result := []DetallesExamenCuantitativo{}
	for _, item := range detallesExamenCuantitativo {
		result = append(result, DetallesExamenCuantitativo{item})
	}
	return result, nil
}

// CreateDetalleExamenCuantitativoRequest represents an detalleExamenCuantitativo creation request.
type CreateDetalleExamenCuantitativoRequest struct {
	IdTipoExamen           int     `json:"id_tipo_examen"`
	Parametro              string  `json:"parametro"`
	RangoReferenciaInicial float32 `json:"rango_referencia_inicial"`
	RangoReferenciaFinal   float32 `json:"rango_referencia_final"`
	Unidad                 *string `json:"unidad"`
	AlertaMenor            *string `json:"alerta_menor"`
	AlertaRango            *string `json:"alerta_rango"`
	AlertaMayor            *string `json:"alerta_mayor"`
}

type UpdateDetalleExamenCuantitativoRequest struct {
	IdDetalleExamenCuantitativo int     `json:"id_detalle_examen_cuantitativo"`
	IdTipoExamen                int     `json:"id_tipo_examen"`
	Parametro                   string  `json:"parametro"`
	RangoReferenciaInicial      float32 `json:"rango_referencia_inicial"`
	RangoReferenciaFinal        float32 `json:"rango_referencia_final"`
	Unidad                      *string `json:"unidad"`
	AlertaMenor                 *string `json:"alerta_menor"`
	AlertaRango                 *string `json:"alerta_rango"`
	AlertaMayor                 *string `json:"alerta_mayor"`
}

// Validate validates the UpdateDetalleExamenCuantitativoRequest fields.
func (m UpdateDetalleExamenCuantitativoRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdTipoExamen, validation.Required),
		validation.Field(&m.Parametro, validation.Required, validation.Length(0, 1000)),
	)
}

// Validate validates the CreateDetalleExamenCuantitativoRequest fields.
func (m CreateDetalleExamenCuantitativoRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdTipoExamen, validation.Required),
		validation.Field(&m.Parametro, validation.Required, validation.Length(0, 1000)),
	)
}

// CrearDetalleExamenCuantitativo creates a new detalleExamenCuantitativo.
func (s service) CrearDetalleExamenCuantitativo(ctx context.Context, req CreateDetalleExamenCuantitativoRequest) (DetallesExamenCuantitativo, error) {
	if err := req.Validate(); err != nil {
		return DetallesExamenCuantitativo{}, err
	}
	detalleExamenCuantitativoG, err := s.repo.CrearDetalleExamenCuantitativo(ctx, entity.DetallesExamenCuantitativo{
		IdTipoExamen:           req.IdTipoExamen,
		Parametro:              req.Parametro,
		RangoReferenciaInicial: req.RangoReferenciaInicial,
		RangoReferenciaFinal:   req.RangoReferenciaFinal,
		Unidad:                 req.Unidad,
		AlertaMenor:            req.AlertaMenor,
		AlertaRango:            req.AlertaRango,
		AlertaMayor:            req.AlertaMayor,
	})
	if err != nil {
		return DetallesExamenCuantitativo{}, err
	}
	return DetallesExamenCuantitativo{detalleExamenCuantitativoG}, nil
}

// ActualizarDetalleExamenCuantitativo creates a new detalleExamenCuantitativo.
func (s service) ActualizarDetalleExamenCuantitativo(ctx context.Context, req UpdateDetalleExamenCuantitativoRequest) (DetallesExamenCuantitativo, error) {
	if err := req.ValidateUpdate(); err != nil {
		return DetallesExamenCuantitativo{}, err
	}
	detalleExamenCuantitativoG, err := s.repo.ActualizarDetalleExamenCuantitativo(ctx, entity.DetallesExamenCuantitativo{
		IdDetalleExamenCuantitativo: req.IdDetalleExamenCuantitativo,
		IdTipoExamen:                req.IdTipoExamen,
		Parametro:                   req.Parametro,
		RangoReferenciaInicial:      req.RangoReferenciaInicial,
		RangoReferenciaFinal:        req.RangoReferenciaFinal,
		Unidad:                      req.Unidad,
		AlertaMenor:                 req.AlertaMenor,
		AlertaRango:                 req.AlertaRango,
		AlertaMayor:                 req.AlertaMayor,
	})
	if err != nil {
		return DetallesExamenCuantitativo{}, err
	}
	return DetallesExamenCuantitativo{detalleExamenCuantitativoG}, nil
}

// GetDetalleExamenCuantitativoPorId returns the detalleExamenCuantitativo with the specified the detalleExamenCuantitativo ID.
func (s service) GetDetalleExamenCuantitativoPorId(ctx context.Context, idDetalleExamenCuantitativo int) (DetallesExamenCuantitativo, error) {
	detalleExamenCuantitativo, err := s.repo.GetDetalleExamenCuantitativoPorId(ctx, idDetalleExamenCuantitativo)
	if err != nil {
		return DetallesExamenCuantitativo{}, err
	}
	return DetallesExamenCuantitativo{detalleExamenCuantitativo}, nil
}

func (s service) GetDetallesExamenCuantitativoPorTipoExamen(ctx context.Context, idTipoDeExamen int) ([]DetallesExamenCuantitativo, error) {
	detallesExamenCuantitativo, err := s.repo.GetDetallesExamenCuantitativoPorTipoExamen(ctx, idTipoDeExamen)
	if err != nil {
		return nil, err
	}
	result := []DetallesExamenCuantitativo{}
	for _, item := range detallesExamenCuantitativo {
		result = append(result, DetallesExamenCuantitativo{item})
	}
	return result, nil
}
