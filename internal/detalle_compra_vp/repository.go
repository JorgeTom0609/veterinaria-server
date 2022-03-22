package detalle_compra_vp

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"
)

// Repository encapsulates the logic to access detallesCompraVP from the data source.
type Repository interface {
	// GetDetalleCompraVPPorId returns the detalleCompraVP with the specified detalleCompraVP ID.
	GetDetalleCompraVPPorId(ctx context.Context, idDetalleCompraVP int) (entity.DetalleCompraVP, error)
	// GetDetallesCompraVP returns the list detallesCompraVP.
	GetDetallesCompraVP(ctx context.Context) ([]entity.DetalleCompraVP, error)
	CrearDetalleCompraVP(ctx context.Context, detalleCompraVP entity.DetalleCompraVP) (entity.DetalleCompraVP, error)
	ActualizarDetalleCompraVP(ctx context.Context, detalleCompraVP entity.DetalleCompraVP) (entity.DetalleCompraVP, error)
}

// repository persists detallesCompraVP in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new detalleCompraVP repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list detallesCompraVP from the database.
func (r repository) GetDetallesCompraVP(ctx context.Context) ([]entity.DetalleCompraVP, error) {
	var detallesCompraVP []entity.DetalleCompraVP

	err := r.db.With(ctx).
		Select().
		From().
		All(&detallesCompraVP)
	if err != nil {
		return detallesCompraVP, err
	}
	return detallesCompraVP, err
}

// Create saves a new DetalleCompraVP record in the database.
// It returns the ID of the newly inserted detalleCompraVP record.
func (r repository) CrearDetalleCompraVP(ctx context.Context, detalleCompraVP entity.DetalleCompraVP) (entity.DetalleCompraVP, error) {
	err := r.db.With(ctx).Model(&detalleCompraVP).Insert()
	if err != nil {
		return entity.DetalleCompraVP{}, err
	}
	return detalleCompraVP, nil
}

// Create saves a new DetalleCompraVP record in the database.
// It returns the ID of the newly inserted detalleCompraVP record.
func (r repository) ActualizarDetalleCompraVP(ctx context.Context, detalleCompraVP entity.DetalleCompraVP) (entity.DetalleCompraVP, error) {
	var err error
	if detalleCompraVP.IdDetalleCompraVP != 0 {
		err = r.db.With(ctx).Model(&detalleCompraVP).Update()
	} else {
		err = r.db.With(ctx).Model(&detalleCompraVP).Insert()
	}
	if err != nil {
		return entity.DetalleCompraVP{}, err
	}
	return detalleCompraVP, nil
}

// Get reads the detalleCompraVP with the specified ID from the database.
func (r repository) GetDetalleCompraVPPorId(ctx context.Context, idDetalleCompraVP int) (entity.DetalleCompraVP, error) {
	var detalleCompraVP entity.DetalleCompraVP
	err := r.db.With(ctx).Select().Model(idDetalleCompraVP, &detalleCompraVP)
	return detalleCompraVP, err
}
