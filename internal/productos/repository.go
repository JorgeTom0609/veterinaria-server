package productos

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access productos from the data source.
type Repository interface {
	// GetProductoPorId returns the producto with the specified producto ID.
	GetProductoPorId(ctx context.Context, idProducto int) (entity.Producto, error)
	// GetProductos returns the list productos.
	GetProductos(ctx context.Context) ([]entity.Producto, error)
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
