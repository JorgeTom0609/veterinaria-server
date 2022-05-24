package detalle_uso_servicio_consulta

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for detallesUsoServicioConsulta.
type Service interface {
	GetDetallesUsoServicioConsulta(ctx context.Context) ([]DetalleUsoServicioConsulta, error)
	GetDetalleUsoServicioConsultaPorId(ctx context.Context, idDetalleUsoServicioConsulta int) (DetalleUsoServicioConsulta, error)
	CrearDetalleUsoServicioConsulta(ctx context.Context, input CreateDetalleUsoServicioConsultaRequest) (DetalleUsoServicioConsulta, error)
	ActualizarDetalleUsoServicioConsulta(ctx context.Context, input UpdateDetalleUsoServicioConsultaRequest) (DetalleUsoServicioConsulta, error)
}

// DetallesUsoServicioConsulta represents the data about an detallesUsoServicioConsulta.
type DetalleUsoServicioConsulta struct {
	entity.DetalleUsoServicioConsulta
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new detallesUsoServicioConsulta service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list detallesUsoServicioConsulta.
func (s service) GetDetallesUsoServicioConsulta(ctx context.Context) ([]DetalleUsoServicioConsulta, error) {
	detallesUsoServicioConsulta, err := s.repo.GetDetallesUsoServicioConsulta(ctx)
	if err != nil {
		return nil, err
	}
	result := []DetalleUsoServicioConsulta{}
	for _, item := range detallesUsoServicioConsulta {
		result = append(result, DetalleUsoServicioConsulta{item})
	}
	return result, nil
}

// CreateDetalleUsoServicioConsultaRequest represents an detalleUsoServicioConsulta creation request.
type CreateDetalleUsoServicioConsultaRequest struct {
	IdDetalleServicioConsulta int     `json:"id_detalle_servicio_consulta"`
	IdReferencia              int     `json:"id_referencia"`
	Tabla                     string  `json:"tabla"`
	Cantidad                  float32 `json:"cantidad"`
}

type UpdateDetalleUsoServicioConsultaRequest struct {
	IdDetalleUsoServicioConsulta int     `json:"id_detalle_uso_servicio_consulta"`
	IdDetalleServicioConsulta    int     `json:"id_detalle_servicio_consulta"`
	IdReferencia                 int     `json:"id_referencia"`
	Tabla                        string  `json:"tabla"`
	Cantidad                     float32 `json:"cantidad"`
}

// Validate validates the UpdateDetalleUsoServicioConsultaRequest fields.
func (m UpdateDetalleUsoServicioConsultaRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdDetalleServicioConsulta, validation.Required),
		validation.Field(&m.IdReferencia, validation.Required),
		validation.Field(&m.Tabla, validation.Required, validation.Length(0, 1000)),
	)
}

// Validate validates the CreateDetalleUsoServicioConsultaRequest fields.
func (m CreateDetalleUsoServicioConsultaRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdDetalleServicioConsulta, validation.Required),
		validation.Field(&m.IdReferencia, validation.Required),
		validation.Field(&m.Tabla, validation.Required, validation.Length(0, 1000)),
	)
}

// CrearDetalleUsoServicioConsulta creates a new detalleUsoServicioConsulta.
func (s service) CrearDetalleUsoServicioConsulta(ctx context.Context, req CreateDetalleUsoServicioConsultaRequest) (DetalleUsoServicioConsulta, error) {
	if err := req.Validate(); err != nil {
		return DetalleUsoServicioConsulta{}, err
	}
	detalleUsoServicioConsultaG, err := s.repo.CrearDetalleUsoServicioConsulta(ctx, entity.DetalleUsoServicioConsulta{
		IdDetalleServicioConsulta: req.IdDetalleServicioConsulta,
		IdReferencia:              req.IdReferencia,
		Tabla:                     req.Tabla,
		Cantidad:                  req.Cantidad,
	})
	if err != nil {
		return DetalleUsoServicioConsulta{}, err
	}
	return DetalleUsoServicioConsulta{detalleUsoServicioConsultaG}, nil
}

// ActualizarDetalleUsoServicioConsulta creates a new detalleUsoServicioConsulta.
func (s service) ActualizarDetalleUsoServicioConsulta(ctx context.Context, req UpdateDetalleUsoServicioConsultaRequest) (DetalleUsoServicioConsulta, error) {
	if err := req.ValidateUpdate(); err != nil {
		return DetalleUsoServicioConsulta{}, err
	}
	detalleUsoServicioConsultaG, err := s.repo.ActualizarDetalleUsoServicioConsulta(ctx, entity.DetalleUsoServicioConsulta{
		IdDetalleUsoServicioConsulta: req.IdDetalleUsoServicioConsulta,
		IdDetalleServicioConsulta:    req.IdDetalleServicioConsulta,
		IdReferencia:                 req.IdReferencia,
		Tabla:                        req.Tabla,
		Cantidad:                     req.Cantidad,
	})
	if err != nil {
		return DetalleUsoServicioConsulta{}, err
	}
	return DetalleUsoServicioConsulta{detalleUsoServicioConsultaG}, nil
}

// GetDetalleUsoServicioConsultaPorId returns the detalleUsoServicioConsulta with the specified the detalleUsoServicioConsulta ID.
func (s service) GetDetalleUsoServicioConsultaPorId(ctx context.Context, idDetalleUsoServicioConsulta int) (DetalleUsoServicioConsulta, error) {
	detalleUsoServicioConsulta, err := s.repo.GetDetalleUsoServicioConsultaPorId(ctx, idDetalleUsoServicioConsulta)
	if err != nil {
		return DetalleUsoServicioConsulta{}, err
	}
	return DetalleUsoServicioConsulta{detalleUsoServicioConsulta}, nil
}
