package productos

import (
	"context"
	"strconv"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access productos from the data source.
type Repository interface {
	// GetProductoPorId returns the producto with the specified producto ID.
	GetProductoPorId(ctx context.Context, idProducto int) (entity.Producto, error)
	// GetProductos returns the list productos.
	GetProductos(ctx context.Context) ([]entity.Producto, error)
	GetProductosConStock(ctx context.Context) ([]ProductosConStock, error)
	GetProductosSinAsignarAProveedor(ctx context.Context, idProveedor int) ([]entity.Producto, error)
	CrearProducto(ctx context.Context, producto entity.Producto) (entity.Producto, error)
	ActualizarProducto(ctx context.Context, producto entity.Producto) (entity.Producto, error)
}

// repository persists productos in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new producto repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list productos from the database.
func (r repository) GetProductos(ctx context.Context) ([]entity.Producto, error) {
	var productos []entity.Producto

	err := r.db.With(ctx).
		Select().
		From().
		All(&productos)
	if err != nil {
		return productos, err
	}
	return productos, err
}

func (r repository) GetProductosConStock(ctx context.Context) ([]ProductosConStock, error) {
	var productos []entity.Producto
	var lotes []entity.Lote
	var stockIndividuales []entity.StockIndividual
	var stockPorProducto int

	var productoConStock ProductosConStock
	var productosConStock []ProductosConStock

	err := r.db.With(ctx).
		Select().
		Where(dbx.NewExp("venta_publico = true")).
		All(&productos)
	if err != nil {
		return []ProductosConStock{}, err
	}
	for i := 0; i < len(productos); i++ {
		productoConStock = ProductosConStock{}
		stockPorProducto = 0
		lotes = []entity.Lote{}
		err := r.db.With(ctx).
			Select().
			From("lote l").
			InnerJoin("proveedor_producto as pp", dbx.NewExp("pp.id_proveedor_producto = l.id_proveedor_producto")).
			Where(dbx.HashExp{"pp.id_producto": productos[i].IdProducto}).
			AndWhere(dbx.NewExp("(DATE(now()) <= fecha_caducidad or fecha_caducidad is null) and stock > 0")).
			All(&lotes)
		if err != nil {
			return []ProductosConStock{}, err
		}

		for j := 0; j < len(lotes); j++ {
			productoConStock.Lote = append(productoConStock.Lote, LoteConStock{})
			productoConStock.Lote[j].Lote = lotes[j]
			stockIndividuales = []entity.StockIndividual{}
			if productos[i].PorMedida.Bool {
				err := r.db.With(ctx).
					Select().
					Where(dbx.HashExp{"id_lote": lotes[j].IdLote}).
					AndWhere(dbx.NewExp("cantidad > 0")).
					All(&stockIndividuales)
				if err != nil {
					return []ProductosConStock{}, err
				}
				productoConStock.Lote[j].StockIndividual = stockIndividuales
			}
		}
		if len(lotes) > 0 {
			err = r.db.With(ctx).
				Select("sum(stock)").
				From("lote l").
				InnerJoin("proveedor_producto as pp", dbx.NewExp("pp.id_proveedor_producto = l.id_proveedor_producto")).
				Where(dbx.HashExp{"pp.id_producto": productos[i].IdProducto}).
				AndWhere(dbx.NewExp("(DATE(now()) <= fecha_caducidad or fecha_caducidad is null) and stock > 0")).
				Row(&stockPorProducto)
			if err != nil {
				return []ProductosConStock{}, err
			}
			productoConStock.StockPorProducto = stockPorProducto
			productoConStock.Producto = productos[i]
			productosConStock = append(productosConStock, productoConStock)
		}

	}
	return productosConStock, err
}

func (r repository) GetProductosSinAsignarAProveedor(ctx context.Context, idProveedor int) ([]entity.Producto, error) {
	var productos []entity.Producto
	var proveedorProductos []entity.ProveedorProducto
	var idsProductos string = ""

	err := r.db.With(ctx).
		Select().
		From().
		Where(dbx.HashExp{"id_proveedor": idProveedor}).
		All(&proveedorProductos)

	for i := 0; i < len(proveedorProductos); i++ {
		if i == 0 {
			idsProductos = idsProductos + strconv.Itoa(proveedorProductos[i].IdProducto)
		} else {
			idsProductos = idsProductos + ", " + strconv.Itoa(proveedorProductos[i].IdProducto)
		}
	}

	if len(proveedorProductos) > 0 {
		err = r.db.With(ctx).
			Select().
			From().
			Where(dbx.NewExp("id_producto not in (" + idsProductos + ")")).
			All(&productos)
	} else {
		err = r.db.With(ctx).
			Select().
			From().
			All(&productos)
	}

	if err != nil {
		return productos, err
	}
	return productos, err
}

// Create saves a new Producto record in the database.
// It returns the ID of the newly inserted producto record.
func (r repository) CrearProducto(ctx context.Context, producto entity.Producto) (entity.Producto, error) {
	err := r.db.With(ctx).Model(&producto).Insert()
	if err != nil {
		return entity.Producto{}, err
	}
	return producto, nil
}

// Create saves a new Producto record in the database.
// It returns the ID of the newly inserted producto record.
func (r repository) ActualizarProducto(ctx context.Context, producto entity.Producto) (entity.Producto, error) {
	var err error
	if producto.IdProducto != 0 {
		err = r.db.With(ctx).Model(&producto).Update()
	} else {
		err = r.db.With(ctx).Model(&producto).Insert()
	}
	if err != nil {
		return entity.Producto{}, err
	}
	return producto, nil
}

// Get reads the producto with the specified ID from the database.
func (r repository) GetProductoPorId(ctx context.Context, idProducto int) (entity.Producto, error) {
	var producto entity.Producto
	err := r.db.With(ctx).Select().Model(idProducto, &producto)
	return producto, err
}
