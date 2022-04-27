package detalle_compra

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for detallesCompra.
type Service interface {
	GetDetallesCompra(ctx context.Context) ([]DetalleCompra, error)
	GetDetalleCompraPorId(ctx context.Context, idDetalleCompra int) (DetalleCompra, error)
	GetDetalleCompraPorIdCompra(ctx context.Context, idCompra int) ([]DetallesCompraConDatos, error)
	CrearDetalleCompra(ctx context.Context, input CreateDetalleCompraRequest) (DetalleCompra, error)
	ActualizarDetalleCompra(ctx context.Context, input UpdateDetalleCompraRequest) (DetalleCompra, error)
}

// DetallesCompra represents the data about an detallesCompra.
type DetalleCompra struct {
	entity.DetalleCompra
}

type DetallesCompraConDatos struct {
	entity.DetalleCompra
	NombreProducto string `json:"nombreProducto"`
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new detallesCompra service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list detallesCompra.
func (s service) GetDetallesCompra(ctx context.Context) ([]DetalleCompra, error) {
	detallesCompra, err := s.repo.GetDetallesCompra(ctx)
	if err != nil {
		return nil, err
	}
	result := []DetalleCompra{}
	for _, item := range detallesCompra {
		result = append(result, DetalleCompra{item})
	}
	return result, nil
}

// CreateDetalleCompraRequest represents an detalleCompra creation request.
type CreateDetalleCompraRequest struct {
	IdCompra int     `json:"id_compra"`
	IdLote   int     `json:"id_lote"`
	Cantidad int     `json:"cantidad"`
	Valor    float32 `json:"valor"`
}

type UpdateDetalleCompraRequest struct {
	IdDetalleCompra int     `json:"id_detalle_compra"`
	IdCompra        int     `json:"id_compra"`
	IdLote          int     `json:"id_lote"`
	Cantidad        int     `json:"cantidad"`
	Valor           float32 `json:"valor"`
}

// Validate validates the UpdateDetalleCompraRequest fields.
func (m UpdateDetalleCompraRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdCompra, validation.Required),
		validation.Field(&m.IdLote, validation.Required),
		validation.Field(&m.Cantidad, validation.Required),
	)
}

// Validate validates the CreateDetalleCompraRequest fields.
func (m CreateDetalleCompraRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdCompra, validation.Required),
		validation.Field(&m.IdLote, validation.Required),
		validation.Field(&m.Cantidad, validation.Required),
	)
}

// CrearDetalleCompra creates a new detalleCompra.
func (s service) CrearDetalleCompra(ctx context.Context, req CreateDetalleCompraRequest) (DetalleCompra, error) {
	if err := req.Validate(); err != nil {
		return DetalleCompra{}, err
	}
	detalleCompraG, err := s.repo.CrearDetalleCompra(ctx, entity.DetalleCompra{
		IdCompra: req.IdCompra,
		IdLote:   req.IdLote,
		Cantidad: req.Cantidad,
		Valor:    req.Valor,
	})
	if err != nil {
		return DetalleCompra{}, err
	}
	return DetalleCompra{detalleCompraG}, nil
}

// ActualizarDetalleCompra creates a new detalleCompra.
func (s service) ActualizarDetalleCompra(ctx context.Context, req UpdateDetalleCompraRequest) (DetalleCompra, error) {
	if err := req.ValidateUpdate(); err != nil {
		return DetalleCompra{}, err
	}
	detalleCompraG, err := s.repo.ActualizarDetalleCompra(ctx, entity.DetalleCompra{
		IdDetalleCompra: req.IdDetalleCompra,
		IdCompra:        req.IdCompra,
		IdLote:          req.IdLote,
		Cantidad:        req.Cantidad,
		Valor:           req.Valor,
	})
	if err != nil {
		return DetalleCompra{}, err
	}
	return DetalleCompra{detalleCompraG}, nil
}

// GetDetalleCompraPorId returns the detalleCompra with the specified the detalleCompra ID.
func (s service) GetDetalleCompraPorId(ctx context.Context, idDetalleCompra int) (DetalleCompra, error) {
	detalleCompra, err := s.repo.GetDetalleCompraPorId(ctx, idDetalleCompra)
	if err != nil {
		return DetalleCompra{}, err
	}
	return DetalleCompra{detalleCompra}, nil
}

func (s service) GetDetalleCompraPorIdCompra(ctx context.Context, idCompra int) ([]DetallesCompraConDatos, error) {
	detalleCompra, err := s.repo.GetDetalleCompraPorIdCompra(ctx, idCompra)
	if err != nil {
		return nil, err
	}
	return detalleCompra, nil
}
