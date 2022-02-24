package resultado_examen_informativo

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access ResultadoDetalleInformativo from the data source.
type Repository interface {
	// GetResultadoDetalleInformativoPorId returns the resultadoDetalleInformativo with the specified resultadoDetalleInformativo ID.
	GetResultadoDetalleInformativoPorId(ctx context.Context, idResultadoDetalleInformativo int) (entity.ResultadoDetalleInformativo, error)
	CrearResultadoDetalleInformativo(ctx context.Context, resultadoDetalleInformativo entity.ResultadoDetalleInformativo) (entity.ResultadoDetalleInformativo, error)
}

// repository persists ResultadoDetalleInformativo in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new resultadoDetalleInformativo repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Create saves a new ResultadoDetalleInformativo record in the database.
// It returns the ID of the newly inserted resultadoDetalleInformativo record.
func (r repository) CrearResultadoDetalleInformativo(ctx context.Context, resultadoDetalleInformativo entity.ResultadoDetalleInformativo) (entity.ResultadoDetalleInformativo, error) {
	err := r.db.With(ctx).Model(&resultadoDetalleInformativo).Insert()
	if err != nil {
		return entity.ResultadoDetalleInformativo{}, err
	}
	return resultadoDetalleInformativo, nil
}

// Get reads the resultadoDetalleInformativo with the specified ID from the database.
func (r repository) GetResultadoDetalleInformativoPorId(ctx context.Context, idResultadoDetalleInformativo int) (entity.ResultadoDetalleInformativo, error) {
	var resultadoDetalleInformativo entity.ResultadoDetalleInformativo
	err := r.db.With(ctx).Select().Model(idResultadoDetalleInformativo, &resultadoDetalleInformativo)
	return resultadoDetalleInformativo, err
}
