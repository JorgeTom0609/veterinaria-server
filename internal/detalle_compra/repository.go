package detalle_compra

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access detallesCompra from the data source.
type Repository interface {
	// GetDetalleCompraPorId returns the detalleCompra with the specified detalleCompra ID.
	GetDetalleCompraPorId(ctx context.Context, idDetalleCompra int) (entity.DetalleCompra, error)
	GetDetalleCompraPorIdCompra(ctx context.Context, idCompra int) ([]DetallesCompraConDatos, error)
	// GetDetallesCompra returns the list detallesCompra.
	GetDetallesCompra(ctx context.Context) ([]entity.DetalleCompra, error)
	CrearDetalleCompra(ctx context.Context, detalleCompra entity.DetalleCompra) (entity.DetalleCompra, error)
	ActualizarDetalleCompra(ctx context.Context, detalleCompra entity.DetalleCompra) (entity.DetalleCompra, error)
}

// repository persists detallesCompra in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new detalleCompra repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list detallesCompra from the database.
func (r repository) GetDetallesCompra(ctx context.Context) ([]entity.DetalleCompra, error) {
	var detallesCompra []entity.DetalleCompra

	err := r.db.With(ctx).
		Select().
		From().
		All(&detallesCompra)
	if err != nil {
		return detallesCompra, err
	}
	return detallesCompra, err
}

// Create saves a new DetalleCompra record in the database.
// It returns the ID of the newly inserted detalleCompra record.
func (r repository) CrearDetalleCompra(ctx context.Context, detalleCompra entity.DetalleCompra) (entity.DetalleCompra, error) {
	err := r.db.With(ctx).Model(&detalleCompra).Insert()
	if err != nil {
		return entity.DetalleCompra{}, err
	}
	return detalleCompra, nil
}

// Create saves a new DetalleCompra record in the database.
// It returns the ID of the newly inserted detalleCompra record.
func (r repository) ActualizarDetalleCompra(ctx context.Context, detalleCompra entity.DetalleCompra) (entity.DetalleCompra, error) {
	var err error
	if detalleCompra.IdDetalleCompra != 0 {
		err = r.db.With(ctx).Model(&detalleCompra).Update()
	} else {
		err = r.db.With(ctx).Model(&detalleCompra).Insert()
	}
	if err != nil {
		return entity.DetalleCompra{}, err
	}
	return detalleCompra, nil
}

// Get reads the detalleCompra with the specified ID from the database.
func (r repository) GetDetalleCompraPorId(ctx context.Context, idDetalleCompra int) (entity.DetalleCompra, error) {
	var detalleCompra entity.DetalleCompra
	err := r.db.With(ctx).Select().Model(idDetalleCompra, &detalleCompra)
	return detalleCompra, err
}

func (r repository) GetDetalleCompraPorIdCompra(ctx context.Context, idCompra int) ([]DetallesCompraConDatos, error) {
	var detallesCompra []entity.DetalleCompra
	var detallesCompraConDatos []DetallesCompraConDatos = []DetallesCompraConDatos{}
	var nombreProducto string

	err := r.db.With(ctx).
		Select().
		Where(dbx.HashExp{"id_compra": idCompra}).
		All(&detallesCompra)

	for i := 0; i < len(detallesCompra); i++ {
		idLote := detallesCompra[i].IdLote
		err := r.db.With(ctx).
			Select("p.descripcion").
			From("lote as l").
			InnerJoin("proveedor_producto as pp", dbx.NewExp("pp.id_proveedor_producto = l.id_proveedor_producto")).
			InnerJoin("producto as p", dbx.NewExp("p.id_producto = pp.id_producto")).
			Where(dbx.HashExp{"l.id_lote": idLote}).
			Row(&nombreProducto)
		if err != nil {
			return []DetallesCompraConDatos{}, err
		}

		detallesCompraConDatos = append(detallesCompraConDatos, DetallesCompraConDatos{
			detallesCompra[i],
			nombreProducto,
		})
	}

	if err != nil {
		return []DetallesCompraConDatos{}, err
	}
	return detallesCompraConDatos, err
}
