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
	var examenesMascotaAll []ExamenMascotaAll
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
	var examenesMascotaAll []ExamenMascotaAll
	var nombre, apellido, titulo, nombreMascota string

	err := r.db.With(ctx).
		Select().
		From().
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
