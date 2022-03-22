package detalle_factura

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access detallesFactura from the data source.
type Repository interface {
	// GetDetalleFacturaPorId returns the detalleFactura with the specified detalleFactura ID.
	GetDetalleFacturaPorId(ctx context.Context, idDetalleFactura int) (entity.DetalleFactura, error)
	// GetDetallesFactura returns the list detallesFactura.
	GetDetallesFactura(ctx context.Context) ([]entity.DetalleFactura, error)
	CrearDetalleFactura(ctx context.Context, detalleFactura entity.DetalleFactura) (entity.DetalleFactura, error)
	ActualizarDetalleFactura(ctx context.Context, detalleFactura entity.DetalleFactura) (entity.DetalleFactura, error)
}

// repository persists detallesFactura in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new detalleFactura repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list detallesFactura from the database.
func (r repository) GetDetallesFactura(ctx context.Context) ([]entity.DetalleFactura, error) {
	var detallesFactura []entity.DetalleFactura

	err := r.db.With(ctx).
		Select().
		From().
		All(&detallesFactura)
	if err != nil {
		return detallesFactura, err
	}
	return detallesFactura, err
}

// Create saves a new DetalleFactura record in the database.
// It returns the ID of the newly inserted detalleFactura record.
func (r repository) CrearDetalleFactura(ctx context.Context, detalleFactura entity.DetalleFactura) (entity.DetalleFactura, error) {
	err := r.db.With(ctx).Model(&detalleFactura).Insert()
	if err != nil {
		return entity.DetalleFactura{}, err
	}
	return detalleFactura, nil
}

// Create saves a new DetalleFactura record in the database.
// It returns the ID of the newly inserted detalleFactura record.
func (r repository) ActualizarDetalleFactura(ctx context.Context, detalleFactura entity.DetalleFactura) (entity.DetalleFactura, error) {
	var err error
	if detalleFactura.IdDetalleFactura != 0 {
		err = r.db.With(ctx).Model(&detalleFactura).Update()
	} else {
		err = r.db.With(ctx).Model(&detalleFactura).Insert()
	}
	if err != nil {
		return entity.DetalleFactura{}, err
	}
	return detalleFactura, nil
}

// Get reads the detalleFactura with the specified ID from the database.
func (r repository) GetDetalleFacturaPorId(ctx context.Context, idDetalleFactura int) (entity.DetalleFactura, error) {
	var detalleFactura entity.DetalleFactura
	err := r.db.With(ctx).Select().Model(idDetalleFactura, &detalleFactura)
	return detalleFactura, err
}
