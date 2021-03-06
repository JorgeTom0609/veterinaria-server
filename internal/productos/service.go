package productos

import (
	"context"
	"database/sql"
	"time"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for productos.
type Service interface {
	GetProductos(ctx context.Context) ([]Producto, error)
	GetProductosSinAsignarAProveedor(ctx context.Context, idProveedor int) ([]Producto, error)
	GetProductosStock(ctx context.Context) ([]ProductoStock, error)
	GetProductosCaducados(ctx context.Context) ([]ProductoStock, error)
	GetProductosConStock(ctx context.Context) ([]ProductosConStock, error)
	GetProductosConStockUsoInternoPorServicio(ctx context.Context, idServicio int) ([]ProductosConStock, error)
	GetProductosUsoInterno(ctx context.Context) ([]ProductoUsoInterno, error)
	GetProductosAComparar(ctx context.Context, idProveedor1 int, idProveedor2 int) ([]ProductoComparado, error)
	GetProductoPorId(ctx context.Context, idProducto int) (Producto, error)
	GetProductoCodigoBarra(ctx context.Context, codigoBarra string) (ProductosConStock, error)
	GetProductosPocoStock(ctx context.Context) ([]ProductoPocoStock, error)
	CrearProducto(ctx context.Context, input CreateProductoRequest) (Producto, error)
	ActualizarProducto(ctx context.Context, input UpdateProductoRequest) (Producto, error)
}

// Productos represents the data about an productos.
type Producto struct {
	entity.Producto
}

type ProductoPocoStock struct {
	Producto    string `json:"producto"`
	Stock       int    `json:"stock"`
	StockMinimo int    `json:"stock_minimo"`
}

type ProductoComparado struct {
	Producto string  `json:"producto"`
	Precio1  float32 `json:"precio1"`
	Precio2  float32 `json:"precio2"`
}

type ProductoStock struct {
	Producto          string     `json:"producto"`
	IdLote            int        `json:"id_lote"`
	Lote              string     `json:"lote"`
	FechaCaducidad    *time.Time `json:"fecha_caducidad"`
	Stock             int        `json:"stock"`
	IdStockIndividual *int       `json:"id_stock_individual"`
	StockIndividual   *string    `json:"stock_individual"`
	Cantidad          *float32   `json:"cantidad"`
	Unidad            *string    `json:"unidad"`
}

type ProductosConStock struct {
	StockPorProducto int             `json:"stock_por_producto"`
	Medida           string          `json:"medida"`
	Producto         entity.Producto `json:"producto"`
	Lote             []LoteConStock  `json:"lotes"`
}

type ProductoUsoInterno struct {
	entity.Producto
	Unidad string `json:"unidad"`
}

type LoteConStock struct {
	Lote            entity.Lote              `json:"lote"`
	StockIndividual []entity.StockIndividual `json:"stock_individual"`
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

func (s service) GetProductosStock(ctx context.Context) ([]ProductoStock, error) {
	productos, err := s.repo.GetProductosStock(ctx)
	if err != nil {
		return nil, err
	}
	return productos, nil
}

func (s service) GetProductosCaducados(ctx context.Context) ([]ProductoStock, error) {
	productos, err := s.repo.GetProductosCaducados(ctx)
	if err != nil {
		return nil, err
	}
	return productos, nil
}

func (s service) GetProductosConStock(ctx context.Context) ([]ProductosConStock, error) {
	productos, err := s.repo.GetProductosConStock(ctx)
	if err != nil {
		return nil, err
	}
	return productos, nil
}

func (s service) GetProductosConStockUsoInternoPorServicio(ctx context.Context, idServicio int) ([]ProductosConStock, error) {
	productos, err := s.repo.GetProductosConStockUsoInternoPorServicio(ctx, idServicio)
	if err != nil {
		return nil, err
	}
	return productos, nil
}

func (s service) GetProductosUsoInterno(ctx context.Context) ([]ProductoUsoInterno, error) {
	productos, err := s.repo.GetProductosUsoInterno(ctx)
	if err != nil {
		return nil, err
	}
	return productos, nil
}

func (s service) GetProductosSinAsignarAProveedor(ctx context.Context, idProveedor int) ([]Producto, error) {
	productos, err := s.repo.GetProductosSinAsignarAProveedor(ctx, idProveedor)
	if err != nil {
		return nil, err
	}
	result := []Producto{}
	for _, item := range productos {
		result = append(result, Producto{item})
	}
	return result, nil
}

func (s service) GetProductosAComparar(ctx context.Context, idProveedor1 int, idProveedor2 int) ([]ProductoComparado, error) {
	productos, err := s.repo.GetProductosAComparar(ctx, idProveedor1, idProveedor2)
	if err != nil {
		return nil, err
	}
	return productos, nil
}

// CreateProductoRequest represents an producto creation request.
type CreateProductoRequest struct {
	Descripcion  string       `json:"descripcion"`
	PrecioVenta  float32      `json:"precio_venta"`
	Iva          sql.NullBool `json:"iva"`
	UsoInterno   sql.NullBool `json:"uso_interno"`
	VentaPublico sql.NullBool `json:"venta_publico"`
	PorMedida    sql.NullBool `json:"por_medida"`
	StockMinimo  int          `json:"stock_minimo"`
	Contenido    *float32     `json:"contenido"`
	IdUnidad     *int         `json:"id_unidad"`
}

type UpdateProductoRequest struct {
	IdProducto   int          `json:"id_producto"`
	Descripcion  string       `json:"descripcion"`
	PrecioVenta  float32      `json:"precio_venta"`
	Iva          sql.NullBool `json:"iva"`
	UsoInterno   sql.NullBool `json:"uso_interno"`
	VentaPublico sql.NullBool `json:"venta_publico"`
	PorMedida    sql.NullBool `json:"por_medida"`
	StockMinimo  int          `json:"stock_minimo"`
	Contenido    *float32     `json:"contenido"`
	IdUnidad     *int         `json:"id_unidad"`
}

// Validate validates the UpdateProductoRequest fields.
func (m UpdateProductoRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 1000)),
	)
}

// Validate validates the CreateProductoRequest fields.
func (m CreateProductoRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 1000)),
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
		StockMinimo:  req.StockMinimo,
		IdUnidad:     req.IdUnidad,
		Contenido:    req.Contenido,
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
		StockMinimo:  req.StockMinimo,
		IdUnidad:     req.IdUnidad,
		Contenido:    req.Contenido,
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

func (s service) GetProductoCodigoBarra(ctx context.Context, codigoBarra string) (ProductosConStock, error) {
	producto, err := s.repo.GetProductoCodigoBarra(ctx, codigoBarra)
	if err != nil {
		return ProductosConStock{}, err
	}
	return producto, nil
}

func (s service) GetProductosPocoStock(ctx context.Context) ([]ProductoPocoStock, error) {
	productos, err := s.repo.GetProductosPocoStock(ctx)
	if err != nil {
		return []ProductoPocoStock{}, err
	}
	return productos, nil
}
