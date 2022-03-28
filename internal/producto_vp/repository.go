package producto_vp

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access productosVP from the data source.
type Repository interface {
	// GetProductoVPPorId returns the productoVP with the specified productoVP ID.
	GetProductoVPPorId(ctx context.Context, idProductoVP int) (entity.ProductoVP, error)
	// GetProductosVP returns the list productosVP.
	GetProductosVP(ctx context.Context) ([]entity.ProductoVP, error)
	GetProductosVPConStock(ctx context.Context) ([]entity.ProductoVP, error)
	GetProductosVPPocoStock(ctx context.Context) ([]entity.ProductoVP, error)
	CrearProductoVP(ctx context.Context, productoVP entity.ProductoVP) (entity.ProductoVP, error)
	ActualizarProductoVP(ctx context.Context, productoVP entity.ProductoVP) (entity.ProductoVP, error)
}

// repository persists productosVP in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new productoVP repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list productosVP from the database.
func (r repository) GetProductosVP(ctx context.Context) ([]entity.ProductoVP, error) {
	var productosVP []entity.ProductoVP

	err := r.db.With(ctx).
		Select().
		From().
		All(&productosVP)
	if err != nil {
		return productosVP, err
	}
	return productosVP, err
}

func (r repository) GetProductosVPConStock(ctx context.Context) ([]entity.ProductoVP, error) {
	var productosVP []entity.ProductoVP

	err := r.db.With(ctx).
		Select().
		From().
		//Where(dbx.NewExp("id={:id}", dbx.Params{"id": 100})).
		Where(dbx.NewExp("stock>0")).
		All(&productosVP)
	if err != nil {
		return productosVP, err
	}
	return productosVP, err
}

func (r repository) GetProductosVPPocoStock(ctx context.Context) ([]entity.ProductoVP, error) {
	var productosVP []entity.ProductoVP

	err := r.db.With(ctx).
		Select().
		From().
		Where(dbx.NewExp("stock<=stock_minimo")).
		All(&productosVP)
	if err != nil {
		return productosVP, err
	}
	return productosVP, err
}

// Create saves a new ProductoVP record in the database.
// It returns the ID of the newly inserted productoVP record.
func (r repository) CrearProductoVP(ctx context.Context, productoVP entity.ProductoVP) (entity.ProductoVP, error) {
	err := r.db.With(ctx).Model(&productoVP).Insert()
	if err != nil {
		return entity.ProductoVP{}, err
	}
	return productoVP, nil
}

// Create saves a new ProductoVP record in the database.
// It returns the ID of the newly inserted productoVP record.
func (r repository) ActualizarProductoVP(ctx context.Context, productoVP entity.ProductoVP) (entity.ProductoVP, error) {
	var err error
	if productoVP.IdProductoVP != 0 {
		err = r.db.With(ctx).Model(&productoVP).Update()
	} else {
		err = r.db.With(ctx).Model(&productoVP).Insert()
	}
	if err != nil {
		return entity.ProductoVP{}, err
	}
	return productoVP, nil
}

// Get reads the productoVP with the specified ID from the database.
func (r repository) GetProductoVPPorId(ctx context.Context, idProductoVP int) (entity.ProductoVP, error) {
	var productoVP entity.ProductoVP
	err := r.db.With(ctx).Select().Model(idProductoVP, &productoVP)
	return productoVP, err
}
