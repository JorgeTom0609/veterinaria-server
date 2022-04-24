package stock_individual

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access stocksIndividual from the data source.
type Repository interface {
	// GetStockIndividualPorId returns the stockIndividual with the specified stockIndividual ID.
	GetStockIndividualPorId(ctx context.Context, idStockIndividual int) (entity.StockIndividual, error)
	// GetStocksIndividual returns the list stocksIndividual.
	GetStocksIndividual(ctx context.Context) ([]entity.StockIndividual, error)
	CrearStockIndividual(ctx context.Context, stockIndividual entity.StockIndividual) (entity.StockIndividual, error)
	ActualizarStockIndividual(ctx context.Context, stockIndividual entity.StockIndividual) (entity.StockIndividual, error)
}

// repository persists stocksIndividual in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new stockIndividual repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list stocksIndividual from the database.
func (r repository) GetStocksIndividual(ctx context.Context) ([]entity.StockIndividual, error) {
	var stocksIndividual []entity.StockIndividual

	err := r.db.With(ctx).
		Select().
		All(&stocksIndividual)
	if err != nil {
		return stocksIndividual, err
	}
	return stocksIndividual, err
}

// Create saves a new StockIndividual record in the database.
// It returns the ID of the newly inserted stockIndividual record.
func (r repository) CrearStockIndividual(ctx context.Context, stockIndividual entity.StockIndividual) (entity.StockIndividual, error) {
	err := r.db.With(ctx).Model(&stockIndividual).Insert()
	if err != nil {
		return entity.StockIndividual{}, err
	}
	return stockIndividual, nil
}

// Create saves a new StockIndividual record in the database.
// It returns the ID of the newly inserted stockIndividual record.
func (r repository) ActualizarStockIndividual(ctx context.Context, stockIndividual entity.StockIndividual) (entity.StockIndividual, error) {
	var err error
	if stockIndividual.IdStockIndividual != 0 {
		err = r.db.With(ctx).Model(&stockIndividual).Update()
	} else {
		err = r.db.With(ctx).Model(&stockIndividual).Insert()
	}
	if err != nil {
		return entity.StockIndividual{}, err
	}
	return stockIndividual, nil
}

// Get reads the stockIndividual with the specified ID from the database.
func (r repository) GetStockIndividualPorId(ctx context.Context, idStockIndividual int) (entity.StockIndividual, error) {
	var stockIndividual entity.StockIndividual
	err := r.db.With(ctx).Select().Model(idStockIndividual, &stockIndividual)
	return stockIndividual, err
}
