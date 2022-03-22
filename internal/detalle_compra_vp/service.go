package detallesCompraVP

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for detallesCompraVP.
type Service interface {
	GetDetallesCompraVP(ctx context.Context) ([]DetalleCompraVP, error)
	GetDetalleCompraVPPorId(ctx context.Context, idDetalleCompraVP int) (DetalleCompraVP, error)
	CrearDetalleCompraVP(ctx context.Context, input CreateDetalleCompraVPRequest) (DetalleCompraVP, error)
	ActualizarDetalleCompraVP(ctx context.Context, input UpdateDetalleCompraVPRequest) (DetalleCompraVP, error)
}

// DetallesCompraVP represents the data about an detallesCompraVP.
type DetalleCompraVP struct {
	entity.DetalleCompraVP
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new detallesCompraVP service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list detallesCompraVP.
func (s service) GetDetallesCompraVP(ctx context.Context) ([]DetalleCompraVP, error) {
	detallesCompraVP, err := s.repo.GetDetallesCompraVP(ctx)
	if err != nil {
		return nil, err
	}
	result := []DetalleCompraVP{}
	for _, item := range detallesCompraVP {
		result = append(result, DetalleCompraVP{item})
	}
	return result, nil
}

// CreateDetalleCompraVPRequest represents an detalleCompraVP creation request.
type CreateDetalleCompraVPRequest struct {
	IdCompra     int     `json:"id_compra"`
	IdProductoVp int     `json:"id_producto_vp"`
	Cantidad     int     `json:"cantidad"`
	Valor        float32 `json:"valor"`
}

type UpdateDetalleCompraVPRequest struct {
	IdDetalleCompraVP int     `json:"id_detalle_compra_vp"`
	IdCompra          int     `json:"id_compra"`
	IdProductoVp      int     `json:"id_producto_vp"`
	Cantidad          int     `json:"cantidad"`
	Valor             float32 `json:"valor"`
}

// Validate validates the UpdateDetalleCompraVPRequest fields.
func (m UpdateDetalleCompraVPRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdCompra, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.IdProductoVp, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Cantidad, validation.Required, validation.Length(0, 128)),
	)
}

// Validate validates the CreateDetalleCompraVPRequest fields.
func (m CreateDetalleCompraVPRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdCompra, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.IdProductoVp, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Cantidad, validation.Required, validation.Length(0, 128)),
	)
}

// CrearDetalleCompraVP creates a new detalleCompraVP.
func (s service) CrearDetalleCompraVP(ctx context.Context, req CreateDetalleCompraVPRequest) (DetalleCompraVP, error) {
	if err := req.Validate(); err != nil {
		return DetalleCompraVP{}, err
	}
	detalleCompraVPG, err := s.repo.CrearDetalleCompraVP(ctx, entity.DetalleCompraVP{
		IdCompra:     req.IdCompra,
		IdProductoVp: req.IdProductoVp,
		Cantidad:     req.Cantidad,
		Valor:        req.Valor,
	})
	if err != nil {
		return DetalleCompraVP{}, err
	}
	return DetalleCompraVP{detalleCompraVPG}, nil
}

// ActualizarDetalleCompraVP creates a new detalleCompraVP.
func (s service) ActualizarDetalleCompraVP(ctx context.Context, req UpdateDetalleCompraVPRequest) (DetalleCompraVP, error) {
	if err := req.ValidateUpdate(); err != nil {
		return DetalleCompraVP{}, err
	}
	detalleCompraVPG, err := s.repo.ActualizarDetalleCompraVP(ctx, entity.DetalleCompraVP{
		IdDetalleCompraVP: req.IdDetalleCompraVP,
		IdCompra:          req.IdCompra,
		IdProductoVp:      req.IdProductoVp,
		Cantidad:          req.Cantidad,
		Valor:             req.Valor,
	})
	if err != nil {
		return DetalleCompraVP{}, err
	}
	return DetalleCompraVP{detalleCompraVPG}, nil
}

// GetDetalleCompraVPPorId returns the detalleCompraVP with the specified the detalleCompraVP ID.
func (s service) GetDetalleCompraVPPorId(ctx context.Context, idDetalleCompraVP int) (DetalleCompraVP, error) {
	detalleCompraVP, err := s.repo.GetDetalleCompraVPPorId(ctx, idDetalleCompraVP)
	if err != nil {
		return DetalleCompraVP{}, err
	}
	return DetalleCompraVP{detalleCompraVP}, nil
}
