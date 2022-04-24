package productos

import (
	"context"
	"database/sql"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for productos.
type Service interface {
	GetProductos(ctx context.Context) ([]Producto, error)
	GetProductoPorId(ctx context.Context, idProducto int) (Producto, error)
	CrearProducto(ctx context.Context, input CreateProductoRequest) (Producto, error)
	ActualizarProducto(ctx context.Context, input UpdateProductoRequest) (Producto, error)
}

// Productos represents the data about an productos.
type Producto struct {
	entity.Producto
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new productos service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list productos.
func (s service) GetProductos(ctx context.Context) ([]Producto, error) {
	productos, err := s.repo.GetProductos(ctx)
	if err != nil {
		return nil, err
	}
	result := []Producto{}
	for _, item := range productos {
		result = append(result, Producto{item})
	}
	return result, nil
}

// CreateProductoRequest represents an producto creation request.
type CreateProductoRequest struct {
	Descripcion  string       `json:"descripcion"`
	PrecioVenta  float32      `json:"precio_venta"`
	Iva          sql.NullBool `json:"iva"`
	UsoInterno   sql.NullBool `json:"uso_interno"`
	VentaPublico sql.NullBool `json:"venta_publico"`
	PorMedida    sql.NullBool `json:"por_medida"`
}

type UpdateProductoRequest struct {
	IdProducto   int          `json:"id_producto"`
	Descripcion  string       `json:"descripcion"`
	PrecioVenta  float32      `json:"precio_venta"`
	Iva          sql.NullBool `json:"iva"`
	UsoInterno   sql.NullBool `json:"uso_interno"`
	VentaPublico sql.NullBool `json:"venta_publico"`
	PorMedida    sql.NullBool `json:"por_medida"`
}

// Validate validates the UpdateProductoRequest fields.
func (m UpdateProductoRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 128)),
	)
}

// Validate validates the CreateProductoRequest fields.
func (m CreateProductoRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 128)),
	)
}

// CrearProducto creates a new producto.
func (s service) CrearProducto(ctx context.Context, req CreateProductoRequest) (Producto, error) {
	if err := req.Validate(); err != nil {
		return Producto{}, err
	}
	productoG, err := s.repo.CrearProducto(ctx, entity.Producto{
		Descripcion:  req.Descripcion,
		PrecioVenta:  req.PrecioVenta,
		Iva:          req.Iva,
		UsoInterno:   req.UsoInterno,
		VentaPublico: req.VentaPublico,
		PorMedida:    req.PorMedida,
	})
	if err != nil {
		return Producto{}, err
	}
	return Producto{productoG}, nil
}

// ActualizarProducto creates a new producto.
func (s service) ActualizarProducto(ctx context.Context, req UpdateProductoRequest) (Producto, error) {
	if err := req.ValidateUpdate(); err != nil {
		return Producto{}, err
	}
	productoG, err := s.repo.ActualizarProducto(ctx, entity.Producto{
		IdProducto:   req.IdProducto,
		Descripcion:  req.Descripcion,
		PrecioVenta:  req.PrecioVenta,
		Iva:          req.Iva,
		UsoInterno:   req.UsoInterno,
		VentaPublico: req.VentaPublico,
		PorMedida:    req.PorMedida,
	})
	if err != nil {
		return Producto{}, err
	}
	return Producto{productoG}, nil
}

// GetProductoPorId returns the producto with the specified the producto ID.
func (s service) GetProductoPorId(ctx context.Context, idProducto int) (Producto, error) {
	producto, err := s.repo.GetProductoPorId(ctx, idProducto)
	if err != nil {
		return Producto{}, err
	}
	return Producto{producto}, nil
}
