package usuario_rol

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for usuarioRoles.
type Service interface {
	GetUsuarioRoles(ctx context.Context) ([]UsuarioRol, error)
	GetUsuarioRolPorId(ctx context.Context, idUsuarioRol int) (UsuarioRol, error)
	GetUsuarioRolPorCedula(ctx context.Context, cedula string) (UsuarioRol, error)
	CrearUsuarioRol(ctx context.Context, input CreateUsuarioRolRequest) (UsuarioRol, error)
	ActualizarUsuarioRol(ctx context.Context, input UpdateUsuarioRolRequest) (UsuarioRol, error)
}

// UsuarioRoles represents the data about an usuarioRoles.
type UsuarioRol struct {
	entity.UsuarioRol
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new usuarioRoles service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list usuarioRoles.
func (s service) GetUsuarioRoles(ctx context.Context) ([]UsuarioRol, error) {
	usuarioRoles, err := s.repo.GetUsuarioRoles(ctx)
	if err != nil {
		return nil, err
	}
	result := []UsuarioRol{}
	for _, item := range usuarioRoles {
		result = append(result, UsuarioRol{item})
	}
	return result, nil
}

// CreateUsuarioRolRequest represents an usuarioRol creation request.
type CreateUsuarioRolRequest struct {
	IdRol     int `json:"id_rol"`
	IdUsuario int `json:"id_usuario"`
}

type UpdateUsuarioRolRequest struct {
	IdUsuarioRol int `json:"id_usuario_rol"`
	IdRol        int `json:"id_rol"`
	IdUsuario    int `json:"id_usuario"`
}

// Validate validates the UpdateUsuarioRolRequest fields.
func (m UpdateUsuarioRolRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdRol, validation.Required),
		validation.Field(&m.IdUsuario, validation.Required),
	)
}

// Validate validates the CreateUsuarioRolRequest fields.
func (m CreateUsuarioRolRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdRol, validation.Required),
		validation.Field(&m.IdUsuario, validation.Required),
	)
}

// CrearUsuarioRol creates a new usuarioRol.
func (s service) CrearUsuarioRol(ctx context.Context, req CreateUsuarioRolRequest) (UsuarioRol, error) {
	if err := req.Validate(); err != nil {
		return UsuarioRol{}, err
	}
	usuarioRolG, err := s.repo.CrearUsuarioRol(ctx, entity.UsuarioRol{
		IdRol:     req.IdRol,
		IdUsuario: req.IdUsuario,
	})
	if err != nil {
		return UsuarioRol{}, err
	}
	return UsuarioRol{usuarioRolG}, nil
}

// ActualizarUsuarioRol creates a new usuarioRol.
func (s service) ActualizarUsuarioRol(ctx context.Context, req UpdateUsuarioRolRequest) (UsuarioRol, error) {
	if err := req.ValidateUpdate(); err != nil {
		return UsuarioRol{}, err
	}
	usuarioRolG, err := s.repo.ActualizarUsuarioRol(ctx, entity.UsuarioRol{
		IdUsuarioRol: req.IdUsuarioRol,
		IdRol:        req.IdRol,
		IdUsuario:    req.IdUsuario,
	})
	if err != nil {
		return UsuarioRol{}, err
	}
	return UsuarioRol{usuarioRolG}, nil
}

// GetUsuarioRolPorId returns the usuarioRol with the specified the usuarioRol ID.
func (s service) GetUsuarioRolPorId(ctx context.Context, idUsuarioRol int) (UsuarioRol, error) {
	usuarioRol, err := s.repo.GetUsuarioRolPorId(ctx, idUsuarioRol)
	if err != nil {
		return UsuarioRol{}, err
	}
	return UsuarioRol{usuarioRol}, nil
}

func (s service) GetUsuarioRolPorCedula(ctx context.Context, cedula string) (UsuarioRol, error) {
	usuarioRol, err := s.repo.GetUsuarioRolPorCedula(ctx, cedula)
	if err != nil {
		return UsuarioRol{}, err
	}
	return UsuarioRol{usuarioRol}, nil
}
