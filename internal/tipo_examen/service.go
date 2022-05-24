package tipo_examen

import (
	"context"
	"veterinaria-server/internal/detalle_examen_cualitativo"
	"veterinaria-server/internal/detalle_examen_cuantitativo"
	"veterinaria-server/internal/detalle_examen_informativo"
	"veterinaria-server/internal/entity"
	"veterinaria-server/internal/resultado_examen_cualitativo"
	"veterinaria-server/internal/resultado_examen_cuantitativo"
	"veterinaria-server/internal/resultado_examen_informativo"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Service interface {
	GetTipoExamenes(ctx context.Context) ([]TipoExamen, error)
	GetTipoExamenPorEspecie(ctx context.Context, idEspecie int) ([]TipoExamen, error)
	GetTipoExamenPorEspecieDisponibles(ctx context.Context, idEspecie int, idMascota int) ([]TipoExamen, error)
	GetDetallesExamenPorTipoExamen(ctx context.Context, idTipoExamen int) (DetallesExamen, error)
	GetTipoExamenPorId(ctx context.Context, idTipoExamen int) (TipoExamen, error)
	CrearTipoExamen(ctx context.Context, input CreateTipoExamenRequest) (TipoExamen, error)
	ActualizarTipoExamen(ctx context.Context, input UpdateTipoExamenRequest) (TipoExamen, error)
}

type TipoExamen struct {
	entity.TipoExamen
}

type DetallesExamen struct {
	DetallesExamenCualitativo  []entity.DetallesExamenCualitativo  `json:"cualitativos"`
	DetallesExamenCuantitativo []entity.DetallesExamenCuantitativo `json:"cuantitativos"`
	DetallesExamenInformativo  []entity.DetallesExamenInformativo  `json:"informativos"`
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

func (s service) GetTipoExamenPorEspecie(ctx context.Context, idEspecie int) ([]TipoExamen, error) {
	tipoExamenes, err := s.repo.GetTipoExamenPorEspecie(ctx, idEspecie)
	if err != nil {
		return nil, err
	}
	result := []TipoExamen{}
	for _, item := range tipoExamenes {
		result = append(result, TipoExamen{item})
	}
	return result, nil
}

func (s service) GetTipoExamenPorEspecieDisponibles(ctx context.Context, idEspecie int, idMascota int) ([]TipoExamen, error) {
	tipoExamenes, err := s.repo.GetTipoExamenPorEspecieDisponibles(ctx, idEspecie, idMascota)
	if err != nil {
		return nil, err
	}
	result := []TipoExamen{}
	for _, item := range tipoExamenes {
		result = append(result, TipoExamen{item})
	}
	return result, nil
}

func (s service) GetDetallesExamenPorTipoExamen(ctx context.Context, idTipoExamen int) (DetallesExamen, error) {
	tipoExamenes, err := s.repo.GetDetallesExamenPorTipoExamen(ctx, idTipoExamen)
	if err != nil {
		return DetallesExamen{}, err
	}
	return tipoExamenes, nil
}

type CreateTipoExamenRequest struct {
	IdEspecie   int     `json:"id_especie"`
	Descripcion string  `json:"descripcion"`
	Titulo      string  `json:"titulo"`
	Muestra     string  `json:"muestra"`
	Valor       float32 `json:"valor"`
}

type UpdateTipoExamenRequest struct {
	IdTipoExamen int     `json:"id_tipo_examen"`
	IdEspecie    int     `json:"id_especie"`
	Titulo       string  `json:"titulo"`
	Muestra      string  `json:"muestra"`
	Descripcion  string  `json:"descripcion"`
	Valor        float32 `json:"valor"`
}

type UpdateTipoExamenConDetallesRequest struct {
	TipoExamen    UpdateTipoExamenRequest                                              `json:"tipoDeExamen"`
	Cualitativos  []detalle_examen_cualitativo.UpdateDetalleExamenCualitativoRequest   `json:"cualitativos"`
	Cuantitativos []detalle_examen_cuantitativo.UpdateDetalleExamenCuantitativoRequest `json:"cuantitativos"`
	Informativos  []detalle_examen_informativo.UpdateDetalleExamenInformativoRequest   `json:"informativos"`
}

type CreateResultadosRequest struct {
	Cualitativos    []resultado_examen_cualitativo.CreateResultadoDetalleCualitativoRequest   `json:"cualitativos"`
	Cuantitativos   []resultado_examen_cuantitativo.CreateResultadoDetalleCuantitativoRequest `json:"cuantitativos"`
	Informativos    []resultado_examen_informativo.CreateResultadoDetalleInformativoRequest   `json:"informativos"`
	IdExamenMascota int                                                                       `json:"id_examen_mascota"`
}

func (m UpdateTipoExamenRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdEspecie, validation.Required),
		validation.Field(&m.Titulo, validation.Required),
		validation.Field(&m.Muestra, validation.Required),
	)
}

func (m CreateTipoExamenRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdEspecie, validation.Required),
		validation.Field(&m.Titulo, validation.Required),
		validation.Field(&m.Muestra, validation.Required),
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
		Titulo:      req.Titulo,
		Valor:       req.Valor,
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
		Titulo:       req.Titulo,
		Valor:        req.Valor,
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
