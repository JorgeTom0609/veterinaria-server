package hospitalizacion

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access hospitalizaciones from the data source.
type Repository interface {
	// GetHospitalizacionPorId returns the hospitalizacion with the specified hospitalizacion ID.
	GetHospitalizacionPorId(ctx context.Context, idHospitalizacion int) (entity.Hospitalizacion, error)
	// GetHospitalizaciones returns the list hospitalizaciones.
	GetHospitalizaciones(ctx context.Context) ([]entity.Hospitalizacion, error)
	CrearHospitalizacion(ctx context.Context, hospitalizacion entity.Hospitalizacion) (entity.Hospitalizacion, error)
	ActualizarHospitalizacion(ctx context.Context, hospitalizacion entity.Hospitalizacion) (entity.Hospitalizacion, error)
}

// repository persists hospitalizaciones in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new hospitalizacion repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list hospitalizaciones from the database.
func (r repository) GetHospitalizaciones(ctx context.Context) ([]entity.Hospitalizacion, error) {
	var hospitalizaciones []entity.Hospitalizacion

	err := r.db.With(ctx).
		Select().
		From().
		All(&hospitalizaciones)
	if err != nil {
		return hospitalizaciones, err
	}
	return hospitalizaciones, err
}

// Create saves a new Hospitalizacion record in the database.
// It returns the ID of the newly inserted hospitalizacion record.
func (r repository) CrearHospitalizacion(ctx context.Context, hospitalizacion entity.Hospitalizacion) (entity.Hospitalizacion, error) {
	err := r.db.With(ctx).Model(&hospitalizacion).Insert()
	if err != nil {
		return entity.Hospitalizacion{}, err
	}
	return hospitalizacion, nil
}

// Create saves a new Hospitalizacion record in the database.
// It returns the ID of the newly inserted hospitalizacion record.
func (r repository) ActualizarHospitalizacion(ctx context.Context, hospitalizacion entity.Hospitalizacion) (entity.Hospitalizacion, error) {
	var err error
	if hospitalizacion.IdHospitalizacion != 0 {
		err = r.db.With(ctx).Model(&hospitalizacion).Update()
	} else {
		err = r.db.With(ctx).Model(&hospitalizacion).Insert()
	}
	if err != nil {
		return entity.Hospitalizacion{}, err
	}
	return hospitalizacion, nil
}

// Get reads the hospitalizacion with the specified ID from the database.
func (r repository) GetHospitalizacionPorId(ctx context.Context, idHospitalizacion int) (entity.Hospitalizacion, error) {
	var hospitalizacion entity.Hospitalizacion
	err := r.db.With(ctx).Select().Model(idHospitalizacion, &hospitalizacion)
	return hospitalizacion, err
}
