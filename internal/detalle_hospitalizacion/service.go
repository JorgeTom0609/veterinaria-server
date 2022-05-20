package detalle_hospitalizacion

import (
	"context"
	"time"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for detallesHospitalizacion.
type Service interface {
	GetDetallesHospitalizacion(ctx context.Context) ([]DetalleHospitalizacion, error)
	GetDetalleHospitalizacionPorId(ctx context.Context, idDetalleHospitalizacion int) (DetalleHospitalizacion, error)
	CrearDetalleHospitalizacion(ctx context.Context, input CreateDetalleHospitalizacionRequest) (DetalleHospitalizacion, error)
	ActualizarDetalleHospitalizacion(ctx context.Context, input UpdateDetalleHospitalizacionRequest) (DetalleHospitalizacion, error)
}

// DetallesHospitalizacion represents the data about an detallesHospitalizacion.
type DetalleHospitalizacion struct {
	entity.DetalleHospitalizacion
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new detallesHospitalizacion service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list detallesHospitalizacion.
func (s service) GetDetallesHospitalizacion(ctx context.Context) ([]DetalleHospitalizacion, error) {
	detallesHospitalizacion, err := s.repo.GetDetallesHospitalizacion(ctx)
	if err != nil {
		return nil, err
	}
	result := []DetalleHospitalizacion{}
	for _, item := range detallesHospitalizacion {
		result = append(result, DetalleHospitalizacion{item})
	}
	return result, nil
}

// CreateDetalleHospitalizacionRequest represents an detalleHospitalizacion creation request.
type CreateDetalleHospitalizacionRequest struct {
	IdHospitalizacion int       `json:"id_hospitalizacion"`
	IdUsuario         int       `json:"id_usuario"`
	Descripcion       string    `json:"descripcion"`
	Fecha             time.Time `json:"fecha"`
}

type UpdateDetalleHospitalizacionRequest struct {
	IdDetalleHospitalizacion int       `json:"id_detalle_hospitalizacion"`
	IdHospitalizacion        int       `json:"id_hospitalizacion"`
	IdUsuario                int       `json:"id_usuario"`
	Descripcion              string    `json:"descripcion"`
	Fecha                    time.Time `json:"fecha"`
}

// Validate validates the UpdateDetalleHospitalizacionRequest fields.
func (m UpdateDetalleHospitalizacionRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdHospitalizacion, validation.Required),
		validation.Field(&m.IdUsuario, validation.Required),
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 128)),
	)
}

// Validate validates the CreateDetalleHospitalizacionRequest fields.
func (m CreateDetalleHospitalizacionRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdHospitalizacion, validation.Required),
		validation.Field(&m.IdUsuario, validation.Required),
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 128)),
	)
}

// CrearDetalleHospitalizacion creates a new detalleHospitalizacion.
func (s service) CrearDetalleHospitalizacion(ctx context.Context, req CreateDetalleHospitalizacionRequest) (DetalleHospitalizacion, error) {
	if err := req.Validate(); err != nil {
		return DetalleHospitalizacion{}, err
	}
	detalleHospitalizacionG, err := s.repo.CrearDetalleHospitalizacion(ctx, entity.DetalleHospitalizacion{
		IdHospitalizacion: req.IdHospitalizacion,
		IdUsuario:         req.IdUsuario,
		Descripcion:       req.Descripcion,
		Fecha:             req.Fecha,
	})
	if err != nil {
		return DetalleHospitalizacion{}, err
	}
	return DetalleHospitalizacion{detalleHospitalizacionG}, nil
}

// ActualizarDetalleHospitalizacion creates a new detalleHospitalizacion.
func (s service) ActualizarDetalleHospitalizacion(ctx context.Context, req UpdateDetalleHospitalizacionRequest) (DetalleHospitalizacion, error) {
	if err := req.ValidateUpdate(); err != nil {
		return DetalleHospitalizacion{}, err
	}
	detalleHospitalizacionG, err := s.repo.ActualizarDetalleHospitalizacion(ctx, entity.DetalleHospitalizacion{
		IdDetalleHospitalizacion: req.IdDetalleHospitalizacion,
		IdHospitalizacion:        req.IdHospitalizacion,
		IdUsuario:                req.IdUsuario,
		Descripcion:              req.Descripcion,
		Fecha:                    req.Fecha,
	})
	if err != nil {
		return DetalleHospitalizacion{}, err
	}
	return DetalleHospitalizacion{detalleHospitalizacionG}, nil
}

// GetDetalleHospitalizacionPorId returns the detalleHospitalizacion with the specified the detalleHospitalizacion ID.
func (s service) GetDetalleHospitalizacionPorId(ctx context.Context, idDetalleHospitalizacion int) (DetalleHospitalizacion, error) {
	detalleHospitalizacion, err := s.repo.GetDetalleHospitalizacionPorId(ctx, idDetalleHospitalizacion)
	if err != nil {
		return DetalleHospitalizacion{}, err
	}
	return DetalleHospitalizacion{detalleHospitalizacion}, nil
}
