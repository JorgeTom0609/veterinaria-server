package resultado_examen_cuantitativo

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access ResultadoDetalleCuantitativo from the data source.
type Repository interface {
	// GetResultadoDetalleCuantitativoPorId returns the resultadoDetalleCuantitativo with the specified resultadoDetalleCuantitativo ID.
	GetResultadoDetalleCuantitativoPorId(ctx context.Context, idResultadoDetalleCuantitativo int) (entity.ResultadoDetalleCuantitativo, error)
	CrearResultadoDetalleCuantitativo(ctx context.Context, resultadoDetalleCuantitativo entity.ResultadoDetalleCuantitativo) (entity.ResultadoDetalleCuantitativo, error)
}

// repository persists ResultadoDetalleCuantitativo in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new resultadoDetalleCuantitativo repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Create saves a new ResultadoDetalleCuantitativo record in the database.
// It returns the ID of the newly inserted resultadoDetalleCuantitativo record.
func (r repository) CrearResultadoDetalleCuantitativo(ctx context.Context, resultadoDetalleCuantitativo entity.ResultadoDetalleCuantitativo) (entity.ResultadoDetalleCuantitativo, error) {
	err := r.db.With(ctx).Model(&resultadoDetalleCuantitativo).Insert()
	if err != nil {
		return entity.ResultadoDetalleCuantitativo{}, err
	}
	return resultadoDetalleCuantitativo, nil
}

// Get reads the resultadoDetalleCuantitativo with the specified ID from the database.
func (r repository) GetResultadoDetalleCuantitativoPorId(ctx context.Context, idResultadoDetalleCuantitativo int) (entity.ResultadoDetalleCuantitativo, error) {
	var resultadoDetalleCuantitativo entity.ResultadoDetalleCuantitativo
	err := r.db.With(ctx).Select().Model(idResultadoDetalleCuantitativo, &resultadoDetalleCuantitativo)
	return resultadoDetalleCuantitativo, err
}
