package proveedor_producto

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for proveedoresProducto.
type Service interface {
	GetProveedoresProducto(ctx context.Context) ([]ProveedorProducto, error)
	GetProveedorProductoPorId(ctx context.Context, idProveedorProducto int) (ProveedorProducto, error)
	GetProveedorProductoPorIdProveedor(ctx context.Context, idProveedor int) ([]ProveedorProductoConDatos, error)
	CrearProveedorProducto(ctx context.Context, input CreateProveedorProductoRequest) (ProveedorProducto, error)
	ActualizarProveedorProducto(ctx context.Context, input UpdateProveedorProductoRequest) (ProveedorProducto, error)
}

// ProveedoresProducto represents the data about an proveedoresProducto.
type ProveedorProducto struct {
	entity.ProveedorProducto
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new proveedoresProducto service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list proveedoresProducto.
func (s service) GetProveedoresProducto(ctx context.Context) ([]ProveedorProducto, error) {
	proveedoresProducto, err := s.repo.GetProveedoresProducto(ctx)
	if err != nil {
		return nil, err
	}
	result := []ProveedorProducto{}
	for _, item := range proveedoresProducto {
		result = append(result, ProveedorProducto{item})
	}
	return result, nil
}

// CreateProveedorProductoRequest represents an proveedorProducto creation request.
type CreateProveedorProductoRequest struct {
	IdProveedor  int     `json:"id_proveedor"`
	IdProducto   int     `json:"id_producto"`
	PrecioCompra float32 `json:"precio_compra"`
}

type UpdateProveedorProductoRequest struct {
	IdProveedorProducto int     `json:"id_proveedor_producto"`
	IdProveedor         int     `json:"id_proveedor"`
	IdProducto          int     `json:"id_producto"`
	PrecioCompra        float32 `json:"precio_compra"`
}

type UpdateProveedorProductosRequest struct {
	ProveedorProductos []UpdateProveedorProductoRequest `json:"proveedorProductos"`
}

type ProveedorProductoConDatos struct {
	entity.ProveedorProducto
	entity.Producto `json:"producto"`
}

// Validate validates the UpdateProveedorProductoRequest fields.
func (m UpdateProveedorProductoRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdProveedor, validation.Required),
		validation.Field(&m.IdProducto, validation.Required),
	)
}

// Validate validates the CreateProveedorProductoRequest fields.
func (m CreateProveedorProductoRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdProveedor, validation.Required),
		validation.Field(&m.IdProducto, validation.Required),
	)
}

// CrearProveedorProducto creates a new proveedorProducto.
func (s service) CrearProveedorProducto(ctx context.Context, req CreateProveedorProductoRequest) (ProveedorProducto, error) {
	if err := req.Validate(); err != nil {
		return ProveedorProducto{}, err
	}
	proveedorProductoG, err := s.repo.CrearProveedorProducto(ctx, entity.ProveedorProducto{
		IdProveedor:  req.IdProveedor,
		IdProducto:   req.IdProducto,
		PrecioCompra: req.PrecioCompra,
	})
	if err != nil {
		return ProveedorProducto{}, err
	}
	return ProveedorProducto{proveedorProductoG}, nil
}

// ActualizarProveedorProducto creates a new proveedorProducto.
func (s service) ActualizarProveedorProducto(ctx context.Context, req UpdateProveedorProductoRequest) (ProveedorProducto, error) {
	if err := req.ValidateUpdate(); err != nil {
		return ProveedorProducto{}, err
	}
	proveedorProductoG, err := s.repo.ActualizarProveedorProducto(ctx, entity.ProveedorProducto{
		IdProveedorProducto: req.IdProveedorProducto,
		IdProveedor:         req.IdProveedor,
		IdProducto:          req.IdProducto,
		PrecioCompra:        req.PrecioCompra,
	})
	if err != nil {
		return ProveedorProducto{}, err
	}
	return ProveedorProducto{proveedorProductoG}, nil
}

// GetProveedorProductoPorId returns the proveedorProducto with the specified the proveedorProducto ID.
func (s service) GetProveedorProductoPorId(ctx context.Context, idProveedorProducto int) (ProveedorProducto, error) {
	proveedorProducto, err := s.repo.GetProveedorProductoPorId(ctx, idProveedorProducto)
	if err != nil {
		return ProveedorProducto{}, err
	}
	return ProveedorProducto{proveedorProducto}, nil
}

func (s service) GetProveedorProductoPorIdProveedor(ctx context.Context, idProveedor int) ([]ProveedorProductoConDatos, error) {
	proveedorProductoConDatos, err := s.repo.GetProveedorProductoPorIdProveedor(ctx, idProveedor)
	if err != nil {
		return []ProveedorProductoConDatos{}, err
	}
	return proveedorProductoConDatos, nil
}
