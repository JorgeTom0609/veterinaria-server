package detalle_examen_cualitativo

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access detallesExamenCualitativo from the data source.
type Repository interface {
	// GetDetalleExamenCualitativoPorId returns the DetalleExamenCualitativo with the specified detalleExamenCualitativo ID.
	GetDetalleExamenCualitativoPorId(ctx context.Context, idDetalleExamenCualitativo int) (entity.DetallesExamenCualitativo, error)
	// GetDetallesExamenCualitativo returns the list detallesExamenCualitativo.
	GetDetallesExamenCualitativo(ctx context.Context) ([]entity.DetallesExamenCualitativo, error)
	GetDetallesExamenCualitativoPorTipoExamen(ctx context.Context, idTipoDeExamen int) ([]entity.DetallesExamenCualitativo, error)
	CrearDetalleExamenCualitativo(ctx context.Context, detalleExamenCualitativo entity.DetallesExamenCualitativo) (entity.DetallesExamenCualitativo, error)
	ActualizarDetalleExamenCualitativo(ctx context.Context, detalleExamenCualitativo entity.DetallesExamenCualitativo) (entity.DetallesExamenCualitativo, error)
}

// repository persists detallesExamenCualitativo in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new detalleExamenCualitativo repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list detallesExamenCualitativo from the database.
func (r repository) GetDetallesExamenCualitativo(ctx context.Context) ([]entity.DetallesExamenCualitativo, error) {
	var detallesExamenCualitativo []entity.DetallesExamenCualitativo

	err := r.db.With(ctx).
		Select().
		From().
		All(&detallesExamenCualitativo)
	if err != nil {
		return detallesExamenCualitativo, err
	}
	return detallesExamenCualitativo, err
}

// Create saves a new DetallesExamenCualitativo record in the database.
// It returns the ID of the newly inserted detalleExamenCualitativo record.
func (r repository) CrearDetalleExamenCualitativo(ctx context.Context, detalleExamenCualitativo entity.DetallesExamenCualitativo) (entity.DetallesExamenCualitativo, error) {
	err := r.db.With(ctx).Model(&detalleExamenCualitativo).Insert()
	if err != nil {
		return entity.DetallesExamenCualitativo{}, err
	}
	return detalleExamenCualitativo, nil
}

// Create saves a new DetallesExamenCualitativo record in the database.
// It returns the ID of the newly inserted detalleExamenCualitativo record.
func (r repository) ActualizarDetalleExamenCualitativo(ctx context.Context, detalleExamenCualitativo entity.DetallesExamenCualitativo) (entity.DetallesExamenCualitativo, error) {
	var err error
	if detalleExamenCualitativo.IdDetalleExamenCualitativo != 0 {
		err = r.db.With(ctx).Model(&detalleExamenCualitativo).Update()
	} else {
		err = r.db.With(ctx).Model(&detalleExamenCualitativo).Insert()
	}
	if err != nil {
		return entity.DetallesExamenCualitativo{}, err
	}
	return detalleExamenCualitativo, nil
}

// GetDetalleExamenCualitativoPorId reads the detalleExamenCualitativo with the specified ID from the database.
func (r repository) GetDetalleExamenCualitativoPorId(ctx context.Context, idDetalleExamenCualitativo int) (entity.DetallesExamenCualitativo, error) {
	var detalleExamenCualitativo entity.DetallesExamenCualitativo
	err := r.db.With(ctx).Select().Model(idDetalleExamenCualitativo, &detalleExamenCualitativo)
	return detalleExamenCualitativo, err
}

func (r repository) GetDetallesExamenCualitativoPorTipoExamen(ctx context.Context, idTipoDeExamen int) ([]entity.DetallesExamenCualitativo, error) {
	var detallesExamenCualitativo []entity.DetallesExamenCualitativo
	err := r.db.With(ctx).
		Select().
		Where(dbx.HashExp{"id_tipo_examen": idTipoDeExamen}).
		All(&detallesExamenCualitativo)
	if err != nil {
		return detallesExamenCualitativo, err
	}
	return detallesExamenCualitativo, err
}
