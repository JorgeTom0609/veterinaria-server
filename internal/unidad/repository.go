package unidad

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access unidades from the data source.
type Repository interface {
	// GetUnidadPorId returns the unidad with the specified unidad ID.
	GetUnidadPorId(ctx context.Context, idUnidad int) (entity.Unidad, error)
	// GetUnidades returns the list unidades.
	GetUnidades(ctx context.Context) ([]entity.Unidad, error)
	CrearUnidad(ctx context.Context, unidad entity.Unidad) (entity.Unidad, error)
	ActualizarUnidad(ctx context.Context, unidad entity.Unidad) (entity.Unidad, error)
}

// repository persists unidades in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new unidad repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list unidades from the database.
func (r repository) GetUnidades(ctx context.Context) ([]entity.Unidad, error) {
	var unidades []entity.Unidad

	err := r.db.With(ctx).
		Select().
		All(&unidades)
	if err != nil {
		return unidades, err
	}
	return unidades, err
}

// Create saves a new Unidad record in the database.
// It returns the ID of the newly inserted unidad record.
func (r repository) CrearUnidad(ctx context.Context, unidad entity.Unidad) (entity.Unidad, error) {
	err := r.db.With(ctx).Model(&unidad).Insert()
	if err != nil {
		return entity.Unidad{}, err
	}
	return unidad, nil
}

// Create saves a new Unidad record in the database.
// It returns the ID of the newly inserted unidad record.
func (r repository) ActualizarUnidad(ctx context.Context, unidad entity.Unidad) (entity.Unidad, error) {
	var err error
	if unidad.IdUnidad != 0 {
		err = r.db.With(ctx).Model(&unidad).Update()
	} else {
		err = r.db.With(ctx).Model(&unidad).Insert()
	}
	if err != nil {
		return entity.Unidad{}, err
	}
	return unidad, nil
}

// Get reads the unidad with the specified ID from the database.
func (r repository) GetUnidadPorId(ctx context.Context, idUnidad int) (entity.Unidad, error) {
	var unidad entity.Unidad
	err := r.db.With(ctx).Select().Model(idUnidad, &unidad)
	return unidad, err
}
