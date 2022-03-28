package factura

import (
	"context"
	"time"
	"veterinaria-server/internal/clientes"
	"veterinaria-server/internal/detalle_factura"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for facturas.
type Service interface {
	GetFacturas(ctx context.Context) ([]Factura, error)
	GetFacturaPorId(ctx context.Context, idFactura int) (Factura, error)
	CrearFactura(ctx context.Context, input CreateFacturaRequest) (Factura, error)
	ActualizarFactura(ctx context.Context, input UpdateFacturaRequest) (Factura, error)
}

// Facturas represents the data about an facturas.
type Factura struct {
	entity.Factura
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new facturas service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list facturas.
func (s service) GetFacturas(ctx context.Context) ([]Factura, error) {
	facturas, err := s.repo.GetFacturas(ctx)
	if err != nil {
		return nil, err
	}
	result := []Factura{}
	for _, item := range facturas {
		result = append(result, Factura{item})
	}
	return result, nil
}

// CreateFacturaRequest represents an factura creation request.
type CreateFacturaRequest struct {
	IdCliente int       `json:"id_cliente"`
	IdUsuario int       `json:"id_usuario"`
	Fecha     time.Time `json:"fecha"`
	Valor     float32   `json:"valor"`
}

type UpdateFacturaRequest struct {
	IdFactura int       `json:"id_factura"`
	IdCliente int       `json:"id_cliente"`
	IdUsuario int       `json:"id_usuario"`
	Fecha     time.Time `json:"fecha"`
	Valor     float32   `json:"valor"`
}

type CreateFacturaConDetallesRequest struct {
	Cliente         clientes.CreateClienteRequest                 `json:"cliente"`
	Factura         CreateFacturaRequest                          `json:"factura"`
	DetallesFactura []detalle_factura.CreateDetalleFacturaRequest `json:"detalles_factura"`
}

// Validate validates the UpdateFacturaRequest fields.
func (m UpdateFacturaRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdCliente, validation.Required),
		validation.Field(&m.IdUsuario, validation.Required),
	)
}

// Validate validates the CreateFacturaRequest fields.
func (m CreateFacturaRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdCliente, validation.Required),
		validation.Field(&m.IdUsuario, validation.Required),
	)
}

// CrearFactura creates a new factura.
func (s service) CrearFactura(ctx context.Context, req CreateFacturaRequest) (Factura, error) {
	if err := req.Validate(); err != nil {
		return Factura{}, err
	}
	facturaG, err := s.repo.CrearFactura(ctx, entity.Factura{
		IdCliente: req.IdCliente,
		IdUsuario: req.IdUsuario,
		Fecha:     req.Fecha,
		Valor:     req.Valor,
	})
	if err != nil {
		return Factura{}, err
	}
	return Factura{facturaG}, nil
}

// ActualizarFactura creates a new factura.
func (s service) ActualizarFactura(ctx context.Context, req UpdateFacturaRequest) (Factura, error) {
	if err := req.ValidateUpdate(); err != nil {
		return Factura{}, err
	}
	facturaG, err := s.repo.ActualizarFactura(ctx, entity.Factura{
		IdFactura: req.IdFactura,
		IdCliente: req.IdCliente,
		IdUsuario: req.IdUsuario,
		Fecha:     req.Fecha,
		Valor:     req.Valor,
	})
	if err != nil {
		return Factura{}, err
	}
	return Factura{facturaG}, nil
}

// GetFacturaPorId returns the factura with the specified the factura ID.
func (s service) GetFacturaPorId(ctx context.Context, idFactura int) (Factura, error) {
	factura, err := s.repo.GetFacturaPorId(ctx, idFactura)
	if err != nil {
		return Factura{}, err
	}
	return Factura{factura}, nil
}
