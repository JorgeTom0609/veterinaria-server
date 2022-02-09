package detalle_examen_cuantitativo

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access detallesExamenCuantitativo from the data source.
type Repository interface {
	// GetDetalleExamenCuantitativoPorId returns the detalleExamenCuantitativo with the specified detalleExamenCuantitativo ID.
	GetDetalleExamenCuantitativoPorId(ctx context.Context, idDetalleExamenCuantitativo int) (entity.DetallesExamenCuantitativo, error)
	// GetDetallesExamenCuantitativo returns the list detallesExamenCuantitativo.
	GetDetallesExamenCuantitativo(ctx context.Context) ([]entity.DetallesExamenCuantitativo, error)
	GetDetallesExamenCuantitativoPorTipoExamen(ctx context.Context, idTipoDeExamen int) ([]entity.DetallesExamenCuantitativo, error)
	CrearDetalleExamenCuantitativo(ctx context.Context, detalleExamenCuantitativo entity.DetallesExamenCuantitativo) (entity.DetallesExamenCuantitativo, error)
	ActualizarDetalleExamenCuantitativo(ctx context.Context, detalleExamenCuantitativo entity.DetallesExamenCuantitativo) (entity.DetallesExamenCuantitativo, error)
}

// repository persists detallesExamenCuantitativo in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new detalleExamenCuantitativo repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list detallesExamenCuantitativo from the database.
func (r repository) GetDetallesExamenCuantitativo(ctx context.Context) ([]entity.DetallesExamenCuantitativo, error) {
	var detallesExamenCuantitativo []entity.DetallesExamenCuantitativo

	err := r.db.With(ctx).
		Select().
		From().
		All(&detallesExamenCuantitativo)
	if err != nil {
		return detallesExamenCuantitativo, err
	}
	return detallesExamenCuantitativo, err
}

// Create saves a new DetallesExamenCuantitativo record in the database.
// It returns the ID of the newly inserted detalleExamenCuantitativo record.
func (r repository) CrearDetalleExamenCuantitativo(ctx context.Context, detalleExamenCuantitativo entity.DetallesExamenCuantitativo) (entity.DetallesExamenCuantitativo, error) {
	err := r.db.With(ctx).Model(&detalleExamenCuantitativo).Insert()
	if err != nil {
		return entity.DetallesExamenCuantitativo{}, err
	}
	return detalleExamenCuantitativo, nil
}

// Create saves a new DetallesExamenCuantitativo record in the database.
// It returns the ID of the newly inserted detalleExamenCuantitativo record.
func (r repository) ActualizarDetalleExamenCuantitativo(ctx context.Context, detalleExamenCuantitativo entity.DetallesExamenCuantitativo) (entity.DetallesExamenCuantitativo, error) {
	var err error
	if detalleExamenCuantitativo.IdDetalleExamenCuantitativo != 0 {
		err = r.db.With(ctx).Model(&detalleExamenCuantitativo).Update()
	} else {
		err = r.db.With(ctx).Model(&detalleExamenCuantitativo).Insert()
	}
	if err != nil {
		return entity.DetallesExamenCuantitativo{}, err
	}
	return detalleExamenCuantitativo, nil
}

// Get reads the detalleExamenCuantitativo with the specified ID from the database.
func (r repository) GetDetalleExamenCuantitativoPorId(ctx context.Context, idDetalleExamenCuantitativo int) (entity.DetallesExamenCuantitativo, error) {
	var detalleExamenCuantitativo entity.DetallesExamenCuantitativo
	err := r.db.With(ctx).Select().Model(idDetalleExamenCuantitativo, &detalleExamenCuantitativo)
	return detalleExamenCuantitativo, err
}

func (r repository) GetDetallesExamenCuantitativoPorTipoExamen(ctx context.Context, idTipoDeExamen int) ([]entity.DetallesExamenCuantitativo, error) {
	var detallesExamenCuantitativo []entity.DetallesExamenCuantitativo
	err := r.db.With(ctx).
		Select().
		Where(dbx.HashExp{"id_tipo_examen": idTipoDeExamen}).
		All(&detallesExamenCuantitativo)
	if err != nil {
		return detallesExamenCuantitativo, err
	}
	return detallesExamenCuantitativo, err
}
