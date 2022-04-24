package proveedor_producto

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access proveedoresProducto from the data source.
type Repository interface {
	// GetProveedorProductoPorId returns the proveedorProducto with the specified proveedorProducto ID.
	GetProveedorProductoPorId(ctx context.Context, idProveedorProducto int) (entity.ProveedorProducto, error)
	// GetProveedoresProducto returns the list proveedoresProducto.
	GetProveedoresProducto(ctx context.Context) ([]entity.ProveedorProducto, error)
	CrearProveedorProducto(ctx context.Context, proveedorProducto entity.ProveedorProducto) (entity.ProveedorProducto, error)
	ActualizarProveedorProducto(ctx context.Context, proveedorProducto entity.ProveedorProducto) (entity.ProveedorProducto, error)
}

// repository persists proveedoresProducto in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new proveedorProducto repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list proveedoresProducto from the database.
func (r repository) GetProveedoresProducto(ctx context.Context) ([]entity.ProveedorProducto, error) {
	var proveedoresProducto []entity.ProveedorProducto

	err := r.db.With(ctx).
		Select().
		From().
		All(&proveedoresProducto)
	if err != nil {
		return proveedoresProducto, err
	}
	return proveedoresProducto, err
}

// Create saves a new ProveedorProducto record in the database.
// It returns the ID of the newly inserted proveedorProducto record.
func (r repository) CrearProveedorProducto(ctx context.Context, proveedorProducto entity.ProveedorProducto) (entity.ProveedorProducto, error) {
	err := r.db.With(ctx).Model(&proveedorProducto).Insert()
	if err != nil {
		return entity.ProveedorProducto{}, err
	}
	return proveedorProducto, nil
}

// Create saves a new ProveedorProducto record in the database.
// It returns the ID of the newly inserted proveedorProducto record.
func (r repository) ActualizarProveedorProducto(ctx context.Context, proveedorProducto entity.ProveedorProducto) (entity.ProveedorProducto, error) {
	var err error
	if proveedorProducto.IdProveedorProducto != 0 {
		err = r.db.With(ctx).Model(&proveedorProducto).Update()
	} else {
		err = r.db.With(ctx).Model(&proveedorProducto).Insert()
	}
	if err != nil {
		return entity.ProveedorProducto{}, err
	}
	return proveedorProducto, nil
}

// Get reads the proveedorProducto with the specified ID from the database.
func (r repository) GetProveedorProductoPorId(ctx context.Context, idProveedorProducto int) (entity.ProveedorProducto, error) {
	var proveedorProducto entity.ProveedorProducto
	err := r.db.With(ctx).Select().Model(idProveedorProducto, &proveedorProducto)
	return proveedorProducto, err
}
