package detalle_hospitalizacion

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access detallesHospitalizacion from the data source.
type Repository interface {
	// GetDetalleHospitalizacionPorId returns the detalleHospitalizacion with the specified detalleHospitalizacion ID.
	GetDetalleHospitalizacionPorId(ctx context.Context, idDetalleHospitalizacion int) (entity.DetalleHospitalizacion, error)
	// GetDetallesHospitalizacion returns the list detallesHospitalizacion.
	GetDetallesHospitalizacion(ctx context.Context) ([]entity.DetalleHospitalizacion, error)
	CrearDetalleHospitalizacion(ctx context.Context, detalleHospitalizacion entity.DetalleHospitalizacion) (entity.DetalleHospitalizacion, error)
	ActualizarDetalleHospitalizacion(ctx context.Context, detalleHospitalizacion entity.DetalleHospitalizacion) (entity.DetalleHospitalizacion, error)
}

// repository persists detallesHospitalizacion in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new detalleHospitalizacion repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list detallesHospitalizacion from the database.
func (r repository) GetDetallesHospitalizacion(ctx context.Context) ([]entity.DetalleHospitalizacion, error) {
	var detallesHospitalizacion []entity.DetalleHospitalizacion

	err := r.db.With(ctx).
		Select().
		From().
		All(&detallesHospitalizacion)
	if err != nil {
		return detallesHospitalizacion, err
	}
	return detallesHospitalizacion, err
}

// Create saves a new DetalleHospitalizacion record in the database.
// It returns the ID of the newly inserted detalleHospitalizacion record.
func (r repository) CrearDetalleHospitalizacion(ctx context.Context, detalleHospitalizacion entity.DetalleHospitalizacion) (entity.DetalleHospitalizacion, error) {
	err := r.db.With(ctx).Model(&detalleHospitalizacion).Insert()
	if err != nil {
		return entity.DetalleHospitalizacion{}, err
	}
	return detalleHospitalizacion, nil
}

// Create saves a new DetalleHospitalizacion record in the database.
// It returns the ID of the newly inserted detalleHospitalizacion record.
func (r repository) ActualizarDetalleHospitalizacion(ctx context.Context, detalleHospitalizacion entity.DetalleHospitalizacion) (entity.DetalleHospitalizacion, error) {
	var err error
	if detalleHospitalizacion.IdDetalleHospitalizacion != 0 {
		err = r.db.With(ctx).Model(&detalleHospitalizacion).Update()
	} else {
		err = r.db.With(ctx).Model(&detalleHospitalizacion).Insert()
	}
	if err != nil {
		return entity.DetalleHospitalizacion{}, err
	}
	return detalleHospitalizacion, nil
}

// Get reads the detalleHospitalizacion with the specified ID from the database.
func (r repository) GetDetalleHospitalizacionPorId(ctx context.Context, idDetalleHospitalizacion int) (entity.DetalleHospitalizacion, error) {
	var detalleHospitalizacion entity.DetalleHospitalizacion
	err := r.db.With(ctx).Select().Model(idDetalleHospitalizacion, &detalleHospitalizacion)
	return detalleHospitalizacion, err
}
