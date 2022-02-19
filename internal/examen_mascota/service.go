package examen_mascota

import (
	"context"
	"time"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for examenesMascota.
type Service interface {
	GetExamenesMascota(ctx context.Context) ([]ExamenMascota, error)
	GetExamenesMascotaPorMascotayEstado(ctx context.Context, idExamenMascota int, estado string) ([]ExamenMascota, error)
	GetExamenMascotaPorId(ctx context.Context, idExamenMascota int) (ExamenMascota, error)
	CrearExamenMascota(ctx context.Context, input CreateExamenMascotaRequest) (ExamenMascota, error)
	ActualizarExamenMascota(ctx context.Context, input UpdateExamenMascotaRequest) (ExamenMascota, error)
}

// ExamenesMascota represents the data about an examenesMascota.
type ExamenMascota struct {
	entity.ExamenMascota
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new examenesMascota service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list examenesMascota.
func (s service) GetExamenesMascota(ctx context.Context) ([]ExamenMascota, error) {
	examenesMascota, err := s.repo.GetExamenesMascota(ctx)
	if err != nil {
		return nil, err
	}
	result := []ExamenMascota{}
	for _, item := range examenesMascota {
		result = append(result, ExamenMascota{item})
	}
	return result, nil
}

func (s service) GetExamenesMascotaPorMascotayEstado(ctx context.Context, idMascota int, estado string) ([]ExamenMascota, error) {
	examenesMascota, err := s.repo.GetExamenesMascotaPorMascotayEstado(ctx, idMascota, estado)
	if err != nil {
		return nil, err
	}
	result := []ExamenMascota{}
	for _, item := range examenesMascota {
		result = append(result, ExamenMascota{item})
	}
	return result, nil
}

// CreateExamenMascotaRequest represents an examenesMascota creation request.
type CreateExamenMascotaRequest struct {
	IdUsuario      int        `json:"id_usuario"`
	IdMascota      string     `json:"id_mascota"`
	FechaSolicitud time.Time  `json:"ruta"`
	FechaLlenado   *time.Time `json:"icono"`
	Estado         string     `json:"principal"`
}

type UpdateExamenMascotaRequest struct {
	IdExamenMascota int        `json:"id_examen_mascota"`
	IdUsuario       int        `json:"id_usuario"`
	IdMascota       string     `json:"id_mascota"`
	FechaSolicitud  time.Time  `json:"ruta"`
	FechaLlenado    *time.Time `json:"icono"`
	Estado          string     `json:"principal"`
}

// Validate validates the UpdateExamenMascotaRequest fields.
func (m UpdateExamenMascotaRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdMascota, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.IdUsuario, validation.Required, validation.Length(0, 128)),
	)
}

// Validate validates the CreateExamenMascotaRequest fields.
func (m CreateExamenMascotaRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdMascota, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.IdUsuario, validation.Required, validation.Length(0, 128)),
	)
}

// CrearExamenMascota creates a new examenesMascota.
func (s service) CrearExamenMascota(ctx context.Context, req CreateExamenMascotaRequest) (ExamenMascota, error) {
	if err := req.Validate(); err != nil {
		return ExamenMascota{}, err
	}
	ExamenMascotaG, err := s.repo.CrearExamenMascota(ctx, entity.ExamenMascota{
		IdUsuario:      req.IdUsuario,
		IdMascota:      req.IdMascota,
		FechaSolicitud: req.FechaSolicitud,
		FechaLlenado:   req.FechaLlenado,
		Estado:         req.Estado,
	})
	if err != nil {
		return ExamenMascota{}, err
	}
	return ExamenMascota{ExamenMascotaG}, nil
}

// ActualizarExamenMascota creates a new examenesMascota.
func (s service) ActualizarExamenMascota(ctx context.Context, req UpdateExamenMascotaRequest) (ExamenMascota, error) {
	if err := req.ValidateUpdate(); err != nil {
		return ExamenMascota{}, err
	}
	ExamenMascotaG, err := s.repo.ActualizarExamenMascota(ctx, entity.ExamenMascota{
		IdExamenMascota: req.IdExamenMascota,
		IdUsuario:       req.IdUsuario,
		IdMascota:       req.IdMascota,
		FechaSolicitud:  req.FechaSolicitud,
		FechaLlenado:    req.FechaLlenado,
		Estado:          req.Estado,
	})
	if err != nil {
		return ExamenMascota{}, err
	}
	return ExamenMascota{ExamenMascotaG}, nil
}

// GetExamenMascotaPorId returns the examenesMascota with the specified the examenesMascota ID.
func (s service) GetExamenMascotaPorId(ctx context.Context, idExamenMascota int) (ExamenMascota, error) {
	examenesMascota, err := s.repo.GetExamenMascotaPorId(ctx, idExamenMascota)
	if err != nil {
		return ExamenMascota{}, err
	}
	return ExamenMascota{examenesMascota}, nil
}
