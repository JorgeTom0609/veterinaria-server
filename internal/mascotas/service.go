package mascotas

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for mascotas.
type Service interface {
	GetMascotasPorCliente(ctx context.Context, idCliente int) ([]Mascota, error)
	GetMascotaPorId(ctx context.Context, idMascota int) (Mascota, error)
	CrearMascota(ctx context.Context, input CreateMascotaRequest) (Mascota, error)
	ActualizarMascotaPorGrupo(ctx context.Context, input CreateMascotaPorGrupoRequest) ([]Mascota, error)
}

// Mascota represents the data about an mascotas.
type Mascota struct {
	entity.Mascota
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new mascotas service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// GetMascotasPorCliente returns the list mascotas by cliente.
func (s service) GetMascotasPorCliente(ctx context.Context, idCliente int) ([]Mascota, error) {
	mascotas, err := s.repo.GetMascotasPorCliente(ctx, idCliente)
	if err != nil {
		return nil, err
	}
	result := []Mascota{}
	for _, item := range mascotas {
		result = append(result, Mascota{item})
	}
	return result, nil
}

// CreateMascotaRequest represents an mascota creation request.
type CreateMascotaRequest struct {
	IdEspecie int     `json:"id_especie"`
	IdCliente int     `json:"id_cliente"`
	IdGenero  int     `json:"id_genero"`
	Nombre    *string `json:"nombre"`
	Raza      *string `json:"raza"`
	Color     *string `json:"color"`
}

type UpdateMascotaRequest struct {
	IdMascota int     `json:"id_mascota"`
	IdEspecie int     `json:"id_especie"`
	IdCliente int     `json:"id_cliente"`
	IdGenero  int     `json:"id_genero"`
	Nombre    *string `json:"nombre"`
	Raza      *string `json:"raza"`
	Color     *string `json:"color"`
}

type CreateMascotaPorGrupoRequest struct {
	Mascotas []UpdateMascotaRequest `json:"mascotas"`
}

// Validate validates the CreateMascotaRequest fields.
func (m UpdateMascotaRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdEspecie, validation.Required),
		validation.Field(&m.IdCliente, validation.Required),
		validation.Field(&m.IdGenero, validation.Required),
	)
}

// Validate validates the CreateMascotaRequest fields.
func (m CreateMascotaRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdEspecie, validation.Required),
		validation.Field(&m.IdCliente, validation.Required),
		validation.Field(&m.IdGenero, validation.Required),
	)
}

// CrearMascota creates a new mascota.
func (s service) CrearMascota(ctx context.Context, req CreateMascotaRequest) (Mascota, error) {
	if err := req.Validate(); err != nil {
		return Mascota{}, err
	}

	mascotaG, err := s.repo.CrearMascota(ctx, entity.Mascota{
		IdCliente: req.IdCliente,
		IdEspecie: req.IdEspecie,
		IdGenero:  req.IdGenero,
		Nombre:    req.Nombre,
		Raza:      req.Raza,
		Color:     req.Color,
	})
	if err != nil {
		return Mascota{}, err
	}
	return Mascota{mascotaG}, nil
}

// ActualizarMascotaPorGrupo creates a new mascota.
func (s service) ActualizarMascotaPorGrupo(ctx context.Context, req CreateMascotaPorGrupoRequest) ([]Mascota, error) {
	result := []Mascota{}
	for i := 0; i < len(req.Mascotas); i++ {
		if err := req.Mascotas[i].Validate(); err != nil {
			return result, err
		}
		mascotaG, err := s.repo.ActualizarMascotaPorGrupo(ctx, entity.Mascota{
			IdMascota: req.Mascotas[i].IdMascota,
			IdCliente: req.Mascotas[i].IdCliente,
			IdEspecie: req.Mascotas[i].IdEspecie,
			IdGenero:  req.Mascotas[i].IdGenero,
			Nombre:    req.Mascotas[i].Nombre,
			Raza:      req.Mascotas[i].Raza,
			Color:     req.Mascotas[i].Color,
		})
		if err != nil {
			return result, err
		}
		result = append(result, Mascota{mascotaG})
	}
	return result, nil
}

// GetMascotaPorId returns the mascota with the specified the mascota ID.
func (s service) GetMascotaPorId(ctx context.Context, idMascota int) (Mascota, error) {
	mascota, err := s.repo.GetMascotaPorId(ctx, idMascota)
	if err != nil {
		return Mascota{}, err
	}
	return Mascota{mascota}, nil
}
