package resultado_examen_cualitativo

import (
	"context"
	"database/sql"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for ResultadoDetalleCualitativo.
type Service interface {
	GetResultadoDetalleCualitativoPorId(ctx context.Context, idResultadoDetalleCualitativo int) (ResultadoDetalleCualitativo, error)
	CrearResultadoDetalleCualitativo(ctx context.Context, input CreateResultadoDetalleCualitativoRequest) (ResultadoDetalleCualitativo, error)
}

// ResultadoDetalleCualitativo represents the data about an ResultadoDetalleCualitativo.
type ResultadoDetalleCualitativo struct {
	entity.ResultadoDetalleCualitativo
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new ResultadoDetalleCualitativo service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// CreateResultadoDetalleCualitativoRequest represents an resultadoDetalleCualitativo creation request.
type CreateResultadoDetalleCualitativoRequest struct {
	IdExamenMascota            int          `json:"id_examen_mascota"`
	IdDetalleExamenCualitativo int          `json:"id_detalle_examen_cualitativo"`
	Resultado                  sql.NullBool `json:"resultado"`
}

// Validate validates the CreateResultadoDetalleCualitativoRequest fields.
func (m CreateResultadoDetalleCualitativoRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdDetalleExamenCualitativo, validation.Required),
		validation.Field(&m.IdDetalleExamenCualitativo, validation.Required),
	)
}

// CrearResultadoDetalleCualitativo creates a new resultadoDetalleCualitativo.
func (s service) CrearResultadoDetalleCualitativo(ctx context.Context, req CreateResultadoDetalleCualitativoRequest) (ResultadoDetalleCualitativo, error) {
	if err := req.Validate(); err != nil {
		return ResultadoDetalleCualitativo{}, err
	}
	clienteG, err := s.repo.CrearResultadoDetalleCualitativo(ctx, entity.ResultadoDetalleCualitativo{
		IdExamenMascota:            req.IdExamenMascota,
		IdDetalleExamenCualitativo: req.IdDetalleExamenCualitativo,
		Resultado:                  req.Resultado,
	})
	if err != nil {
		return ResultadoDetalleCualitativo{}, err
	}
	return ResultadoDetalleCualitativo{clienteG}, nil
}

// GetResultadoDetalleCualitativoPorId returns the resultadoDetalleCualitativo with the specified the resultadoDetalleCualitativo ID.
func (s service) GetResultadoDetalleCualitativoPorId(ctx context.Context, idResultadoDetalleCualitativo int) (ResultadoDetalleCualitativo, error) {
	resultadoDetalleCualitativo, err := s.repo.GetResultadoDetalleCualitativoPorId(ctx, idResultadoDetalleCualitativo)
	if err != nil {
		return ResultadoDetalleCualitativo{}, err
	}
	return ResultadoDetalleCualitativo{resultadoDetalleCualitativo}, nil
}
