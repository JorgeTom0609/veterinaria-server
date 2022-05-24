package examen_mascota

import (
	"context"
	"database/sql"
	"time"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for examenesMascota.
type Service interface {
	GetExamenesMascota(ctx context.Context) ([]ExamenMascota, error)
	GetExamenesMascotaPorMascotayEstado(ctx context.Context, idExamenMascota int, estado string) ([]ExamenMascotaAll, error)
	GetExamenesMascotaPorEstado(ctx context.Context, estado string) ([]ExamenMascotaAll, error)
	GetExamenMascotaPorId(ctx context.Context, idExamenMascota int) (ExamenMascota, error)
	ObtenerResultadosPorExamen(ctx context.Context, idExamenMascota int) (Resultados, error)
	CrearExamenMascota(ctx context.Context, input CreateExamenMascotaRequest) (ExamenMascota, error)
	ActualizarExamenMascota(ctx context.Context, input UpdateExamenMascotaRequest) (ExamenMascota, error)
}

// ExamenesMascota represents the data about an examenesMascota.
type ExamenMascota struct {
	entity.ExamenMascota
}

type ExamenMascotaAll struct {
	IdExamenMascota int        `json:"id_examen_mascota"`
	IdUsuario       int        `json:"id_usuario"`
	IdMascota       int        `json:"id_mascota"`
	IdTipoExamen    int        `json:"id_tipo_examen"`
	FechaSolicitud  time.Time  `json:"fecha_solicitud"`
	FechaLlenado    *time.Time `json:"fecha_llenado"`
	Estado          string     `json:"estado"`
	Solicitante     string     `json:"solicitante"`
	Titulo          string     `json:"titulo"`
	Mascota         string     `json:"mascota"`
	Muestra         string     `json:"muestra"`
}

type ResultadosCualitativos struct {
	Parametro string       `json:"parametro" db:"parametro"`
	Resultado sql.NullBool `json:"resultado" db:"resultado"`
}

type ResultadosCuantitativos struct {
	Parametro              string  `json:"parametro" db:"parametro"`
	Resultado              float32 `json:"resultado" db:"resultado"`
	Unidad                 string  `json:"unidad" db:"unidad"`
	AlertaMenor            string  `json:"alerta_menor" db:"alerta_menor"`
	AlertaRango            string  `json:"alerta_rango" db:"alerta_rango"`
	AlertaMayor            string  `json:"alerta_mayor" db:"alerta_mayor"`
	RangoReferenciaInicial float32 `json:"rango_referencia_inicial" db:"rango_referencia_inicial"`
	RangoReferenciaFinal   float32 `json:"rango_referencia_final" db:"rango_referencia_final"`
}

type ResultadosInformativos struct {
	Parametro string `json:"parametro" db:"parametro"`
	Resultado string `json:"resultado" db:"resultado"`
}

type Resultados struct {
	Cualitativos  []ResultadosCualitativos  `json:"cualitativos"`
	Cuantitativos []ResultadosCuantitativos `json:"cuantitativos"`
	Informativos  []ResultadosInformativos  `json:"informativos"`
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new examenesMascota service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the list examenesMascota.
func (s service) GetExamenesMascota(ctx context.Context) ([]ExamenMascota, error) {
	examenesMascota, err := s.repo.GetExamenesMascota(ctx)
	if err != nil {
		return nil, err
	}
	result := []ExamenMascota{}
	for _, item := range examenesMascota {
		result = append(result, ExamenMascota{item})
	}
	return result, nil
}

func (s service) GetExamenesMascotaPorMascotayEstado(ctx context.Context, idMascota int, estado string) ([]ExamenMascotaAll, error) {
	examenesMascota, err := s.repo.GetExamenesMascotaPorMascotayEstado(ctx, idMascota, estado)
	if err != nil {
		return nil, err
	}
	return examenesMascota, nil
}

func (s service) GetExamenesMascotaPorEstado(ctx context.Context, estado string) ([]ExamenMascotaAll, error) {
	examenesMascota, err := s.repo.GetExamenesMascotaPorEstado(ctx, estado)
	if err != nil {
		return nil, err
	}
	return examenesMascota, nil
}

// CreateExamenMascotaRequest represents an examenesMascota creation request.
type CreateExamenMascotaRequest struct {
	IdUsuario      int        `json:"id_usuario"`
	IdMascota      int        `json:"id_mascota"`
	IdTipoExamen   int        `json:"id_tipo_examen"`
	FechaSolicitud time.Time  `json:"fecha_solicitud"`
	FechaLlenado   *time.Time `json:"fecha_llenado"`
	Estado         string     `json:"estado"`
	IdReferencia   int        `json:"id_referencia"`
	Tabla          string     `json:"tabla"`
}

type ResultadoRequest struct {
	Parametro string `json:"parametro"`
	Resultado string `json:"resultado"`
	Alerta    string `json:"alerta"`
}

type DatosMascotaRequest struct {
	Paciente     string
	Propietario  string
	Medico       string
	Muestra      string
	Especie      string
	Genero       string
	Raza         string
	FechaLlenado time.Time
}

type ResultadosRequest struct {
	Resultados []ResultadoRequest  `json:"resultados"`
	Datos      DatosMascotaRequest `json:"datos"`
}

type DatosMascotaDueñoRequest struct {
	NumAutorizacion int       `json:"num_autorizacion"`
	Paciente        string    `json:"paciente"`
	Propietario     string    `json:"propietario"`
	Nacionalidad    string    `json:"nacionalidad"`
	Cedula          string    `json:"cedula"`
	Sexo            string    `json:"sexo"`
	Direccion       string    `json:"direccion"`
	Raza            string    `json:"raza"`
	Edad            string    `json:"edad"`
	Enfermedad      string    `json:"enfermedad"`
	Intervencion    string    `json:"intervencion"`
	Profesional     string    `json:"profesional"`
	Abono           float32   `json:"abono"`
	Autoriza        string    `json:"autoriza"`
	Fecha           time.Time `json:"fecha"`
}

type UpdateExamenMascotaRequest struct {
	IdExamenMascota int        `json:"id_examen_mascota"`
	IdUsuario       int        `json:"id_usuario"`
	IdMascota       int        `json:"id_mascota"`
	IdTipoExamen    int        `json:"id_tipo_examen"`
	FechaSolicitud  time.Time  `json:"fecha_solicitud"`
	FechaLlenado    *time.Time `json:"fecha_llenado"`
	IdReferencia    int        `json:"id_referencia"`
	Tabla           string     `json:"tabla"`
	Estado          string     `json:"estado"`
	Valor           float32    `json:"valor"`
}

// Validate validates the UpdateExamenMascotaRequest fields.
func (m UpdateExamenMascotaRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdMascota, validation.Required),
		validation.Field(&m.IdUsuario, validation.Required),
		validation.Field(&m.IdTipoExamen, validation.Required),
	)
}

// Validate validates the CreateExamenMascotaRequest fields.
func (m CreateExamenMascotaRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.IdMascota, validation.Required),
		validation.Field(&m.IdUsuario, validation.Required),
		validation.Field(&m.IdTipoExamen, validation.Required),
	)
}

// CrearExamenMascota creates a new examenesMascota.
func (s service) CrearExamenMascota(ctx context.Context, req CreateExamenMascotaRequest) (ExamenMascota, error) {
	if err := req.Validate(); err != nil {
		return ExamenMascota{}, err
	}
	ExamenMascotaG, err := s.repo.CrearExamenMascota(ctx, entity.ExamenMascota{
		IdUsuario:      req.IdUsuario,
		IdMascota:      req.IdMascota,
		IdTipoExamen:   req.IdTipoExamen,
		FechaSolicitud: req.FechaSolicitud,
		FechaLlenado:   req.FechaLlenado,
		IdReferencia:   req.IdReferencia,
		Tabla:          req.Tabla,
		Estado:         req.Estado,
	})
	if err != nil {
		return ExamenMascota{}, err
	}
	return ExamenMascota{ExamenMascotaG}, nil
}

// ActualizarExamenMascota creates a new examenesMascota.
func (s service) ActualizarExamenMascota(ctx context.Context, req UpdateExamenMascotaRequest) (ExamenMascota, error) {
	if err := req.ValidateUpdate(); err != nil {
		return ExamenMascota{}, err
	}
	ExamenMascotaG, err := s.repo.ActualizarExamenMascota(ctx, entity.ExamenMascota{
		IdExamenMascota: req.IdExamenMascota,
		IdUsuario:       req.IdUsuario,
		IdMascota:       req.IdMascota,
		IdTipoExamen:    req.IdTipoExamen,
		FechaSolicitud:  req.FechaSolicitud,
		FechaLlenado:    req.FechaLlenado,
		Estado:          req.Estado,
		IdReferencia:    req.IdReferencia,
		Tabla:           req.Tabla,
	})
	if err != nil {
		return ExamenMascota{}, err
	}
	return ExamenMascota{ExamenMascotaG}, nil
}

// GetExamenMascotaPorId returns the examenesMascota with the specified the examenesMascota ID.
func (s service) GetExamenMascotaPorId(ctx context.Context, idExamenMascota int) (ExamenMascota, error) {
	examenesMascota, err := s.repo.GetExamenMascotaPorId(ctx, idExamenMascota)
	if err != nil {
		return ExamenMascota{}, err
	}
	return ExamenMascota{examenesMascota}, nil
}

func (s service) ObtenerResultadosPorExamen(ctx context.Context, idExamenMascota int) (Resultados, error) {
	resultados, err := s.repo.ObtenerResultadosPorExamen(ctx, idExamenMascota)
	if err != nil {
		return Resultados{}, err
	}
	return resultados, nil
}
