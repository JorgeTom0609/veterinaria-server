package detalle_servicio_consulta

import (
	"context"
	"time"
	"veterinaria-server/internal/detalle_uso_servicio_consulta"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for detallesServicioConsulta.
type Service interface {
	GetDetallesServicioConsulta(ctx context.Context) ([]DetalleServicioConsulta, error)
	GetDetalleServicioConsultaPorId(ctx context.Context, idDetalleServicioConsulta int) (DetalleServicioConsulta, error)
	GetDetalleServicioConsultaPorConsulta(ctx context.Context, idConsulta int) ([]DetalleServicioConsultaConDatos, error)
	CrearDetalleServicioConsulta(ctx context.Context, input CreateDetalleServicioConsultaRequest) (DetalleServicioConsulta, error)
	ActualizarDetalleServicioConsulta(ctx context.Context, input UpdateDetalleServicioConsultaRequest) (DetalleServicioConsulta, error)
}

// DetallesServicioConsulta represents the data about an detallesServicioConsulta.
type DetalleServicioConsulta struct {
	entity.DetalleServicioConsulta
}

type DetalleServicioConsultaConDatos struct {
	entity.DetalleServicioConsulta
	Servicio string `json:"servicio"`
}
type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new detallesServicioConsulta service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list detallesServicioConsulta.
func (s service) GetDetallesServicioConsulta(ctx context.Context) ([]DetalleServicioConsulta, error) {
	detallesServicioConsulta, err := s.repo.GetDetallesServicioConsulta(ctx)
	if err != nil {
		return nil, err
	}
	result := []DetalleServicioConsulta{}
	for _, item := range detallesServicioConsulta {
		result = append(result, DetalleServicioConsulta{item})
	}
	return result, nil
}

type CreateDetalleServicioConsultaConDetallesRequest struct {
	DetalleServicioConsulta CreateDetalleServicioConsultaRequest                                    `json:"detalle_servicio_consulta"`
	Productos               []detalle_uso_servicio_consulta.CreateDetalleUsoServicioConsultaRequest `json:"productos"`
}

// CreateDetalleServicioConsultaRequest represents an detalleServicioConsulta creation request.
type CreateDetalleServicioConsultaRequest struct {
	IdConsulta int       `json:"id_consulta"`
	IdServicio int       `json:"id_servicio"`
	Valor      float32   `json:"valor"`
	Fecha      time.Time `json:"fecha"`
	Servicio   string    `json:"servicio"`
}

type UpdateDetalleServicioConsultaRequest struct {
	IdDetalleServicioConsulta int       `json:"id_detalle_servicio_consulta"`
	IdConsulta                int       `json:"id_consulta"`
	IdServicio                int       `json:"id_servicio"`
	Valor                     float32   `json:"valor"`
	Fecha                     time.Time `json:"fecha"`
}

// Validate validates the UpdateDetalleServicioConsultaRequest fields.
func (m UpdateDetalleServicioConsultaRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdConsulta, validation.Required),
		validation.Field(&m.IdServicio, validation.Required),
	)
}

// Validate validates the CreateDetalleServicioConsultaRequest fields.
func (m CreateDetalleServicioConsultaRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdConsulta, validation.Required),
		validation.Field(&m.IdServicio, validation.Required),
	)
}

// CrearDetalleServicioConsulta creates a new detalleServicioConsulta.
func (s service) CrearDetalleServicioConsulta(ctx context.Context, req CreateDetalleServicioConsultaRequest) (DetalleServicioConsulta, error) {
	if err := req.Validate(); err != nil {
		return DetalleServicioConsulta{}, err
	}
	detalleServicioConsultaG, err := s.repo.CrearDetalleServicioConsulta(ctx, entity.DetalleServicioConsulta{
		IdConsulta: req.IdConsulta,
		IdServicio: req.IdServicio,
		Valor:      req.Valor,
		Fecha:      req.Fecha,
	})
	if err != nil {
		return DetalleServicioConsulta{}, err
	}
	return DetalleServicioConsulta{detalleServicioConsultaG}, nil
}

// ActualizarDetalleServicioConsulta creates a new detalleServicioConsulta.
func (s service) ActualizarDetalleServicioConsulta(ctx context.Context, req UpdateDetalleServicioConsultaRequest) (DetalleServicioConsulta, error) {
	if err := req.ValidateUpdate(); err != nil {
		return DetalleServicioConsulta{}, err
	}
	detalleServicioConsultaG, err := s.repo.ActualizarDetalleServicioConsulta(ctx, entity.DetalleServicioConsulta{
		IdDetalleServicioConsulta: req.IdDetalleServicioConsulta,
		IdConsulta:                req.IdConsulta,
		IdServicio:                req.IdServicio,
		Valor:                     req.Valor,
		Fecha:                     req.Fecha,
	})
	if err != nil {
		return DetalleServicioConsulta{}, err
	}
	return DetalleServicioConsulta{detalleServicioConsultaG}, nil
}

// GetDetalleServicioConsultaPorId returns the detalleServicioConsulta with the specified the detalleServicioConsulta ID.
func (s service) GetDetalleServicioConsultaPorId(ctx context.Context, idDetalleServicioConsulta int) (DetalleServicioConsulta, error) {
	detalleServicioConsulta, err := s.repo.GetDetalleServicioConsultaPorId(ctx, idDetalleServicioConsulta)
	if err != nil {
		return DetalleServicioConsulta{}, err
	}
	return DetalleServicioConsulta{detalleServicioConsulta}, nil

}
func (s service) GetDetalleServicioConsultaPorConsulta(ctx context.Context, idConsulta int) ([]DetalleServicioConsultaConDatos, error) {
	detallesServicioConsulta, err := s.repo.GetDetalleServicioConsultaPorConsulta(ctx, idConsulta)
	if err != nil {
		return []DetalleServicioConsultaConDatos{}, err
	}
	return detallesServicioConsulta, nil
}
