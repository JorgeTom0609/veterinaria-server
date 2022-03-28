package detalle_factura

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for detallesFactura.
type Service interface {
	GetDetallesFactura(ctx context.Context) ([]DetalleFactura, error)
	GetDetalleFacturaPorId(ctx context.Context, idDetalleFactura int) (DetalleFactura, error)
	CrearDetalleFactura(ctx context.Context, input CreateDetalleFacturaRequest) (DetalleFactura, error)
	ActualizarDetalleFactura(ctx context.Context, input UpdateDetalleFacturaRequest) (DetalleFactura, error)
}

// DetallesFactura represents the data about an detallesFactura.
type DetalleFactura struct {
	entity.DetalleFactura
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new detallesFactura service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list detallesFactura.
func (s service) GetDetallesFactura(ctx context.Context) ([]DetalleFactura, error) {
	detallesFactura, err := s.repo.GetDetallesFactura(ctx)
	if err != nil {
		return nil, err
	}
	result := []DetalleFactura{}
	for _, item := range detallesFactura {
		result = append(result, DetalleFactura{item})
	}
	return result, nil
}

// CreateDetalleFacturaRequest represents an detalleFactura creation request.
type CreateDetalleFacturaRequest struct {
	IdFactura    int     `json:"id_factura"`
	IdProductoVp int     `json:"id_producto_vp"`
	Cantidad     int     `json:"cantidad"`
	Valor        float32 `json:"valor"`
}

type UpdateDetalleFacturaRequest struct {
	IdDetalleFactura int     `json:"id_detalle_factura"`
	IdFactura        int     `json:"id_factura"`
	IdProductoVp     int     `json:"id_producto_vp"`
	Cantidad         int     `json:"cantidad"`
	Valor            float32 `json:"valor"`
}

// Validate validates the UpdateDetalleFacturaRequest fields.
func (m UpdateDetalleFacturaRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdFactura, validation.Required),
		validation.Field(&m.IdProductoVp, validation.Required),
		validation.Field(&m.Cantidad, validation.Required),
	)
}

// Validate validates the CreateDetalleFacturaRequest fields.
func (m CreateDetalleFacturaRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdFactura, validation.Required),
		validation.Field(&m.IdProductoVp, validation.Required),
		validation.Field(&m.Cantidad, validation.Required),
	)
}

// CrearDetalleFactura creates a new detalleFactura.
func (s service) CrearDetalleFactura(ctx context.Context, req CreateDetalleFacturaRequest) (DetalleFactura, error) {
	if err := req.Validate(); err != nil {
		return DetalleFactura{}, err
	}
	detalleFacturaG, err := s.repo.CrearDetalleFactura(ctx, entity.DetalleFactura{
		IdFactura:    req.IdFactura,
		IdProductoVp: req.IdProductoVp,
		Cantidad:     req.Cantidad,
		Valor:        req.Valor,
	})
	if err != nil {
		return DetalleFactura{}, err
	}
	return DetalleFactura{detalleFacturaG}, nil
}

// ActualizarDetalleFactura creates a new detalleFactura.
func (s service) ActualizarDetalleFactura(ctx context.Context, req UpdateDetalleFacturaRequest) (DetalleFactura, error) {
	if err := req.ValidateUpdate(); err != nil {
		return DetalleFactura{}, err
	}
	detalleFacturaG, err := s.repo.ActualizarDetalleFactura(ctx, entity.DetalleFactura{
		IdDetalleFactura: req.IdDetalleFactura,
		IdFactura:        req.IdFactura,
		IdProductoVp:     req.IdProductoVp,
		Cantidad:         req.Cantidad,
		Valor:            req.Valor,
	})
	if err != nil {
		return DetalleFactura{}, err
	}
	return DetalleFactura{detalleFacturaG}, nil
}

// GetDetalleFacturaPorId returns the detalleFactura with the specified the detalleFactura ID.
func (s service) GetDetalleFacturaPorId(ctx context.Context, idDetalleFactura int) (DetalleFactura, error) {
	detalleFactura, err := s.repo.GetDetalleFacturaPorId(ctx, idDetalleFactura)
	if err != nil {
		return DetalleFactura{}, err
	}
	return DetalleFactura{detalleFactura}, nil
}
