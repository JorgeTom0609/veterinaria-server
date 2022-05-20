package detalle_servicio_hospitalizacion

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access detallesServicioHospitalizacion from the data source.
type Repository interface {
	// GetDetalleServicioHospitalizacionPorId returns the detalleServicioHospitalizacion with the specified detalleServicioHospitalizacion ID.
	GetDetalleServicioHospitalizacionPorId(ctx context.Context, idDetalleServicioHospitalizacion int) (entity.DetalleServicioHospitalizacion, error)
	// GetDetallesServicioHospitalizacion returns the list detallesServicioHospitalizacion.
	GetDetallesServicioHospitalizacion(ctx context.Context) ([]entity.DetalleServicioHospitalizacion, error)
	CrearDetalleServicioHospitalizacion(ctx context.Context, detalleServicioHospitalizacion entity.DetalleServicioHospitalizacion) (entity.DetalleServicioHospitalizacion, error)
	ActualizarDetalleServicioHospitalizacion(ctx context.Context, detalleServicioHospitalizacion entity.DetalleServicioHospitalizacion) (entity.DetalleServicioHospitalizacion, error)
}

// repository persists detallesServicioHospitalizacion in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new detalleServicioHospitalizacion repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list detallesServicioHospitalizacion from the database.
func (r repository) GetDetallesServicioHospitalizacion(ctx context.Context) ([]entity.DetalleServicioHospitalizacion, error) {
	var detallesServicioHospitalizacion []entity.DetalleServicioHospitalizacion

	err := r.db.With(ctx).
		Select().
		From().
		All(&detallesServicioHospitalizacion)
	if err != nil {
		return detallesServicioHospitalizacion, err
	}
	return detallesServicioHospitalizacion, err
}

// Create saves a new DetalleServicioHospitalizacion record in the database.
// It returns the ID of the newly inserted detalleServicioHospitalizacion record.
func (r repository) CrearDetalleServicioHospitalizacion(ctx context.Context, detalleServicioHospitalizacion entity.DetalleServicioHospitalizacion) (entity.DetalleServicioHospitalizacion, error) {
	err := r.db.With(ctx).Model(&detalleServicioHospitalizacion).Insert()
	if err != nil {
		return entity.DetalleServicioHospitalizacion{}, err
	}
	return detalleServicioHospitalizacion, nil
}

// Create saves a new DetalleServicioHospitalizacion record in the database.
// It returns the ID of the newly inserted detalleServicioHospitalizacion record.
func (r repository) ActualizarDetalleServicioHospitalizacion(ctx context.Context, detalleServicioHospitalizacion entity.DetalleServicioHospitalizacion) (entity.DetalleServicioHospitalizacion, error) {
	var err error
	if detalleServicioHospitalizacion.IdDetalleServicioHospitalizacion != 0 {
		err = r.db.With(ctx).Model(&detalleServicioHospitalizacion).Update()
	} else {
		err = r.db.With(ctx).Model(&detalleServicioHospitalizacion).Insert()
	}
	if err != nil {
		return entity.DetalleServicioHospitalizacion{}, err
	}
	return detalleServicioHospitalizacion, nil
}

// Get reads the detalleServicioHospitalizacion with the specified ID from the database.
func (r repository) GetDetalleServicioHospitalizacionPorId(ctx context.Context, idDetalleServicioHospitalizacion int) (entity.DetalleServicioHospitalizacion, error) {
	var detalleServicioHospitalizacion entity.DetalleServicioHospitalizacion
	err := r.db.With(ctx).Select().Model(idDetalleServicioHospitalizacion, &detalleServicioHospitalizacion)
	return detalleServicioHospitalizacion, err
}
