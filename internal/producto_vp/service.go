package producto_vp

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for productosVP.
type Service interface {
	GetProductosVP(ctx context.Context) ([]ProductoVP, error)
	GetProductosVPConStock(ctx context.Context) ([]ProductoVP, error)
	GetProductosVPPocoStock(ctx context.Context) ([]ProductoVP, error)
	GetProductoVPPorId(ctx context.Context, idProductoVP int) (ProductoVP, error)
	CrearProductoVP(ctx context.Context, input CreateProductoVPRequest) (ProductoVP, error)
	ActualizarProductoVP(ctx context.Context, input UpdateProductoVPRequest) (ProductoVP, error)
}

// ProductosVP represents the data about an productosVP.
type ProductoVP struct {
	entity.ProductoVP
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new productosVP service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list productosVP.
func (s service) GetProductosVP(ctx context.Context) ([]ProductoVP, error) {
	productosVP, err := s.repo.GetProductosVP(ctx)
	if err != nil {
		return nil, err
	}
	result := []ProductoVP{}
	for _, item := range productosVP {
		result = append(result, ProductoVP{item})
	}
	return result, nil
}

func (s service) GetProductosVPConStock(ctx context.Context) ([]ProductoVP, error) {
	productosVP, err := s.repo.GetProductosVPConStock(ctx)
	if err != nil {
		return nil, err
	}
	result := []ProductoVP{}
	for _, item := range productosVP {
		result = append(result, ProductoVP{item})
	}
	return result, nil
}

func (s service) GetProductosVPPocoStock(ctx context.Context) ([]ProductoVP, error) {
	productosVP, err := s.repo.GetProductosVPPocoStock(ctx)
	if err != nil {
		return nil, err
	}
	result := []ProductoVP{}
	for _, item := range productosVP {
		result = append(result, ProductoVP{item})
	}
	return result, nil
}

// CreateProductoVPRequest represents an productoVP creation request.
type CreateProductoVPRequest struct {
	Descripcion  string  `json:"descripcion"`
	PrecioCompra float32 `json:"precio_compra"`
	PrecioVenta  float32 `json:"precio_venta"`
	Stock        int     `json:"stock"`
	StockMinimo  *int    `json:"stock_minimo"`
}

type UpdateProductoVPRequest struct {
	IdProductoVP int     `json:"id_producto_vp"`
	Descripcion  string  `json:"descripcion"`
	PrecioCompra float32 `json:"precio_compra"`
	PrecioVenta  float32 `json:"precio_venta"`
	Stock        int     `json:"stock"`
	StockMinimo  *int    `json:"stock_minimo"`
}

// Validate validates the UpdateProductoVPRequest fields.
func (m UpdateProductoVPRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 128)),
	)
}

// Validate validates the CreateProductoVPRequest fields.
func (m CreateProductoVPRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 128)),
	)
}

// CrearProductoVP creates a new productoVP.
func (s service) CrearProductoVP(ctx context.Context, req CreateProductoVPRequest) (ProductoVP, error) {
	if err := req.Validate(); err != nil {
		return ProductoVP{}, err
	}
	productoVPG, err := s.repo.CrearProductoVP(ctx, entity.ProductoVP{
		Descripcion:  req.Descripcion,
		PrecioCompra: req.PrecioCompra,
		PrecioVenta:  req.PrecioVenta,
		Stock:        req.Stock,
		StockMinimo:  req.StockMinimo,
	})
	if err != nil {
		return ProductoVP{}, err
	}
	return ProductoVP{productoVPG}, nil
}

// ActualizarProductoVP creates a new productoVP.
func (s service) ActualizarProductoVP(ctx context.Context, req UpdateProductoVPRequest) (ProductoVP, error) {
	if err := req.ValidateUpdate(); err != nil {
		return ProductoVP{}, err
	}
	productoVPG, err := s.repo.ActualizarProductoVP(ctx, entity.ProductoVP{
		IdProductoVP: req.IdProductoVP,
		Descripcion:  req.Descripcion,
		PrecioCompra: req.PrecioCompra,
		PrecioVenta:  req.PrecioVenta,
		Stock:        req.Stock,
		StockMinimo:  req.StockMinimo,
	})
	if err != nil {
		return ProductoVP{}, err
	}
	return ProductoVP{productoVPG}, nil
}

// GetProductoVPPorId returns the productoVP with the specified the productoVP ID.
func (s service) GetProductoVPPorId(ctx context.Context, idProductoVP int) (ProductoVP, error) {
	productoVP, err := s.repo.GetProductoVPPorId(ctx, idProductoVP)
	if err != nil {
		return ProductoVP{}, err
	}
	return ProductoVP{productoVP}, nil
}
