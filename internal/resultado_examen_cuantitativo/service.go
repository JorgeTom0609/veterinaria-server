package resultado_examen_cuantitativo

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for ResultadoDetalleCuantitativo.
type Service interface {
	GetResultadoDetalleCuantitativoPorId(ctx context.Context, idResultadoDetalleCuantitativo int) (ResultadoDetalleCuantitativo, error)
	CrearResultadoDetalleCuantitativo(ctx context.Context, input CreateResultadoDetalleCuantitativoRequest) (ResultadoDetalleCuantitativo, error)
}

// ResultadoDetalleCuantitativo represents the data about an ResultadoDetalleCuantitativo.
type ResultadoDetalleCuantitativo struct {
	entity.ResultadoDetalleCuantitativo
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new ResultadoDetalleCuantitativo service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// CreateResultadoDetalleCuantitativoRequest represents an resultadoDetalleCuantitativo creation request.
type CreateResultadoDetalleCuantitativoRequest struct {
	IdExamenMascota             int     `json:"id_examen_mascota"`
	IdDetalleExamenCuantitativo int     `json:"id_detalle_examen_cuantitativo"`
	Resultado                   float32 `json:"resultado"`
}

// Validate validates the CreateResultadoDetalleCuantitativoRequest fields.
func (m CreateResultadoDetalleCuantitativoRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdDetalleExamenCuantitativo, validation.Required),
		validation.Field(&m.IdDetalleExamenCuantitativo, validation.Required),
	)
}

// CrearResultadoDetalleCuantitativo creates a new resultadoDetalleCuantitativo.
func (s service) CrearResultadoDetalleCuantitativo(ctx context.Context, req CreateResultadoDetalleCuantitativoRequest) (ResultadoDetalleCuantitativo, error) {
	if err := req.Validate(); err != nil {
		return ResultadoDetalleCuantitativo{}, err
	}
	clienteG, err := s.repo.CrearResultadoDetalleCuantitativo(ctx, entity.ResultadoDetalleCuantitativo{
		IdExamenMascota:             req.IdExamenMascota,
		IdDetalleExamenCuantitativo: req.IdDetalleExamenCuantitativo,
		Resultado:                   req.Resultado,
	})
	if err != nil {
		return ResultadoDetalleCuantitativo{}, err
	}
	return ResultadoDetalleCuantitativo{clienteG}, nil
}

// GetResultadoDetalleCuantitativoPorId returns the resultadoDetalleCuantitativo with the specified the resultadoDetalleCuantitativo ID.
func (s service) GetResultadoDetalleCuantitativoPorId(ctx context.Context, idResultadoDetalleCuantitativo int) (ResultadoDetalleCuantitativo, error) {
	resultadoDetalleCuantitativo, err := s.repo.GetResultadoDetalleCuantitativoPorId(ctx, idResultadoDetalleCuantitativo)
	if err != nil {
		return ResultadoDetalleCuantitativo{}, err
	}
	return ResultadoDetalleCuantitativo{resultadoDetalleCuantitativo}, nil
}
