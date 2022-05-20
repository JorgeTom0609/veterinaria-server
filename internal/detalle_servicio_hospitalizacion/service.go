package detalle_servicio_hospitalizacion

import (
	"context"
	"time"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for detallesServicioHospitalizacion.
type Service interface {
	GetDetallesServicioHospitalizacion(ctx context.Context) ([]DetalleServicioHospitalizacion, error)
	GetDetalleServicioHospitalizacionPorId(ctx context.Context, idDetalleServicioHospitalizacion int) (DetalleServicioHospitalizacion, error)
	CrearDetalleServicioHospitalizacion(ctx context.Context, input CreateDetalleServicioHospitalizacionRequest) (DetalleServicioHospitalizacion, error)
	ActualizarDetalleServicioHospitalizacion(ctx context.Context, input UpdateDetalleServicioHospitalizacionRequest) (DetalleServicioHospitalizacion, error)
}

// DetallesServicioHospitalizacion represents the data about an detallesServicioHospitalizacion.
type DetalleServicioHospitalizacion struct {
	entity.DetalleServicioHospitalizacion
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new detallesServicioHospitalizacion service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list detallesServicioHospitalizacion.
func (s service) GetDetallesServicioHospitalizacion(ctx context.Context) ([]DetalleServicioHospitalizacion, error) {
	detallesServicioHospitalizacion, err := s.repo.GetDetallesServicioHospitalizacion(ctx)
	if err != nil {
		return nil, err
	}
	result := []DetalleServicioHospitalizacion{}
	for _, item := range detallesServicioHospitalizacion {
		result = append(result, DetalleServicioHospitalizacion{item})
	}
	return result, nil
}

// CreateDetalleServicioHospitalizacionRequest represents an detalleServicioHospitalizacion creation request.
type CreateDetalleServicioHospitalizacionRequest struct {
	IdHospitalizacion int       `json:"id_hospitalizacion"`
	IdUsuario         int       `json:"id_usuario"`
	IdServicio        int       `json:"id_servicio"`
	Descripcion       string    `json:"descripcion"`
	Fecha             time.Time `json:"fecha"`
}

type UpdateDetalleServicioHospitalizacionRequest struct {
	IdDetalleServicioHospitalizacion int       `json:"id_detalle_servicio_hospitalizacion"`
	IdHospitalizacion                int       `json:"id_hospitalizacion"`
	IdUsuario                        int       `json:"id_usuario"`
	IdServicio                       int       `json:"id_servicio"`
	Descripcion                      string    `json:"descripcion"`
	Fecha                            time.Time `json:"fecha"`
}

// Validate validates the UpdateDetalleServicioHospitalizacionRequest fields.
func (m UpdateDetalleServicioHospitalizacionRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdHospitalizacion, validation.Required),
		validation.Field(&m.IdUsuario, validation.Required),
		validation.Field(&m.IdServicio, validation.Required),
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 128)),
	)
}

// Validate validates the CreateDetalleServicioHospitalizacionRequest fields.
func (m CreateDetalleServicioHospitalizacionRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdHospitalizacion, validation.Required),
		validation.Field(&m.IdUsuario, validation.Required),
		validation.Field(&m.IdServicio, validation.Required),
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 128)),
	)
}

// CrearDetalleServicioHospitalizacion creates a new detalleServicioHospitalizacion.
func (s service) CrearDetalleServicioHospitalizacion(ctx context.Context, req CreateDetalleServicioHospitalizacionRequest) (DetalleServicioHospitalizacion, error) {
	if err := req.Validate(); err != nil {
		return DetalleServicioHospitalizacion{}, err
	}
	detalleServicioHospitalizacionG, err := s.repo.CrearDetalleServicioHospitalizacion(ctx, entity.DetalleServicioHospitalizacion{
		IdHospitalizacion: req.IdHospitalizacion,
		IdUsuario:         req.IdUsuario,
		IdServicio:        req.IdServicio,
		Descripcion:       req.Descripcion,
		Fecha:             req.Fecha,
	})
	if err != nil {
		return DetalleServicioHospitalizacion{}, err
	}
	return DetalleServicioHospitalizacion{detalleServicioHospitalizacionG}, nil
}

// ActualizarDetalleServicioHospitalizacion creates a new detalleServicioHospitalizacion.
func (s service) ActualizarDetalleServicioHospitalizacion(ctx context.Context, req UpdateDetalleServicioHospitalizacionRequest) (DetalleServicioHospitalizacion, error) {
	if err := req.ValidateUpdate(); err != nil {
		return DetalleServicioHospitalizacion{}, err
	}
	detalleServicioHospitalizacionG, err := s.repo.ActualizarDetalleServicioHospitalizacion(ctx, entity.DetalleServicioHospitalizacion{
		IdDetalleServicioHospitalizacion: req.IdDetalleServicioHospitalizacion,
		IdHospitalizacion:                req.IdHospitalizacion,
		IdUsuario:                        req.IdUsuario,
		IdServicio:                       req.IdServicio,
		Descripcion:                      req.Descripcion,
		Fecha:                            req.Fecha,
	})
	if err != nil {
		return DetalleServicioHospitalizacion{}, err
	}
	return DetalleServicioHospitalizacion{detalleServicioHospitalizacionG}, nil
}

// GetDetalleServicioHospitalizacionPorId returns the detalleServicioHospitalizacion with the specified the detalleServicioHospitalizacion ID.
func (s service) GetDetalleServicioHospitalizacionPorId(ctx context.Context, idDetalleServicioHospitalizacion int) (DetalleServicioHospitalizacion, error) {
	detalleServicioHospitalizacion, err := s.repo.GetDetalleServicioHospitalizacionPorId(ctx, idDetalleServicioHospitalizacion)
	if err != nil {
		return DetalleServicioHospitalizacion{}, err
	}
	return DetalleServicioHospitalizacion{detalleServicioHospitalizacion}, nil
}
