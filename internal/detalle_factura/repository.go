package detalle_factura

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access detallesFactura from the data source.
type Repository interface {
	// GetDetalleFacturaPorId returns the detalleFactura with the specified detalleFactura ID.
	GetDetalleFacturaPorId(ctx context.Context, idDetalleFactura int) (entity.DetalleFactura, error)
	GetDetalleFacturaPorIdFactura(ctx context.Context, idFactura int) ([]DetallesFacturaConDatos, error)
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

func (r repository) GetDetalleFacturaPorIdFactura(ctx context.Context, idFactura int) ([]DetallesFacturaConDatos, error) {
	var detallesFactura []entity.DetalleFactura
	var detallesFacturaConDatos []DetallesFacturaConDatos = []DetallesFacturaConDatos{}
	var nombreProducto, nombreLote, nombreStock, nombreUnidad string

	err := r.db.With(ctx).
		Select().
		Where(dbx.HashExp{"id_factura": idFactura}).
		All(&detallesFactura)

	for i := 0; i < len(detallesFactura); i++ {
		//Limpiar
		nombreProducto = ""
		nombreLote = ""
		nombreStock = ""
		nombreUnidad = ""

		idReferencia := detallesFactura[i].IdReferencia

		if detallesFactura[i].Tabla == "lote" {
			err := r.db.With(ctx).
				Select("p.descripcion", "l.descripcion").
				From("lote as l").
				InnerJoin("proveedor_producto as pp", dbx.NewExp("pp.id_proveedor_producto = l.id_proveedor_producto")).
				InnerJoin("producto as p", dbx.NewExp("p.id_producto = pp.id_producto")).
				Where(dbx.HashExp{"l.id_lote": idReferencia}).
				Row(&nombreProducto, &nombreLote)
			if err != nil {
				return []DetallesFacturaConDatos{}, err
			}
		} else {
			err := r.db.With(ctx).
				Select("p.descripcion", "l.descripcion", "si.descripcion", "u.descripcion").
				From("stock_individual as si").
				InnerJoin("lote as l", dbx.NewExp("l.id_lote = si.id_lote")).
				InnerJoin("proveedor_producto as pp", dbx.NewExp("pp.id_proveedor_producto = l.id_proveedor_producto")).
				InnerJoin("producto as p", dbx.NewExp("p.id_producto = pp.id_producto")).
				InnerJoin("unidad as u", dbx.NewExp("u.id_unidad = p.id_unidad")).
				Where(dbx.HashExp{"si.id_stock_individual": idReferencia}).
				Row(&nombreProducto, &nombreLote, &nombreStock, &nombreUnidad)
			if err != nil {
				return []DetallesFacturaConDatos{}, err
			}
		}

		detallesFacturaConDatos = append(detallesFacturaConDatos, DetallesFacturaConDatos{
			DetalleFactura: detallesFactura[i],
			NombreProducto: nombreProducto,
			NombreUnidad:   nombreUnidad,
			NombreLote:     nombreLote,
			NombreStock:    nombreStock,
		})
	}

	if err != nil {
		return []DetallesFacturaConDatos{}, err
	}
	return detallesFacturaConDatos, err
}
