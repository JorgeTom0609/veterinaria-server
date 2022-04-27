package stock_individual

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for stocksIndividual.
type Service interface {
	GetStocksIndividual(ctx context.Context) ([]StockIndividual, error)
	GetStockIndividualPorId(ctx context.Context, idStockIndividual int) (StockIndividual, error)
	CrearStockIndividual(ctx context.Context, input CreateStockIndividualRequest) (StockIndividual, error)
	ActualizarStockIndividual(ctx context.Context, input UpdateStockIndividualRequest) (StockIndividual, error)
}

// StocksIndividual represents the data about an stocksIndividual.
type StockIndividual struct {
	entity.StockIndividual
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new stocksIndividual service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list stocksIndividual.
func (s service) GetStocksIndividual(ctx context.Context) ([]StockIndividual, error) {
	stocksIndividual, err := s.repo.GetStocksIndividual(ctx)
	if err != nil {
		return nil, err
	}
	result := []StockIndividual{}
	for _, item := range stocksIndividual {
		result = append(result, StockIndividual{item})
	}
	return result, nil
}

// CreateStockIndividualRequest represents an stockIndividual creation request.
type CreateStockIndividualRequest struct {
	IdLote          int     `json:"id_lote"`
	Descripcion     string  `json:"descripcion"`
	CantidadInicial float32 `json:"cantidad_inicial"`
	Cantidad        float32 `json:"cantidad"`
}

type UpdateStockIndividualRequest struct {
	IdStockIndividual int     `json:"id_stock_individual"`
	IdLote            int     `json:"id_lote"`
	Descripcion       string  `json:"descripcion"`
	CantidadInicial   float32 `json:"cantidad_inicial"`
	Cantidad          float32 `json:"cantidad"`
}

// Validate validates the UpdateStockIndividualRequest fields.
func (m UpdateStockIndividualRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdLote, validation.Required),
		validation.Field(&m.CantidadInicial, validation.Required),
		validation.Field(&m.Cantidad, validation.Required),
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 128)),
	)
}

// Validate validates the CreateStockIndividualRequest fields.
func (m CreateStockIndividualRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdLote, validation.Required),
		validation.Field(&m.CantidadInicial, validation.Required),
		validation.Field(&m.Cantidad, validation.Required),
		validation.Field(&m.Descripcion, validation.Required, validation.Length(0, 128)),
	)
}

// CrearStockIndividual creates a new stockIndividual.
func (s service) CrearStockIndividual(ctx context.Context, req CreateStockIndividualRequest) (StockIndividual, error) {
	if err := req.Validate(); err != nil {
		return StockIndividual{}, err
	}
	stockIndividualG, err := s.repo.CrearStockIndividual(ctx, entity.StockIndividual{
		IdLote:          req.IdLote,
		CantidadInicial: req.CantidadInicial,
		Descripcion:     req.Descripcion,
		Cantidad:        req.Cantidad,
	})
	if err != nil {
		return StockIndividual{}, err
	}
	return StockIndividual{stockIndividualG}, nil
}

// ActualizarStockIndividual creates a new stockIndividual.
func (s service) ActualizarStockIndividual(ctx context.Context, req UpdateStockIndividualRequest) (StockIndividual, error) {
	if err := req.ValidateUpdate(); err != nil {
		return StockIndividual{}, err
	}
	stockIndividualG, err := s.repo.ActualizarStockIndividual(ctx, entity.StockIndividual{
		IdStockIndividual: req.IdStockIndividual,
		IdLote:            req.IdLote,
		CantidadInicial:   req.CantidadInicial,
		Descripcion:       req.Descripcion,
		Cantidad:          req.Cantidad,
	})
	if err != nil {
		return StockIndividual{}, err
	}
	return StockIndividual{stockIndividualG}, nil
}

// GetStockIndividualPorId returns the stockIndividual with the specified the stockIndividual ID.
func (s service) GetStockIndividualPorId(ctx context.Context, idStockIndividual int) (StockIndividual, error) {
	stockIndividual, err := s.repo.GetStockIndividualPorId(ctx, idStockIndividual)
	if err != nil {
		return StockIndividual{}, err
	}
	return StockIndividual{stockIndividual}, nil
}
