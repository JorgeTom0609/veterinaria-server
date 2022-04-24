package proveedor

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access proveedores from the data source.
type Repository interface {
	// GetProveedorPorId returns the proveedor with the specified proveedor ID.
	GetProveedorPorId(ctx context.Context, idProveedor int) (entity.Proveedor, error)
	// GetProveedores returns the list proveedores.
	GetProveedores(ctx context.Context) ([]entity.Proveedor, error)
	CrearProveedor(ctx context.Context, proveedor entity.Proveedor) (entity.Proveedor, error)
	ActualizarProveedor(ctx context.Context, proveedor entity.Proveedor) (entity.Proveedor, error)
}

// repository persists proveedores in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new proveedor repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list proveedores from the database.
func (r repository) GetProveedores(ctx context.Context) ([]entity.Proveedor, error) {
	var proveedores []entity.Proveedor

	err := r.db.With(ctx).
		Select().
		From().
		All(&proveedores)
	if err != nil {
		return proveedores, err
	}
	return proveedores, err
}

// Create saves a new Proveedor record in the database.
// It returns the ID of the newly inserted proveedor record.
func (r repository) CrearProveedor(ctx context.Context, proveedor entity.Proveedor) (entity.Proveedor, error) {
	err := r.db.With(ctx).Model(&proveedor).Insert()
	if err != nil {
		return entity.Proveedor{}, err
	}
	return proveedor, nil
}

// Create saves a new Proveedor record in the database.
// It returns the ID of the newly inserted proveedor record.
func (r repository) ActualizarProveedor(ctx context.Context, proveedor entity.Proveedor) (entity.Proveedor, error) {
	var err error
	if proveedor.IdProveedor != 0 {
		err = r.db.With(ctx).Model(&proveedor).Update()
	} else {
		err = r.db.With(ctx).Model(&proveedor).Insert()
	}
	if err != nil {
		return entity.Proveedor{}, err
	}
	return proveedor, nil
}

// Get reads the proveedor with the specified ID from the database.
func (r repository) GetProveedorPorId(ctx context.Context, idProveedor int) (entity.Proveedor, error) {
	var proveedor entity.Proveedor
	err := r.db.With(ctx).Select().Model(idProveedor, &proveedor)
	return proveedor, err
}
