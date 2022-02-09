package detalle_examen_informativo

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access detallesExamenInformativo from the data source.
type Repository interface {
	// GetDetalleExamenInformativoPorId returns the detalleExamenInformativo with the specified detalleExamenInformativo ID.
	GetDetalleExamenInformativoPorId(ctx context.Context, idDetalleExamenInformativo int) (entity.DetallesExamenInformativo, error)
	// GetDetallesExamenInformativo returns the list detallesExamenInformativo.
	GetDetallesExamenInformativo(ctx context.Context) ([]entity.DetallesExamenInformativo, error)
	GetDetallesExamenInformativoPorTipoExamen(ctx context.Context, idTipoDeExamen int) ([]entity.DetallesExamenInformativo, error)
	CrearDetalleExamenInformativo(ctx context.Context, detalleExamenInformativo entity.DetallesExamenInformativo) (entity.DetallesExamenInformativo, error)
	ActualizarDetalleExamenInformativo(ctx context.Context, detalleExamenInformativo entity.DetallesExamenInformativo) (entity.DetallesExamenInformativo, error)
}

// repository persists detallesExamenInformativo in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new detalleExamenInformativo repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list detallesExamenInformativo from the database.
func (r repository) GetDetallesExamenInformativo(ctx context.Context) ([]entity.DetallesExamenInformativo, error) {
	var detallesExamenInformativo []entity.DetallesExamenInformativo

	err := r.db.With(ctx).
		Select().
		From().
		All(&detallesExamenInformativo)
	if err != nil {
		return detallesExamenInformativo, err
	}
	return detallesExamenInformativo, err
}

// Create saves a new DetallesExamenInformativo record in the database.
// It returns the ID of the newly inserted detalleExamenInformativo record.
func (r repository) CrearDetalleExamenInformativo(ctx context.Context, detalleExamenInformativo entity.DetallesExamenInformativo) (entity.DetallesExamenInformativo, error) {
	err := r.db.With(ctx).Model(&detalleExamenInformativo).Insert()
	if err != nil {
		return entity.DetallesExamenInformativo{}, err
	}
	return detalleExamenInformativo, nil
}

// Create saves a new DetallesExamenInformativo record in the database.
// It returns the ID of the newly inserted detalleExamenInformativo record.
func (r repository) ActualizarDetalleExamenInformativo(ctx context.Context, detalleExamenInformativo entity.DetallesExamenInformativo) (entity.DetallesExamenInformativo, error) {
	var err error
	if detalleExamenInformativo.IdDetalleExamenInformativo != 0 {
		err = r.db.With(ctx).Model(&detalleExamenInformativo).Update()
	} else {
		err = r.db.With(ctx).Model(&detalleExamenInformativo).Insert()
	}
	if err != nil {
		return entity.DetallesExamenInformativo{}, err
	}
	return detalleExamenInformativo, nil
}

// Get reads the detalleExamenInformativo with the specified ID from the database.
func (r repository) GetDetalleExamenInformativoPorId(ctx context.Context, idDetalleExamenInformativo int) (entity.DetallesExamenInformativo, error) {
	var detalleExamenInformativo entity.DetallesExamenInformativo
	err := r.db.With(ctx).Select().Model(idDetalleExamenInformativo, &detalleExamenInformativo)
	return detalleExamenInformativo, err
}

func (r repository) GetDetallesExamenInformativoPorTipoExamen(ctx context.Context, idTipoDeExamen int) ([]entity.DetallesExamenInformativo, error) {
	var detallesExamenInformativo []entity.DetallesExamenInformativo
	err := r.db.With(ctx).
		Select().
		Where(dbx.HashExp{"id_tipo_examen": idTipoDeExamen}).
		All(&detallesExamenInformativo)
	if err != nil {
		return detallesExamenInformativo, err
	}
	return detallesExamenInformativo, err
}
