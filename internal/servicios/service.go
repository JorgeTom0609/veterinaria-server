package servicios

import (
	"context"
	"database/sql"
	"veterinaria-server/internal/entity"
	"veterinaria-server/internal/servicio_producto"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for servicios.
type Service interface {
	GetServicios(ctx context.Context) ([]Servicio, error)
	GetServicioPorEspecie(ctx context.Context, idEspecie int, modo int) ([]ServicioTieneProductos, error)
	GetServiciosConProductos(ctx context.Context) ([]ServicioTieneProductos, error)
	GetServicioPorId(ctx context.Context, idServicio int) (Servicio, error)
	CrearServicio(ctx context.Context, input CreateServicioRequest) (Servicio, error)
	ActualizarServicio(ctx context.Context, input UpdateServicioRequest) (Servicio, error)
}

// Servicios represents the data about an servicios.
type Servicio struct {
	entity.Servicio
}

type ServicioTieneProductos struct {
	entity.Servicio
	CantidadProducto int `json:"cantidad_producto"`
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new servicios service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list servicios.
func (s service) GetServicios(ctx context.Context) ([]Servicio, error) {
	servicios, err := s.repo.GetServicios(ctx)
	if err != nil {
		return nil, err
	}
	result := []Servicio{}
	for _, item := range servicios {
		result = append(result, Servicio{item})
	}
	return result, nil
}

func (s service) GetServicioPorEspecie(ctx context.Context, idEspecie int, modo int) ([]ServicioTieneProductos, error) {
	servicios, err := s.repo.GetServicioPorEspecie(ctx, idEspecie, modo)
	if err != nil {
		return nil, err
	}
	return servicios, nil
}

func (s service) GetServiciosConProductos(ctx context.Context) ([]ServicioTieneProductos, error) {
	servicios, err := s.repo.GetServiciosConProductos(ctx)
	if err != nil {
		return nil, err
	}
	return servicios, nil
}

// CreateServicioRequest represents an servicio creation request.
type CreateServicioRequest struct {
	IdEspecie             int          `json:"id_especie"`
	IdUsuario             int          `json:"id_usuario"`
	Descripcion           string       `json:"descripcion"`
	Valor                 float32      `json:"valor"`
	AplicaConsulta        sql.NullBool `json:"aplica_consulta"`
	AplicaHospitalizacion sql.NullBool `json:"aplica_hospitalizacion"`
}

type UpdateServicioRequest struct {
	IdServicio            int          `json:"id_servicio"`
	IdEspecie             int          `json:"id_especie"`
	IdUsuario             int          `json:"id_usuario"`
	Descripcion           string       `json:"descripcion"`
	Valor                 float32      `json:"valor"`
	AplicaConsulta        sql.NullBool `json:"aplica_consulta"`
	AplicaHospitalizacion sql.NullBool `json:"aplica_hospitalizacion"`
}

type UpdateServicioConDetallesRequest struct {
	Servicio          UpdateServicioRequest                             `json:"servicio"`
	ServicioProductos []servicio_producto.UpdateServicioProductoRequest `json:"productos"`
}

// Validate validates the UpdateServicioRequest fields.
func (m UpdateServicioRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 1000)),
	)
}

// Validate validates the CreateServicioRequest fields.
func (m CreateServicioRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 1000)),
	)
}

// CrearServicio creates a new servicio.
func (s service) CrearServicio(ctx context.Context, req CreateServicioRequest) (Servicio, error) {
	if err := req.Validate(); err != nil {
		return Servicio{}, err
	}
	servicioG, err := s.repo.CrearServicio(ctx, entity.Servicio{
		IdUsuario:             req.IdUsuario,
		IdEspecie:             req.IdEspecie,
		Descripcion:           req.Descripcion,
		Valor:                 req.Valor,
		AplicaConsulta:        req.AplicaConsulta,
		AplicaHospitalizacion: req.AplicaHospitalizacion,
	})
	if err != nil {
		return Servicio{}, err
	}
	return Servicio{servicioG}, nil
}

// ActualizarServicio creates a new servicio.
func (s service) ActualizarServicio(ctx context.Context, req UpdateServicioRequest) (Servicio, error) {
	if err := req.ValidateUpdate(); err != nil {
		return Servicio{}, err
	}
	servicioG, err := s.repo.ActualizarServicio(ctx, entity.Servicio{
		IdServicio:            req.IdServicio,
		IdUsuario:             req.IdUsuario,
		IdEspecie:             req.IdEspecie,
		Descripcion:           req.Descripcion,
		Valor:                 req.Valor,
		AplicaConsulta:        req.AplicaConsulta,
		AplicaHospitalizacion: req.AplicaHospitalizacion,
	})
	if err != nil {
		return Servicio{}, err
	}
	return Servicio{servicioG}, nil
}

// GetServicioPorId returns the servicio with the specified the servicio ID.
func (s service) GetServicioPorId(ctx context.Context, idServicio int) (Servicio, error) {
	servicio, err := s.repo.GetServicioPorId(ctx, idServicio)
	if err != nil {
		return Servicio{}, err
	}
	return Servicio{servicio}, nil
}
