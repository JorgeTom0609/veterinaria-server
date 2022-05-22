package servicio_producto

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for servicioProductos.
type Service interface {
	GetServicioProductos(ctx context.Context) ([]ServicioProducto, error)
	GetServicioProductosConDatos(ctx context.Context) ([]ServicioProductoConDatos, error)
	GetServicioProductoPorId(ctx context.Context, idServicioProducto int) (ServicioProducto, error)
	GetServicioProductoPorServicio(ctx context.Context, idServicio int) ([]ServicioProductoConCantidad, error)
	CrearServicioProducto(ctx context.Context, input CreateServicioProductoRequest) (ServicioProducto, error)
	ActualizarServicioProducto(ctx context.Context, input UpdateServicioProductoRequest) (ServicioProducto, error)
}

// ServicioProductos represents the data about an servicioProductos.
type ServicioProducto struct {
	entity.ServicioProducto
}

type ServicioProductoConDatos struct {
	entity.ServicioProducto
	entity.Producto `json:"producto"`
	Unidad          string `json:"unidad"`
}

type ServicioProductoConCantidad struct {
	entity.ServicioProducto
	CantidadUsar *float32 `json:"cantidad_usar"`
	UnidadUsar   *float32 `json:"unidad_usar"`
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new servicioProductos service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list servicioProductos.
func (s service) GetServicioProductos(ctx context.Context) ([]ServicioProducto, error) {
	servicioProductos, err := s.repo.GetServicioProductos(ctx)
	if err != nil {
		return nil, err
	}
	result := []ServicioProducto{}
	for _, item := range servicioProductos {
		result = append(result, ServicioProducto{item})
	}
	return result, nil
}

// Get returns the list servicioProductos.
func (s service) GetServicioProductosConDatos(ctx context.Context) ([]ServicioProductoConDatos, error) {
	servicioProductos, err := s.repo.GetServicioProductosConDatos(ctx)
	if err != nil {
		return nil, err
	}
	return servicioProductos, nil
}

// CreateServicioProductoRequest represents an servicioProducto creation request.
type CreateServicioProductoRequest struct {
	IdServicio int      `json:"id_servicio"`
	IdProducto int      `json:"id_producto"`
	Cantidad   float32  `json:"cantidad"`
	Razon      *float32 `json:"razon"`
}

type UpdateServicioProductoRequest struct {
	IdServicioProducto int      `json:"id_servicio_producto"`
	IdServicio         int      `json:"id_servicio"`
	IdProducto         int      `json:"id_producto"`
	Cantidad           float32  `json:"cantidad"`
	Razon              *float32 `json:"razon"`
}

// Validate validates the UpdateServicioProductoRequest fields.
func (m UpdateServicioProductoRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdServicio, validation.Required),
		validation.Field(&m.IdProducto, validation.Required),
		validation.Field(&m.Cantidad, validation.Required),
	)
}

// Validate validates the CreateServicioProductoRequest fields.
func (m CreateServicioProductoRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdServicio, validation.Required),
		validation.Field(&m.IdProducto, validation.Required),
		validation.Field(&m.Cantidad, validation.Required),
	)
}

// CrearServicioProducto creates a new servicioProducto.
func (s service) CrearServicioProducto(ctx context.Context, req CreateServicioProductoRequest) (ServicioProducto, error) {
	if err := req.Validate(); err != nil {
		return ServicioProducto{}, err
	}
	servicioProductoG, err := s.repo.CrearServicioProducto(ctx, entity.ServicioProducto{
		IdServicio: req.IdServicio,
		IdProducto: req.IdProducto,
		Cantidad:   req.Cantidad,
		Razon:      req.Razon,
	})
	if err != nil {
		return ServicioProducto{}, err
	}
	return ServicioProducto{servicioProductoG}, nil
}

// ActualizarServicioProducto creates a new servicioProducto.
func (s service) ActualizarServicioProducto(ctx context.Context, req UpdateServicioProductoRequest) (ServicioProducto, error) {
	if err := req.ValidateUpdate(); err != nil {
		return ServicioProducto{}, err
	}
	servicioProductoG, err := s.repo.ActualizarServicioProducto(ctx, entity.ServicioProducto{
		IdServicioProducto: req.IdServicioProducto,
		IdServicio:         req.IdServicio,
		IdProducto:         req.IdProducto,
		Cantidad:           req.Cantidad,
		Razon:              req.Razon,
	})
	if err != nil {
		return ServicioProducto{}, err
	}
	return ServicioProducto{servicioProductoG}, nil
}

// GetServicioProductoPorId returns the servicioProducto with the specified the servicioProducto ID.
func (s service) GetServicioProductoPorId(ctx context.Context, idServicioProducto int) (ServicioProducto, error) {
	servicioProducto, err := s.repo.GetServicioProductoPorId(ctx, idServicioProducto)
	if err != nil {
		return ServicioProducto{}, err
	}
	return ServicioProducto{servicioProducto}, nil
}

func (s service) GetServicioProductoPorServicio(ctx context.Context, idServicio int) ([]ServicioProductoConCantidad, error) {
	servicioProductos, err := s.repo.GetServicioProductoPorServicio(ctx, idServicio)
	if err != nil {
		return []ServicioProductoConCantidad{}, err
	}
	return servicioProductos, nil
}
