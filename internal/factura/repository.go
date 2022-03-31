package factura

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access facturas from the data source.
type Repository interface {
	// GetFacturaPorId returns the factura with the specified factura ID.
	GetFacturaPorId(ctx context.Context, idFactura int) (entity.Factura, error)
	// GetFacturas returns the list facturas.
	GetFacturas(ctx context.Context) ([]entity.Factura, error)
	GetFacturasConDatos(ctx context.Context) ([]FacturaConDatos, error)
	CrearFactura(ctx context.Context, factura entity.Factura) (entity.Factura, error)
	ActualizarFactura(ctx context.Context, factura entity.Factura) (entity.Factura, error)
}

// repository persists facturas in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new factura repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list facturas from the database.
func (r repository) GetFacturas(ctx context.Context) ([]entity.Factura, error) {
	var facturas []entity.Factura

	err := r.db.With(ctx).
		Select().
		From().
		All(&facturas)
	if err != nil {
		return facturas, err
	}
	return facturas, err
}

func (r repository) GetFacturasConDatos(ctx context.Context) ([]FacturaConDatos, error) {
	var facturas []entity.Factura
	var facturasConDatos []FacturaConDatos = []FacturaConDatos{}
	var clienteNombre, clienteApellido, vendedorNombre, vendedorApellido string

	err := r.db.With(ctx).
		Select().
		All(&facturas)

	for i := 0; i < len(facturas); i++ {
		idCliente := facturas[i].IdCliente
		err := r.db.With(ctx).
			Select("nombres", "apellidos").
			From("clientes").
			Where(dbx.HashExp{"id_cliente": idCliente}).
			Row(&clienteNombre, &clienteApellido)
		if err != nil {
			return []FacturaConDatos{}, err
		}

		idUsuario := facturas[i].IdUsuario
		err = r.db.With(ctx).
			Select("nombre", "apellido").
			From("usuarios").
			Where(dbx.HashExp{"id_usuario": idUsuario}).
			Row(&vendedorNombre, &vendedorApellido)
		if err != nil {
			return []FacturaConDatos{}, err
		}

		facturasConDatos = append(facturasConDatos, FacturaConDatos{
			facturas[i],
			clienteApellido + " " + clienteNombre,
			vendedorApellido + " " + vendedorNombre,
		})
	}

	if err != nil {
		return []FacturaConDatos{}, err
	}
	return facturasConDatos, err
}

// Create saves a new Factura record in the database.
// It returns the ID of the newly inserted factura record.
func (r repository) CrearFactura(ctx context.Context, factura entity.Factura) (entity.Factura, error) {
	err := r.db.With(ctx).Model(&factura).Insert()
	if err != nil {
		return entity.Factura{}, err
	}
	return factura, nil
}

// Create saves a new Factura record in the database.
// It returns the ID of the newly inserted factura record.
func (r repository) ActualizarFactura(ctx context.Context, factura entity.Factura) (entity.Factura, error) {
	var err error
	if factura.IdFactura != 0 {
		err = r.db.With(ctx).Model(&factura).Update()
	} else {
		err = r.db.With(ctx).Model(&factura).Insert()
	}
	if err != nil {
		return entity.Factura{}, err
	}
	return factura, nil
}

// Get reads the factura with the specified ID from the database.
func (r repository) GetFacturaPorId(ctx context.Context, idFactura int) (entity.Factura, error) {
	var factura entity.Factura
	err := r.db.With(ctx).Select().Model(idFactura, &factura)
	return factura, err
}
