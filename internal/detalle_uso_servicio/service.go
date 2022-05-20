package detalle_uso_servicio

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for detallesUsoServicio.
type Service interface {
	GetDetallesUsoServicio(ctx context.Context) ([]DetalleUsoServicio, error)
	GetDetalleUsoServicioPorId(ctx context.Context, idDetalleUsoServicio int) (DetalleUsoServicio, error)
	CrearDetalleUsoServicio(ctx context.Context, input CreateDetalleUsoServicioRequest) (DetalleUsoServicio, error)
	ActualizarDetalleUsoServicio(ctx context.Context, input UpdateDetalleUsoServicioRequest) (DetalleUsoServicio, error)
}

// DetallesUsoServicio represents the data about an detallesUsoServicio.
type DetalleUsoServicio struct {
	entity.DetalleUsoServicio
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new detallesUsoServicio service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list detallesUsoServicio.
func (s service) GetDetallesUsoServicio(ctx context.Context) ([]DetalleUsoServicio, error) {
	detallesUsoServicio, err := s.repo.GetDetallesUsoServicio(ctx)
	if err != nil {
		return nil, err
	}
	result := []DetalleUsoServicio{}
	for _, item := range detallesUsoServicio {
		result = append(result, DetalleUsoServicio{item})
	}
	return result, nil
}

// CreateDetalleUsoServicioRequest represents an detalleUsoServicio creation request.
type CreateDetalleUsoServicioRequest struct {
	IdDetalleServicioHospitalizacion int     `json:"id_detalle_servicio_hospitalizacion"`
	IdReferencia                     int     `json:"id_referencia"`
	Tabla                            string  `json:"tabla"`
	Cantidad                         float32 `json:"cantidad"`
	Valor                            float32 `json:"valor"`
}

type UpdateDetalleUsoServicioRequest struct {
	IdDetalleUsoServicio             int     `json:"id_detalle_uso_servicio"`
	IdDetalleServicioHospitalizacion int     `json:"id_detalle_servicio_hospitalizacion"`
	IdReferencia                     int     `json:"id_referencia"`
	Tabla                            string  `json:"tabla"`
	Cantidad                         float32 `json:"cantidad"`
	Valor                            float32 `json:"valor"`
}

// Validate validates the UpdateDetalleUsoServicioRequest fields.
func (m UpdateDetalleUsoServicioRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdDetalleServicioHospitalizacion, validation.Required),
		validation.Field(&m.IdReferencia, validation.Required),
		validation.Field(&m.Tabla, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Cantidad, validation.Required),
	)
}

// Validate validates the CreateDetalleUsoServicioRequest fields.
func (m CreateDetalleUsoServicioRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdDetalleServicioHospitalizacion, validation.Required),
		validation.Field(&m.IdReferencia, validation.Required),
		validation.Field(&m.Tabla, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Cantidad, validation.Required),
	)
}

// CrearDetalleUsoServicio creates a new detalleUsoServicio.
func (s service) CrearDetalleUsoServicio(ctx context.Context, req CreateDetalleUsoServicioRequest) (DetalleUsoServicio, error) {
	if err := req.Validate(); err != nil {
		return DetalleUsoServicio{}, err
	}
	detalleUsoServicioG, err := s.repo.CrearDetalleUsoServicio(ctx, entity.DetalleUsoServicio{
		IdDetalleServicioHospitalizacion: req.IdDetalleServicioHospitalizacion,
		IdReferencia:                     req.IdReferencia,
		Tabla:                            req.Tabla,
		Cantidad:                         req.Cantidad,
		Valor:                            req.Valor,
	})
	if err != nil {
		return DetalleUsoServicio{}, err
	}
	return DetalleUsoServicio{detalleUsoServicioG}, nil
}

// ActualizarDetalleUsoServicio creates a new detalleUsoServicio.
func (s service) ActualizarDetalleUsoServicio(ctx context.Context, req UpdateDetalleUsoServicioRequest) (DetalleUsoServicio, error) {
	if err := req.ValidateUpdate(); err != nil {
		return DetalleUsoServicio{}, err
	}
	detalleUsoServicioG, err := s.repo.ActualizarDetalleUsoServicio(ctx, entity.DetalleUsoServicio{
		IdDetalleUsoServicio:             req.IdDetalleUsoServicio,
		IdDetalleServicioHospitalizacion: req.IdDetalleServicioHospitalizacion,
		IdReferencia:                     req.IdReferencia,
		Tabla:                            req.Tabla,
		Cantidad:                         req.Cantidad,
		Valor:                            req.Valor,
	})
	if err != nil {
		return DetalleUsoServicio{}, err
	}
	return DetalleUsoServicio{detalleUsoServicioG}, nil
}

// GetDetalleUsoServicioPorId returns the detalleUsoServicio with the specified the detalleUsoServicio ID.
func (s service) GetDetalleUsoServicioPorId(ctx context.Context, idDetalleUsoServicio int) (DetalleUsoServicio, error) {
	detalleUsoServicio, err := s.repo.GetDetalleUsoServicioPorId(ctx, idDetalleUsoServicio)
	if err != nil {
		return DetalleUsoServicio{}, err
	}
	return DetalleUsoServicio{detalleUsoServicio}, nil
}
