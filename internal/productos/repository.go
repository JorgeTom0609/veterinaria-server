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
	GetProductosCaducados(ctx context.Context) ([]ProductoStock, error)
	GetProductosStock(ctx context.Context) ([]ProductoStock, error)
	GetProductosConStock(ctx context.Context) ([]ProductosConStock, error)
	GetProductosConStockUsoInternoPorServicio(ctx context.Context, idServicio int) ([]ProductosConStock, error)
	GetProductosUsoInterno(ctx context.Context) ([]ProductoUsoInterno, error)
	GetProductosAComparar(ctx context.Context, idProveedor1 int, idProveedor2 int) ([]ProductoComparado, error)
	GetProductosSinAsignarAProveedor(ctx context.Context, idProveedor int) ([]entity.Producto, error)
	GetProductoCodigoBarra(ctx context.Context, codigoBarra string) (ProductosConStock, error)
	GetProductosPocoStock(ctx context.Context) ([]ProductoPocoStock, error)
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

func (r repository) GetProductosUsoInterno(ctx context.Context) ([]ProductoUsoInterno, error) {
	var productos []entity.Producto
	productosUsoInterno := []ProductoUsoInterno{}

	err := r.db.With(ctx).
		Select().
		Where(dbx.NewExp("uso_interno = true")).
		All(&productos)
	if err != nil {
		return []ProductoUsoInterno{}, err
	}

	for i := 0; i < len(productos); i++ {
		var unidad entity.Unidad
		if productos[i].PorMedida.Bool {
			err := r.db.With(ctx).
				Select().
				Where(dbx.HashExp{"id_unidad": productos[i].IdUnidad}).
				One(&unidad)
			if err != nil {
				return []ProductoUsoInterno{}, err
			}
			productosUsoInterno = append(productosUsoInterno, ProductoUsoInterno{Producto: productos[i], Unidad: unidad.Descripcion})
		} else {
			productosUsoInterno = append(productosUsoInterno, ProductoUsoInterno{Producto: productos[i], Unidad: ""})
		}
	}
	return productosUsoInterno, err
}

func (r repository) GetProductosStock(ctx context.Context) ([]ProductoStock, error) {
	var productosStock []ProductoStock = []ProductoStock{}
	err := r.db.With(ctx).
		Select("p.descripcion as producto", "l.id_lote", "l.descripcion as lote", "l.fecha_caducidad", "l.stock", "si.id_stock_individual", "si.descripcion as stock_individual", "si.cantidad", "u.descripcion as unidad").
		From("producto p").
		LeftJoin("unidad u", dbx.NewExp("u.id_unidad = p.id_unidad")).
		InnerJoin("proveedor_producto pp", dbx.NewExp("pp.id_producto = p.id_producto")).
		InnerJoin("lote l", dbx.NewExp("l.id_proveedor_producto = pp.id_proveedor_producto")).
		LeftJoin("stock_individual si", dbx.NewExp("si.id_lote = l.id_lote and cantidad > 0")).
		Where(dbx.NewExp("(DATE(now()) <= l.fecha_caducidad or l.fecha_caducidad is null) and l.stock > 0")).
		OrderBy("fecha_caducidad asc").
		All(&productosStock)
	if err != nil {
		return []ProductoStock{}, err
	}
	return productosStock, err
}

func (r repository) GetProductosCaducados(ctx context.Context) ([]ProductoStock, error) {
	var productosStock []ProductoStock = []ProductoStock{}
	err := r.db.With(ctx).
		Select("p.descripcion as producto", "l.id_lote", "l.descripcion as lote", "l.fecha_caducidad", "l.stock", "si.id_stock_individual", "si.descripcion as stock_individual", "si.cantidad", "u.descripcion as unidad").
		From("producto p").
		LeftJoin("unidad u", dbx.NewExp("u.id_unidad = p.id_unidad")).
		InnerJoin("proveedor_producto pp", dbx.NewExp("pp.id_producto = p.id_producto")).
		InnerJoin("lote l", dbx.NewExp("l.id_proveedor_producto = pp.id_proveedor_producto")).
		LeftJoin("stock_individual si", dbx.NewExp("si.id_lote = l.id_lote and cantidad > 0")).
		Where(dbx.NewExp("(DATE(now()) > l.fecha_caducidad) and l.stock > 0")).
		OrderBy("fecha_caducidad asc").
		All(&productosStock)
	if err != nil {
		return []ProductoStock{}, err
	}
	return productosStock, err
}

func (r repository) GetProductosConStock(ctx context.Context) ([]ProductosConStock, error) {
	var productos []entity.Producto
	var lotes []entity.Lote
	var stockIndividuales []entity.StockIndividual
	var stockPorProducto int
	var medida string

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
		medida = ""
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

				err = r.db.With(ctx).
					Select("descripcion").
					From("Unidad").
					Where(dbx.HashExp{"id_unidad": productos[i].IdUnidad}).
					Row(&medida)
				if err != nil {
					return []ProductosConStock{}, err
				}
				productoConStock.Medida = medida

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

func (r repository) GetProductosConStockUsoInternoPorServicio(ctx context.Context, idServicio int) ([]ProductosConStock, error) {
	var productos []entity.Producto
	var lotes []entity.Lote
	var stockIndividuales []entity.StockIndividual
	var stockPorProducto int

	var productoConStock ProductosConStock
	var productosConStock []ProductosConStock

	err := r.db.With(ctx).
		Select().
		From("producto p").
		InnerJoin("servicio_producto as sp", dbx.NewExp("sp.id_producto = p.id_producto and sp.estado = 'A'")).
		Where(dbx.NewExp("p.uso_interno = true")).
		AndWhere(dbx.HashExp{"sp.id_servicio": idServicio}).
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

func (r repository) GetProductosAComparar(ctx context.Context, idProveedor1 int, idProveedor2 int) ([]ProductoComparado, error) {
	var productosComparados []ProductoComparado
	err := r.db.With(ctx).
		Select("pr.descripcion as producto", "pp1.precio_compra as precio1", "pp2.precio_compra as precio2").
		From("producto pr").
		InnerJoin("proveedor_producto as pp1", dbx.NewExp("pp1.id_producto = pr.id_producto and pp1.id_proveedor = "+strconv.Itoa(idProveedor1))).
		InnerJoin("proveedor_producto as pp2", dbx.NewExp("pp2.id_producto = pr.id_producto and pp2.id_proveedor = "+strconv.Itoa(idProveedor2))).
		All(&productosComparados)
	if err != nil {
		return []ProductoComparado{}, err
	}
	return productosComparados, err
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

func (r repository) GetProductoCodigoBarra(ctx context.Context, codigoBarra string) (ProductosConStock, error) {
	var producto entity.Producto
	var lote entity.Lote
	var stockIndividuales []entity.StockIndividual
	var stockPorProducto int
	var medida string

	var productoConStock ProductosConStock
	productoConStock = ProductosConStock{}
	stockPorProducto = 0

	err := r.db.With(ctx).
		Select().
		From("lote l").
		Where(dbx.HashExp{"l.codigo_barra": codigoBarra}).
		AndWhere(dbx.NewExp("(DATE(now()) <= fecha_caducidad or fecha_caducidad is null) and stock > 0")).
		One(&lote)
	if err != nil {
		return ProductosConStock{}, err
	}
	err = r.db.With(ctx).
		Select().
		From("proveedor_producto pp").
		InnerJoin("producto as p", dbx.NewExp("p.id_producto = pp.id_producto")).
		Where(dbx.HashExp{"pp.id_proveedor_producto": lote.IdProveedorProducto}).
		AndWhere(dbx.HashExp{"p.venta_publico": true}).
		One(&producto)
	if err != nil {
		return ProductosConStock{}, err
	}

	productoConStock.Lote = append(productoConStock.Lote, LoteConStock{})
	productoConStock.Lote[0].Lote = lote
	stockIndividuales = []entity.StockIndividual{}
	if producto.PorMedida.Bool {
		err := r.db.With(ctx).
			Select().
			Where(dbx.HashExp{"id_lote": lote.IdLote}).
			AndWhere(dbx.NewExp("cantidad > 0")).
			All(&stockIndividuales)
		if err != nil {
			return ProductosConStock{}, err
		}
		productoConStock.Lote[0].StockIndividual = stockIndividuales

		err = r.db.With(ctx).
			Select("descripcion").
			From("Unidad").
			Where(dbx.HashExp{"id_unidad": producto.IdUnidad}).
			Row(&medida)
		if err != nil {
			return ProductosConStock{}, err
		}
		productoConStock.Medida = medida

	}

	err = r.db.With(ctx).
		Select("sum(stock)").
		From("lote l").
		InnerJoin("proveedor_producto as pp", dbx.NewExp("pp.id_proveedor_producto = l.id_proveedor_producto")).
		Where(dbx.HashExp{"pp.id_producto": producto.IdProducto}).
		AndWhere(dbx.NewExp("(DATE(now()) <= fecha_caducidad or fecha_caducidad is null) and stock > 0")).
		Row(&stockPorProducto)
	if err != nil {
		return ProductosConStock{}, err
	}
	productoConStock.StockPorProducto = stockPorProducto
	productoConStock.Producto = producto
	return productoConStock, err
}

func (r repository) GetProductosPocoStock(ctx context.Context) ([]ProductoPocoStock, error) {

	var productoPocoStock []ProductoPocoStock

	err := r.db.With(ctx).
		Select("p.descripcion as producto", "sum(stock) as stock", "p.stock_minimo").
		From("lote l").
		InnerJoin("proveedor_producto as pp", dbx.NewExp("pp.id_proveedor_producto = l.id_proveedor_producto")).
		InnerJoin("producto as p", dbx.NewExp("p.id_producto = pp.id_producto")).
		Where(dbx.NewExp("(DATE(now()) <= fecha_caducidad or fecha_caducidad is null) and stock > 0")).
		GroupBy("p.id_producto").
		Having(dbx.NewExp("sum(stock) <= p.stock_minimo")).
		All(&productoPocoStock)
	if err != nil {
		return []ProductoPocoStock{}, err
	}
	return productoPocoStock, err
}
