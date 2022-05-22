package receta

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for recetas.
type Service interface {
	GetRecetas(ctx context.Context) ([]Receta, error)
	GetRecetaPorId(ctx context.Context, idReceta int) (Receta, error)
	GetRecetaPorConsulta(ctx context.Context, idConsulta int) ([]Receta, error)
	CrearReceta(ctx context.Context, input CreateRecetaRequest) (Receta, error)
	ActualizarReceta(ctx context.Context, input UpdateRecetaRequest) (Receta, error)
}

// Recetas represents the data about an recetas.
type Receta struct {
	entity.Receta
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new recetas service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list recetas.
func (s service) GetRecetas(ctx context.Context) ([]Receta, error) {
	recetas, err := s.repo.GetRecetas(ctx)
	if err != nil {
		return nil, err
	}
	result := []Receta{}
	for _, item := range recetas {
		result = append(result, Receta{item})
	}
	return result, nil
}

// CreateRecetaRequest represents an receta creation request.
type CreateRecetaRequest struct {
	IdProducto   int    `json:"id_producto"`
	IdConsulta   int    `json:"id_consulta"`
	Prescripcion string `json:"prescripcion"`
}

type UpdateRecetaRequest struct {
	IdReceta     int    `json:"id_receta"`
	IdProducto   int    `json:"id_producto"`
	IdConsulta   int    `json:"id_consulta"`
	Prescripcion string `json:"prescripcion"`
}

// Validate validates the UpdateRecetaRequest fields.
func (m UpdateRecetaRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdConsulta, validation.Required),
		validation.Field(&m.IdProducto, validation.Required),
		validation.Field(&m.Prescripcion, validation.Required, validation.Length(0, 1000)),
	)
}

// Validate validates the CreateRecetaRequest fields.
func (m CreateRecetaRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdConsulta, validation.Required),
		validation.Field(&m.IdProducto, validation.Required),
		validation.Field(&m.Prescripcion, validation.Required, validation.Length(0, 1000)),
	)
}

// CrearReceta creates a new receta.
func (s service) CrearReceta(ctx context.Context, req CreateRecetaRequest) (Receta, error) {
	if err := req.Validate(); err != nil {
		return Receta{}, err
	}
	recetaG, err := s.repo.CrearReceta(ctx, entity.Receta{
		IdProducto:   req.IdProducto,
		IdConsulta:   req.IdConsulta,
		Prescripcion: req.Prescripcion,
	})
	if err != nil {
		return Receta{}, err
	}
	return Receta{recetaG}, nil
}

// ActualizarReceta creates a new receta.
func (s service) ActualizarReceta(ctx context.Context, req UpdateRecetaRequest) (Receta, error) {
	if err := req.ValidateUpdate(); err != nil {
		return Receta{}, err
	}
	recetaG, err := s.repo.ActualizarReceta(ctx, entity.Receta{
		IdReceta:     req.IdReceta,
		IdProducto:   req.IdProducto,
		IdConsulta:   req.IdConsulta,
		Prescripcion: req.Prescripcion,
	})
	if err != nil {
		return Receta{}, err
	}
	return Receta{recetaG}, nil
}

// GetRecetaPorId returns the receta with the specified the receta ID.
func (s service) GetRecetaPorId(ctx context.Context, idReceta int) (Receta, error) {
	receta, err := s.repo.GetRecetaPorId(ctx, idReceta)
	if err != nil {
		return Receta{}, err
	}
	return Receta{receta}, nil
}

func (s service) GetRecetaPorConsulta(ctx context.Context, idConsulta int) ([]Receta, error) {
	recetas, err := s.repo.GetRecetaPorConsulta(ctx, idConsulta)
	if err != nil {
		return nil, err
	}
	result := []Receta{}
	for _, item := range recetas {
		result = append(result, Receta{item})
	}
	return result, nil
}
