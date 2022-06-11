package cita_medica

import (
	"context"
	"time"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for citasMedica.
type Service interface {
	GetCitasMedica(ctx context.Context) ([]CitaMedica, error)
	GetCitasMedicaPendientes(ctx context.Context) ([]CitaMedica, error)
	GetCitasMedicaSinNotificar(ctx context.Context) ([]CitaMedicaDatos, error)
	GetCitaMedicaPorId(ctx context.Context, idCitaMedica int) (CitaMedica, error)
	CrearCitaMedica(ctx context.Context, input CreateCitaMedicaRequest) (CitaMedica, error)
	ActualizarCitaMedica(ctx context.Context, input UpdateCitaMedicaRequest) (CitaMedica, error)
}

// CitasMedica represents the data about an citasMedica.
type CitaMedica struct {
	entity.CitaMedica
}

type CitaMedicaDatos struct {
	entity.CitaMedica
	Duenio   string `json:"duenio"`
	Telefono string `json:"telefono"`
	Mascota  string `json:"mascota"`
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new citasMedica service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list citasMedica.
func (s service) GetCitasMedica(ctx context.Context) ([]CitaMedica, error) {
	citasMedica, err := s.repo.GetCitasMedica(ctx)
	if err != nil {
		return nil, err
	}
	result := []CitaMedica{}
	for _, item := range citasMedica {
		result = append(result, CitaMedica{item})
	}
	return result, nil
}

func (s service) GetCitasMedicaPendientes(ctx context.Context) ([]CitaMedica, error) {
	citasMedica, err := s.repo.GetCitasMedicaPendientes(ctx)
	if err != nil {
		return nil, err
	}
	result := []CitaMedica{}
	for _, item := range citasMedica {
		result = append(result, CitaMedica{item})
	}
	return result, nil
}

func (s service) GetCitasMedicaSinNotificar(ctx context.Context) ([]CitaMedicaDatos, error) {
	citasMedica, err := s.repo.GetCitasMedicaSinNotificar(ctx)
	if err != nil {
		return nil, err
	}
	return citasMedica, nil
}

// CreateCitaMedicaRequest represents an citaMedica creation request.
type CreateCitaMedicaRequest struct {
	IdMascota          int       `json:"id_mascota"`
	Motivo             string    `json:"motivo"`
	Fecha              time.Time `json:"fecha"`
	EstadoNotificacion string    `json:"estado_notificacion"`
}

type UpdateCitaMedicaRequest struct {
	IdCitaMedica       int       `json:"id_cita_medica"`
	IdMascota          int       `json:"id_mascota"`
	Motivo             string    `json:"motivo"`
	Fecha              time.Time `json:"fecha"`
	EstadoNotificacion string    `json:"estado_notificacion"`
}

// Validate validates the UpdateCitaMedicaRequest fields.
func (m UpdateCitaMedicaRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdMascota, validation.Required),
		validation.Field(&m.Motivo, validation.Required, validation.Length(0, 1000)),
	)
}

// Validate validates the CreateCitaMedicaRequest fields.
func (m CreateCitaMedicaRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdMascota, validation.Required),
		validation.Field(&m.Motivo, validation.Required, validation.Length(0, 1000)),
	)
}

// CrearCitaMedica creates a new citaMedica.
func (s service) CrearCitaMedica(ctx context.Context, req CreateCitaMedicaRequest) (CitaMedica, error) {
	if err := req.Validate(); err != nil {
		return CitaMedica{}, err
	}
	citaMedicaG, err := s.repo.CrearCitaMedica(ctx, entity.CitaMedica{
		IdMascota:          req.IdMascota,
		Motivo:             req.Motivo,
		Fecha:              req.Fecha,
		EstadoNotificacion: req.EstadoNotificacion,
	})
	if err != nil {
		return CitaMedica{}, err
	}
	return CitaMedica{citaMedicaG}, nil
}

// ActualizarCitaMedica creates a new citaMedica.
func (s service) ActualizarCitaMedica(ctx context.Context, req UpdateCitaMedicaRequest) (CitaMedica, error) {
	if err := req.ValidateUpdate(); err != nil {
		return CitaMedica{}, err
	}
	citaMedicaG, err := s.repo.ActualizarCitaMedica(ctx, entity.CitaMedica{
		IdCitaMedica:       req.IdCitaMedica,
		IdMascota:          req.IdMascota,
		Motivo:             req.Motivo,
		Fecha:              req.Fecha,
		EstadoNotificacion: req.EstadoNotificacion,
	})
	if err != nil {
		return CitaMedica{}, err
	}
	return CitaMedica{citaMedicaG}, nil
}

// GetCitaMedicaPorId returns the citaMedica with the specified the citaMedica ID.
func (s service) GetCitaMedicaPorId(ctx context.Context, idCitaMedica int) (CitaMedica, error) {
	citaMedica, err := s.repo.GetCitaMedicaPorId(ctx, idCitaMedica)
	if err != nil {
		return CitaMedica{}, err
	}
	return CitaMedica{citaMedica}, nil
}
