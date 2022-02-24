package resultado_examen_cualitativo

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access ResultadoDetalleCualitativo from the data source.
type Repository interface {
	// GetResultadoDetalleCualitativoPorId returns the resultadoDetalleCualitativo with the specified resultadoDetalleCualitativo ID.
	GetResultadoDetalleCualitativoPorId(ctx context.Context, idResultadoDetalleCualitativo int) (entity.ResultadoDetalleCualitativo, error)
	CrearResultadoDetalleCualitativo(ctx context.Context, resultadoDetalleCualitativo entity.ResultadoDetalleCualitativo) (entity.ResultadoDetalleCualitativo, error)
}

// repository persists ResultadoDetalleCualitativo in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new resultadoDetalleCualitativo repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Create saves a new ResultadoDetalleCualitativo record in the database.
// It returns the ID of the newly inserted resultadoDetalleCualitativo record.
func (r repository) CrearResultadoDetalleCualitativo(ctx context.Context, resultadoDetalleCualitativo entity.ResultadoDetalleCualitativo) (entity.ResultadoDetalleCualitativo, error) {
	err := r.db.With(ctx).Model(&resultadoDetalleCualitativo).Insert()
	if err != nil {
		return entity.ResultadoDetalleCualitativo{}, err
	}
	return resultadoDetalleCualitativo, nil
}

// Get reads the resultadoDetalleCualitativo with the specified ID from the database.
func (r repository) GetResultadoDetalleCualitativoPorId(ctx context.Context, idResultadoDetalleCualitativo int) (entity.ResultadoDetalleCualitativo, error) {
	var resultadoDetalleCualitativo entity.ResultadoDetalleCualitativo
	err := r.db.With(ctx).Select().Model(idResultadoDetalleCualitativo, &resultadoDetalleCualitativo)
	return resultadoDetalleCualitativo, err
}
