package detalle_uso_servicio

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access detallesUsoServicio from the data source.
type Repository interface {
	// GetDetalleUsoServicioPorId returns the detalleUsoServicio with the specified detalleUsoServicio ID.
	GetDetalleUsoServicioPorId(ctx context.Context, idDetalleUsoServicio int) (entity.DetalleUsoServicio, error)
	// GetDetallesUsoServicio returns the list detallesUsoServicio.
	GetDetallesUsoServicio(ctx context.Context) ([]entity.DetalleUsoServicio, error)
	CrearDetalleUsoServicio(ctx context.Context, detalleUsoServicio entity.DetalleUsoServicio) (entity.DetalleUsoServicio, error)
	ActualizarDetalleUsoServicio(ctx context.Context, detalleUsoServicio entity.DetalleUsoServicio) (entity.DetalleUsoServicio, error)
}

// repository persists detallesUsoServicio in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new detalleUsoServicio repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list detallesUsoServicio from the database.
func (r repository) GetDetallesUsoServicio(ctx context.Context) ([]entity.DetalleUsoServicio, error) {
	var detallesUsoServicio []entity.DetalleUsoServicio

	err := r.db.With(ctx).
		Select().
		From().
		All(&detallesUsoServicio)
	if err != nil {
		return detallesUsoServicio, err
	}
	return detallesUsoServicio, err
}

// Create saves a new DetalleUsoServicio record in the database.
// It returns the ID of the newly inserted detalleUsoServicio record.
func (r repository) CrearDetalleUsoServicio(ctx context.Context, detalleUsoServicio entity.DetalleUsoServicio) (entity.DetalleUsoServicio, error) {
	err := r.db.With(ctx).Model(&detalleUsoServicio).Insert()
	if err != nil {
		return entity.DetalleUsoServicio{}, err
	}
	return detalleUsoServicio, nil
}

// Create saves a new DetalleUsoServicio record in the database.
// It returns the ID of the newly inserted detalleUsoServicio record.
func (r repository) ActualizarDetalleUsoServicio(ctx context.Context, detalleUsoServicio entity.DetalleUsoServicio) (entity.DetalleUsoServicio, error) {
	var err error
	if detalleUsoServicio.IdDetalleUsoServicio != 0 {
		err = r.db.With(ctx).Model(&detalleUsoServicio).Update()
	} else {
		err = r.db.With(ctx).Model(&detalleUsoServicio).Insert()
	}
	if err != nil {
		return entity.DetalleUsoServicio{}, err
	}
	return detalleUsoServicio, nil
}

// Get reads the detalleUsoServicio with the specified ID from the database.
func (r repository) GetDetalleUsoServicioPorId(ctx context.Context, idDetalleUsoServicio int) (entity.DetalleUsoServicio, error) {
	var detalleUsoServicio entity.DetalleUsoServicio
	err := r.db.With(ctx).Select().Model(idDetalleUsoServicio, &detalleUsoServicio)
	return detalleUsoServicio, err
}
