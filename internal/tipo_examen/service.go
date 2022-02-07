package tipo_examen

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Service interface {
	GetTipoExamenes(ctx context.Context) ([]TipoExamen, error)
	GetTipoExamenPorId(ctx context.Context, idTipoExamen int) (TipoExamen, error)
	CrearTipoExamen(ctx context.Context, input CreateTipoExamenRequest) (TipoExamen, error)
	ActualizarTipoExamen(ctx context.Context, input UpdateTipoExamenRequest) (TipoExamen, error)
}

type TipoExamen struct {
	entity.TipoExamen
}

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

func (s service) GetTipoExamenes(ctx context.Context) ([]TipoExamen, error) {
	tipoExamenes, err := s.repo.GetTipoExamenes(ctx)
	if err != nil {
		return nil, err
	}
	result := []TipoExamen{}
	for _, item := range tipoExamenes {
		result = append(result, TipoExamen{item})
	}
	return result, nil
}

type CreateTipoExamenRequest struct {
	IdEspecie   int     `json:"id_especie"`
	Descripcion *string `json:"descripcion"`
	Muestra     *string `json:"muestra"`
}

type UpdateTipoExamenRequest struct {
	IdTipoExamen int     `json:"id_tipo_examen"`
	IdEspecie    int     `json:"id_especie"`
	Descripcion  *string `json:"descripcion"`
	Muestra      *string `json:"muestra"`
}

func (m UpdateTipoExamenRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdEspecie, validation.Required),
	)
}

func (m CreateTipoExamenRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdEspecie, validation.Required),
	)
}

func (s service) CrearTipoExamen(ctx context.Context, req CreateTipoExamenRequest) (TipoExamen, error) {
	if err := req.Validate(); err != nil {
		return TipoExamen{}, err
	}
	tipoExamenG, err := s.repo.CrearTipoExamen(ctx, entity.TipoExamen{
		IdEspecie:   req.IdEspecie,
		Descripcion: req.Descripcion,
		Muestra:     req.Muestra,
	})
	if err != nil {
		return TipoExamen{}, err
	}
	return TipoExamen{tipoExamenG}, nil
}

func (s service) ActualizarTipoExamen(ctx context.Context, req UpdateTipoExamenRequest) (TipoExamen, error) {
	if err := req.ValidateUpdate(); err != nil {
		return TipoExamen{}, err
	}
	tipoExamenG, err := s.repo.ActualizarTipoExamen(ctx, entity.TipoExamen{
		IdTipoExamen: req.IdTipoExamen,
		IdEspecie:    req.IdEspecie,
		Descripcion:  req.Descripcion,
		Muestra:      req.Muestra,
	})
	if err != nil {
		return TipoExamen{}, err
	}
	return TipoExamen{tipoExamenG}, nil
}

func (s service) GetTipoExamenPorId(ctx context.Context, idTipoExamen int) (TipoExamen, error) {
	tipoExamen, err := s.repo.GetTipoExamenPorId(ctx, idTipoExamen)
	if err != nil {
		return TipoExamen{}, err
	}
	return TipoExamen{tipoExamen}, nil
}
