package examen_mascota

import (
	"context"
	"time"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access examenesMascota from the data source.
type Repository interface {
	// GetExamenMascotaPorId returns the examenesMascota with the specified examenesMascota ID.
	GetExamenMascotaPorId(ctx context.Context, idExamenMascota int) (entity.ExamenMascota, error)
	// GetExamenesMascota returns the list examenesMascota.
	GetExamenesMascota(ctx context.Context) ([]entity.ExamenMascota, error)
	GetExamenesMascotaPorMascotayEstado(ctx context.Context, idExamenMascota int, estado string) ([]ExamenMascotaAll, error)
	GetExamenesMascotaPorEstado(ctx context.Context, estado string) ([]ExamenMascotaAll, error)
	ObtenerResultadosPorExamen(ctx context.Context, idExamenMascota int) (Resultados, error)
	CrearExamenMascota(ctx context.Context, examenesMascota entity.ExamenMascota) (entity.ExamenMascota, error)
	ActualizarExamenMascota(ctx context.Context, examenesMascota entity.ExamenMascota) (entity.ExamenMascota, error)
}

// repository persists examenesMascota in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new examenesMascota repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list examenesMascota from the database.
func (r repository) GetExamenesMascota(ctx context.Context) ([]entity.ExamenMascota, error) {
	var examenesMascota []entity.ExamenMascota

	err := r.db.With(ctx).
		Select().
		From().
		All(&examenesMascota)
	if err != nil {
		return examenesMascota, err
	}
	return examenesMascota, err
}

func (r repository) GetExamenesMascotaPorMascotayEstado(ctx context.Context, idMascota int, estado string) ([]ExamenMascotaAll, error) {
	var examenesMascota []entity.ExamenMascota
	var examenesMascotaAll []ExamenMascotaAll = []ExamenMascotaAll{}
	var nombre, apellido, titulo string

	err := r.db.With(ctx).
		Select().
		From().
		Where(dbx.HashExp{"id_mascota": idMascota}).
		AndWhere(dbx.HashExp{"estado": estado}).
		All(&examenesMascota)

	for i := 0; i < len(examenesMascota); i++ {
		idUsuario := examenesMascota[i].IdUsuario
		err := r.db.With(ctx).
			Select("apellido", "nombre").
			From("usuarios").
			Where(dbx.HashExp{"id_usuario": idUsuario}).
			Row(&nombre, &apellido)
		if err != nil {
			return []ExamenMascotaAll{}, err
		}

		idTipoDeExamen := examenesMascota[i].IdTipoExamen
		err = r.db.With(ctx).
			Select("titulo").
			From("tipos_examenes").
			Where(dbx.HashExp{"id_tipo_examen": idTipoDeExamen}).
			Row(&titulo)
		if err != nil {
			return []ExamenMascotaAll{}, err
		}

		examenesMascotaAll = append(examenesMascotaAll, ExamenMascotaAll{
			examenesMascota[i].IdExamenMascota,
			examenesMascota[i].IdUsuario,
			examenesMascota[i].IdMascota,
			examenesMascota[i].IdTipoExamen,
			examenesMascota[i].FechaSolicitud,
			examenesMascota[i].FechaLlenado,
			examenesMascota[i].Estado,
			apellido + " " + nombre,
			titulo,
			"",
		})
	}

	if err != nil {
		return []ExamenMascotaAll{}, err
	}
	return examenesMascotaAll, err
}

func (r repository) GetExamenesMascotaPorEstado(ctx context.Context, estado string) ([]ExamenMascotaAll, error) {
	var examenesMascota []entity.ExamenMascota
	var examenesMascotaAll []ExamenMascotaAll = []ExamenMascotaAll{}
	var nombre, apellido, titulo, nombreMascota string

	err := r.db.With(ctx).
		Select().
		From().
		Where(dbx.HashExp{"estado": estado}).
		All(&examenesMascota)

	for i := 0; i < len(examenesMascota); i++ {
		idUsuario := examenesMascota[i].IdUsuario
		err := r.db.With(ctx).
			Select("apellido", "nombre").
			From("usuarios").
			Where(dbx.HashExp{"id_usuario": idUsuario}).
			Row(&nombre, &apellido)
		if err != nil {
			return []ExamenMascotaAll{}, err
		}

		idTipoDeExamen := examenesMascota[i].IdTipoExamen
		err = r.db.With(ctx).
			Select("titulo").
			From("tipos_examenes").
			Where(dbx.HashExp{"id_tipo_examen": idTipoDeExamen}).
			Row(&titulo)
		if err != nil {
			return []ExamenMascotaAll{}, err
		}

		idMascota := examenesMascota[i].IdMascota
		err = r.db.With(ctx).
			Select("nombre").
			From("mascotas").
			Where(dbx.HashExp{"id_mascota": idMascota}).
			Row(&nombreMascota)
		if err != nil {
			return []ExamenMascotaAll{}, err
		}

		examenesMascotaAll = append(examenesMascotaAll, ExamenMascotaAll{
			examenesMascota[i].IdExamenMascota,
			examenesMascota[i].IdUsuario,
			examenesMascota[i].IdMascota,
			examenesMascota[i].IdTipoExamen,
			examenesMascota[i].FechaSolicitud,
			examenesMascota[i].FechaLlenado,
			examenesMascota[i].Estado,
			apellido + " " + nombre,
			titulo,
			nombreMascota,
		})
	}

	if err != nil {
		return []ExamenMascotaAll{}, err
	}
	return examenesMascotaAll, err
}

// Create saves a new ExamenMascota record in the database.
// It returns the ID of the newly inserted examenesMascota record.
func (r repository) CrearExamenMascota(ctx context.Context, examenesMascota entity.ExamenMascota) (entity.ExamenMascota, error) {
	err := r.db.With(ctx).Model(&examenesMascota).Insert()
	if err != nil {
		return entity.ExamenMascota{}, err
	}
	return examenesMascota, nil
}

// Create saves a new ExamenMascota record in the database.
// It returns the ID of the newly inserted examenesMascota record.
func (r repository) ActualizarExamenMascota(ctx context.Context, examenesMascota entity.ExamenMascota) (entity.ExamenMascota, error) {
	var err error
	if examenesMascota.IdExamenMascota != 0 {
		err = r.db.With(ctx).Model(&examenesMascota).Update()
	} else {
		err = r.db.With(ctx).Model(&examenesMascota).Insert()
	}
	if err != nil {
		return entity.ExamenMascota{}, err
	}
	return examenesMascota, nil
}

// Get reads the examenesMascota with the specified ID from the database.
func (r repository) GetExamenMascotaPorId(ctx context.Context, idExamenMascota int) (entity.ExamenMascota, error) {
	var examenesMascota entity.ExamenMascota
	err := r.db.With(ctx).Select().Model(idExamenMascota, &examenesMascota)
	return examenesMascota, err
}

func ActualizarEstadoExamenMascota(ctx context.Context, idExamenMascota int, db *dbcontext.DB) (bool, error) {
	var examenMascota entity.ExamenMascota
	examenMascota.IdExamenMascota = idExamenMascota
	fecha := time.Now()
	examenMascota.FechaLlenado = &fecha
	examenMascota.Estado = "FINALIZADO"
	err := db.With(ctx).Model(&examenMascota).Update("FechaLlenado", "Estado")
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r repository) ObtenerResultadosPorExamen(ctx context.Context, idExamenMascota int) (Resultados, error) {
	var resultadosCualitativos []ResultadosCualitativos = []ResultadosCualitativos{}
	var resultadosCuantitativos []ResultadosCuantitativos = []ResultadosCuantitativos{}
	var resultadosInformativos []ResultadosInformativos = []ResultadosInformativos{}
	err := r.db.With(ctx).
		Select("parametro", "resultado").
		From("examenes_mascota as em").
		InnerJoin("resultados_detalle_cualitativo as rdc", dbx.NewExp("rdc.id_examen_mascota = em.id_examen_mascota")).
		InnerJoin("detalles_examen_cualitativo as dc", dbx.NewExp("dc.id_detalle_examen_cualitativo = rdc.id_detalle_examen_cualitativo")).
		Where(dbx.HashExp{"em.id_examen_mascota": idExamenMascota}).
		All(&resultadosCualitativos)
	if err != nil {
		return Resultados{}, err
	}

	err = r.db.With(ctx).
		Select("parametro", "resultado", "unidad", "alerta_menor", "alerta_rango", "alerta_mayor", "rango_referencia_inicial", "rango_referencia_final").
		From("examenes_mascota as em").
		InnerJoin("resultados_detalle_cuantitativo as rdc", dbx.NewExp("rdc.id_examen_mascota = em.id_examen_mascota")).
		InnerJoin("detalles_examen_cuantitativo as dc", dbx.NewExp("dc.id_detalle_examen_cuantitativo = rdc.id_detalle_examen_cuantitativo")).
		Where(dbx.HashExp{"em.id_examen_mascota": idExamenMascota}).
		All(&resultadosCuantitativos)
	if err != nil {
		return Resultados{}, err
	}

	err = r.db.With(ctx).
		Select("parametro", "resultado").
		From("examenes_mascota as em").
		InnerJoin("resultados_detalle_informativo as rdi", dbx.NewExp("rdi.id_examen_mascota = em.id_examen_mascota")).
		InnerJoin("detalles_examen_informativo as di", dbx.NewExp("di.id_detalle_examen_informativo = rdi.id_detalle_examen_informativo")).
		Where(dbx.HashExp{"em.id_examen_mascota": idExamenMascota}).
		All(&resultadosInformativos)
	if err != nil {
		return Resultados{}, err
	}

	return Resultados{resultadosCualitativos, resultadosCuantitativos, resultadosInformativos}, nil
}
