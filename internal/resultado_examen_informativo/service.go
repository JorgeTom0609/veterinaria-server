package resultado_examen_informativo

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for ResultadoDetalleInformativo.
type Service interface {
	GetResultadoDetalleInformativoPorId(ctx context.Context, idResultadoDetalleInformativo int) (ResultadoDetalleInformativo, error)
	CrearResultadoDetalleInformativo(ctx context.Context, input CreateResultadoDetalleInformativoRequest) (ResultadoDetalleInformativo, error)
}

// ResultadoDetalleInformativo represents the data about an ResultadoDetalleInformativo.
type ResultadoDetalleInformativo struct {
	entity.ResultadoDetalleInformativo
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new ResultadoDetalleInformativo service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// CreateResultadoDetalleInformativoRequest represents an resultadoDetalleInformativo creation request.
type CreateResultadoDetalleInformativoRequest struct {
	IdExamenMascota            int    `json:"id_examen_mascota"`
	IdDetalleExamenInformativo int    `json:"id_detalle_examen_Informativo"`
	Resultado                  string `json:"resultado"`
}

// Validate validates the CreateResultadoDetalleInformativoRequest fields.
func (m CreateResultadoDetalleInformativoRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdDetalleExamenInformativo, validation.Required),
		validation.Field(&m.IdDetalleExamenInformativo, validation.Required),
	)
}

// CrearResultadoDetalleInformativo creates a new resultadoDetalleInformativo.
func (s service) CrearResultadoDetalleInformativo(ctx context.Context, req CreateResultadoDetalleInformativoRequest) (ResultadoDetalleInformativo, error) {
	if err := req.Validate(); err != nil {
		return ResultadoDetalleInformativo{}, err
	}
	clienteG, err := s.repo.CrearResultadoDetalleInformativo(ctx, entity.ResultadoDetalleInformativo{
		IdExamenMascota:            req.IdExamenMascota,
		IdDetalleExamenInformativo: req.IdDetalleExamenInformativo,
		Resultado:                  req.Resultado,
	})
	if err != nil {
		return ResultadoDetalleInformativo{}, err
	}
	return ResultadoDetalleInformativo{clienteG}, nil
}

// GetResultadoDetalleInformativoPorId returns the resultadoDetalleInformativo with the specified the resultadoDetalleInformativo ID.
func (s service) GetResultadoDetalleInformativoPorId(ctx context.Context, idResultadoDetalleInformativo int) (ResultadoDetalleInformativo, error) {
	resultadoDetalleInformativo, err := s.repo.GetResultadoDetalleInformativoPorId(ctx, idResultadoDetalleInformativo)
	if err != nil {
		return ResultadoDetalleInformativo{}, err
	}
	return ResultadoDetalleInformativo{resultadoDetalleInformativo}, nil
}
