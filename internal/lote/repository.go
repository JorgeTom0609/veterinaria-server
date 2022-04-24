package lote

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access lotes from the data source.
type Repository interface {
	// GetLotePorId returns the lote with the specified lote ID.
	GetLotePorId(ctx context.Context, idLote int) (entity.Lote, error)
	// GetLotes returns the list lotes.
	GetLotes(ctx context.Context) ([]entity.Lote, error)
	CrearLote(ctx context.Context, lote entity.Lote) (entity.Lote, error)
	ActualizarLote(ctx context.Context, lote entity.Lote) (entity.Lote, error)
}

// repository persists lotes in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new lote repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list lotes from the database.
func (r repository) GetLotes(ctx context.Context) ([]entity.Lote, error) {
	var lotes []entity.Lote

	err := r.db.With(ctx).
		Select().
		All(&lotes)
	if err != nil {
		return lotes, err
	}
	return lotes, err
}

// Create saves a new Lote record in the database.
// It returns the ID of the newly inserted lote record.
func (r repository) CrearLote(ctx context.Context, lote entity.Lote) (entity.Lote, error) {
	err := r.db.With(ctx).Model(&lote).Insert()
	if err != nil {
		return entity.Lote{}, err
	}
	return lote, nil
}

// Create saves a new Lote record in the database.
// It returns the ID of the newly inserted lote record.
func (r repository) ActualizarLote(ctx context.Context, lote entity.Lote) (entity.Lote, error) {
	var err error
	if lote.IdLote != 0 {
		err = r.db.With(ctx).Model(&lote).Update()
	} else {
		err = r.db.With(ctx).Model(&lote).Insert()
	}
	if err != nil {
		return entity.Lote{}, err
	}
	return lote, nil
}

// Get reads the Lote with the specified ID from the database.
func (r repository) GetLotePorId(ctx context.Context, idLote int) (entity.Lote, error) {
	var lote entity.Lote
	err := r.db.With(ctx).Select().Model(idLote, &lote)
	return lote, err
}
