package consultas

import (
	"context"
	"time"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for consultas.
type Service interface {
	GetConsultas(ctx context.Context) ([]Consulta, error)
	GetConsultaPorId(ctx context.Context, idConsulta int) (Consulta, error)
	GetConsultaActiva(ctx context.Context, idUsuario int) (Consulta, error)
	GetConsultaPorMesYAnio(ctx context.Context, mes int, anio int) ([]ConsultaConDatos, error)
	GetConsultaPorMascota(ctx context.Context, idMascota int) ([]Consulta, error)
	GetConsultaRecetaServicios(ctx context.Context, idConsulta int) (RecetaServicios, error)
	CrearConsulta(ctx context.Context, input CreateConsultaRequest) (Consulta, error)
	ActualizarConsulta(ctx context.Context, input UpdateConsultaRequest) (Consulta, error)
}

// Consultas represents the data about an consultas.
type Consulta struct {
	entity.Consulta
}

type RecetaServicios struct {
	Receta    []RecetaConDatos   `json:"receta"`
	Servicios []ServicioConDatos `json:"servicios"`
}

type RecetaConDatos struct {
	Producto     string `json:"producto"`
	Prescripcion string `json:"prescripcion"`
}

type ServicioConDatos struct {
	Servicio string  `json:"servicio"`
	Valor    float32 `json:"valor"`
}

type ConsultaConDatos struct {
	entity.Consulta
	Mascota string `json:"mascota"`
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new consultas service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list consultas.
func (s service) GetConsultas(ctx context.Context) ([]Consulta, error) {
	consultas, err := s.repo.GetConsultas(ctx)
	if err != nil {
		return nil, err
	}
	result := []Consulta{}
	for _, item := range consultas {
		result = append(result, Consulta{item})
	}
	return result, nil
}

// CreateConsultaRequest represents an consulta creation request.
type CreateConsultaRequest struct {
	IdMascota              int       `json:"id_mascota"`
	IdUsuario              int       `json:"id_usuario"`
	Fecha                  time.Time `json:"fecha"`
	Valor                  float32   `json:"valor"`
	Motivo                 *string   `json:"motivo"`
	Temperatura            *float32  `json:"temperatura"`
	Peso                   *float32  `json:"peso"`
	Tamaño                 *float32  `json:"tamanio"`
	CondicionCorporal      *string   `json:"condicion_corporal"`
	NivelesDeshidratacion  *string   `json:"niveles_deshidratacion"`
	Diagnostico            *string   `json:"diagnostico"`
	Edad                   *string   `json:"edad"`
	TiempoLlenadoCapilar   int       `json:"tiempo_llenado_capilar"`
	FrecuenciaCardiaca     int       `json:"frecuencia_cardiaca"`
	FrecuenciaRespiratoria int       `json:"frecuencia_respiratoria"`
	EstadoConsulta         string    `json:"estado_consulta"`
}

type UpdateConsultaRequest struct {
	IdConsulta             int       `json:"id_consulta"`
	IdMascota              int       `json:"id_mascota"`
	IdUsuario              int       `json:"id_usuario"`
	Fecha                  time.Time `json:"fecha"`
	Valor                  float32   `json:"valor"`
	Motivo                 *string   `json:"motivo"`
	Temperatura            *float32  `json:"temperatura"`
	Peso                   *float32  `json:"peso"`
	Tamaño                 *float32  `json:"tamanio"`
	CondicionCorporal      *string   `json:"condicion_corporal"`
	NivelesDeshidratacion  *string   `json:"niveles_deshidratacion"`
	Diagnostico            *string   `json:"diagnostico"`
	Edad                   *string   `json:"edad"`
	TiempoLlenadoCapilar   int       `json:"tiempo_llenado_capilar"`
	FrecuenciaCardiaca     int       `json:"frecuencia_cardiaca"`
	FrecuenciaRespiratoria int       `json:"frecuencia_respiratoria"`
	EstadoConsulta         string    `json:"estado_consulta"`
}

// Validate validates the UpdateConsultaRequest fields.
func (m UpdateConsultaRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdMascota, validation.Required),
		validation.Field(&m.IdUsuario, validation.Required),
	)
}

// Validate validates the CreateConsultaRequest fields.
func (m CreateConsultaRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdMascota, validation.Required),
		validation.Field(&m.IdUsuario, validation.Required),
	)
}

// CrearConsulta creates a new consulta.
func (s service) CrearConsulta(ctx context.Context, req CreateConsultaRequest) (Consulta, error) {
	if err := req.Validate(); err != nil {
		return Consulta{}, err
	}
	consultaG, err := s.repo.CrearConsulta(ctx, entity.Consulta{
		IdMascota:              req.IdMascota,
		IdUsuario:              req.IdUsuario,
		Fecha:                  req.Fecha,
		Valor:                  req.Valor,
		Motivo:                 req.Motivo,
		Temperatura:            req.Temperatura,
		Peso:                   req.Peso,
		Tamaño:                 req.Tamaño,
		CondicionCorporal:      req.CondicionCorporal,
		NivelesDeshidratacion:  req.NivelesDeshidratacion,
		Diagnostico:            req.Diagnostico,
		TiempoLlenadoCapilar:   req.TiempoLlenadoCapilar,
		FrecuenciaCardiaca:     req.FrecuenciaCardiaca,
		FrecuenciaRespiratoria: req.FrecuenciaRespiratoria,
		EstadoConsulta:         req.EstadoConsulta,
		Edad:                   req.Edad,
	})
	if err != nil {
		return Consulta{}, err
	}
	return Consulta{consultaG}, nil
}

// ActualizarConsulta creates a new consulta.
func (s service) ActualizarConsulta(ctx context.Context, req UpdateConsultaRequest) (Consulta, error) {
	if err := req.ValidateUpdate(); err != nil {
		return Consulta{}, err
	}
	consultaG, err := s.repo.ActualizarConsulta(ctx, entity.Consulta{
		IdConsulta:             req.IdConsulta,
		IdMascota:              req.IdMascota,
		IdUsuario:              req.IdUsuario,
		Fecha:                  req.Fecha,
		Valor:                  req.Valor,
		Motivo:                 req.Motivo,
		Temperatura:            req.Temperatura,
		Peso:                   req.Peso,
		Tamaño:                 req.Tamaño,
		CondicionCorporal:      req.CondicionCorporal,
		NivelesDeshidratacion:  req.NivelesDeshidratacion,
		Diagnostico:            req.Diagnostico,
		TiempoLlenadoCapilar:   req.TiempoLlenadoCapilar,
		FrecuenciaCardiaca:     req.FrecuenciaCardiaca,
		FrecuenciaRespiratoria: req.FrecuenciaRespiratoria,
		EstadoConsulta:         req.EstadoConsulta,
		Edad:                   req.Edad,
	})
	if err != nil {
		return Consulta{}, err
	}
	return Consulta{consultaG}, nil
}

// GetConsultaPorId returns the consulta with the specified the consulta ID.
func (s service) GetConsultaPorId(ctx context.Context, idConsulta int) (Consulta, error) {
	consulta, err := s.repo.GetConsultaPorId(ctx, idConsulta)
	if err != nil {
		return Consulta{}, err
	}
	return Consulta{consulta}, nil
}

func (s service) GetConsultaActiva(ctx context.Context, idUsuario int) (Consulta, error) {
	consulta, err := s.repo.GetConsultaActiva(ctx, idUsuario)
	if err != nil {
		return Consulta{}, err
	}
	return Consulta{consulta}, nil
}

func (s service) GetConsultaPorMesYAnio(ctx context.Context, mes int, anio int) ([]ConsultaConDatos, error) {
	consultas, err := s.repo.GetConsultaPorMesYAnio(ctx, mes, anio)
	if err != nil {
		return nil, err
	}
	return consultas, nil
}

func (s service) GetConsultaPorMascota(ctx context.Context, idMascota int) ([]Consulta, error) {
	consultas, err := s.repo.GetConsultaPorMascota(ctx, idMascota)
	if err != nil {
		return nil, err
	}
	result := []Consulta{}
	for _, item := range consultas {
		result = append(result, Consulta{item})
	}
	return result, nil
}

func (s service) GetConsultaRecetaServicios(ctx context.Context, idConsulta int) (RecetaServicios, error) {
	recetaServicios, err := s.repo.GetConsultaRecetaServicios(ctx, idConsulta)
	if err != nil {
		return RecetaServicios{}, err
	}
	return recetaServicios, nil
}
