package hospitalizacion

import (
	"context"
	"time"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for hospitalizaciones.
type Service interface {
	GetHospitalizaciones(ctx context.Context) ([]Hospitalizacion, error)
	GetHospitalizacionPorId(ctx context.Context, idHospitalizacion int) (Hospitalizacion, error)
	CrearHospitalizacion(ctx context.Context, input CreateHospitalizacionRequest) (Hospitalizacion, error)
	ActualizarHospitalizacion(ctx context.Context, input UpdateHospitalizacionRequest) (Hospitalizacion, error)
}

// Hospitalizaciones represents the data about an hospitalizaciones.
type Hospitalizacion struct {
	entity.Hospitalizacion
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new hospitalizaciones service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list hospitalizaciones.
func (s service) GetHospitalizaciones(ctx context.Context) ([]Hospitalizacion, error) {
	hospitalizaciones, err := s.repo.GetHospitalizaciones(ctx)
	if err != nil {
		return nil, err
	}
	result := []Hospitalizacion{}
	for _, item := range hospitalizaciones {
		result = append(result, Hospitalizacion{item})
	}
	return result, nil
}

// CreateHospitalizacionRequest represents an hospitalizacion creation request.
type CreateHospitalizacionRequest struct {
	IdConsulta            int        `json:"id_consulta"`
	Motivo                string     `json:"motivo"`
	FechaIngreso          time.Time  `json:"fecha_ingreso"`
	FechaSalida           *time.Time `json:"fecha_salida"`
	Valor                 *float32   `json:"valor"`
	EstadoHospitalizacion string     `json:"estado_hospitalizacion"`
}

type UpdateHospitalizacionRequest struct {
	IdHospitalizacion     int        `json:"id_hospitalizacion"`
	IdConsulta            int        `json:"id_consulta"`
	Motivo                string     `json:"motivo"`
	FechaIngreso          time.Time  `json:"fecha_ingreso"`
	FechaSalida           *time.Time `json:"fecha_salida"`
	Valor                 *float32   `json:"valor"`
	EstadoHospitalizacion string     `json:"estado_hospitalizacion"`
}

// Validate validates the UpdateHospitalizacionRequest fields.
func (m UpdateHospitalizacionRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdConsulta, validation.Required),
		validation.Field(&m.Motivo, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.EstadoHospitalizacion, validation.Required, validation.Length(0, 128)),
	)
}

// Validate validates the CreateHospitalizacionRequest fields.
func (m CreateHospitalizacionRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdConsulta, validation.Required),
		validation.Field(&m.Motivo, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.EstadoHospitalizacion, validation.Required, validation.Length(0, 128)),
	)
}

// CrearHospitalizacion creates a new hospitalizacion.
func (s service) CrearHospitalizacion(ctx context.Context, req CreateHospitalizacionRequest) (Hospitalizacion, error) {
	if err := req.Validate(); err != nil {
		return Hospitalizacion{}, err
	}
	hospitalizacionG, err := s.repo.CrearHospitalizacion(ctx, entity.Hospitalizacion{
		IdConsulta:            req.IdConsulta,
		Motivo:                req.Motivo,
		FechaIngreso:          req.FechaIngreso,
		FechaSalida:           &req.FechaIngreso,
		Valor:                 req.Valor,
		EstadoHospitalizacion: req.EstadoHospitalizacion,
	})
	if err != nil {
		return Hospitalizacion{}, err
	}
	return Hospitalizacion{hospitalizacionG}, nil
}

// ActualizarHospitalizacion creates a new hospitalizacion.
func (s service) ActualizarHospitalizacion(ctx context.Context, req UpdateHospitalizacionRequest) (Hospitalizacion, error) {
	if err := req.ValidateUpdate(); err != nil {
		return Hospitalizacion{}, err
	}
	hospitalizacionG, err := s.repo.ActualizarHospitalizacion(ctx, entity.Hospitalizacion{
		IdHospitalizacion:     req.IdHospitalizacion,
		IdConsulta:            req.IdConsulta,
		Motivo:                req.Motivo,
		FechaIngreso:          req.FechaIngreso,
		FechaSalida:           &req.FechaIngreso,
		Valor:                 req.Valor,
		EstadoHospitalizacion: req.EstadoHospitalizacion,
	})
	if err != nil {
		return Hospitalizacion{}, err
	}
	return Hospitalizacion{hospitalizacionG}, nil
}

// GetHospitalizacionPorId returns the hospitalizacion with the specified the hospitalizacion ID.
func (s service) GetHospitalizacionPorId(ctx context.Context, idHospitalizacion int) (Hospitalizacion, error) {
	hospitalizacion, err := s.repo.GetHospitalizacionPorId(ctx, idHospitalizacion)
	if err != nil {
		return Hospitalizacion{}, err
	}
	return Hospitalizacion{hospitalizacion}, nil
}
