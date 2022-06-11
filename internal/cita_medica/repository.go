package cita_medica

import (
	"context"
	"veterinaria-server/internal/entity"
	"veterinaria-server/pkg/dbcontext"
	"veterinaria-server/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access citasMedica from the data source.
type Repository interface {
	// GetCitaMedicaPorId returns the citaMedica with the specified citaMedica ID.
	GetCitaMedicaPorId(ctx context.Context, idCitaMedica int) (entity.CitaMedica, error)
	// GetCitasMedica returns the list citasMedica.
	GetCitasMedica(ctx context.Context) ([]entity.CitaMedica, error)
	GetCitasMedicaPendientes(ctx context.Context) ([]entity.CitaMedica, error)
	GetCitasMedicaSinNotificar(ctx context.Context) ([]CitaMedicaDatos, error)
	CrearCitaMedica(ctx context.Context, citaMedica entity.CitaMedica) (entity.CitaMedica, error)
	ActualizarCitaMedica(ctx context.Context, citaMedica entity.CitaMedica) (entity.CitaMedica, error)
}

// repository persists citasMedica in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new citaMedica repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the list citasMedica from the database.
func (r repository) GetCitasMedica(ctx context.Context) ([]entity.CitaMedica, error) {
	var citasMedica []entity.CitaMedica

	err := r.db.With(ctx).
		Select().
		From().
		All(&citasMedica)
	if err != nil {
		return citasMedica, err
	}
	return citasMedica, err
}

func (r repository) GetCitasMedicaPendientes(ctx context.Context) ([]entity.CitaMedica, error) {
	var citasMedica []entity.CitaMedica

	err := r.db.With(ctx).
		Select().
		From().
		Where(dbx.NewExp("DATE(now()) <= fecha")).
		All(&citasMedica)
	if err != nil {
		return citasMedica, err
	}
	return citasMedica, err
}

func (r repository) GetCitasMedicaSinNotificar(ctx context.Context) ([]CitaMedicaDatos, error) {
	var citasMedica []entity.CitaMedica
	var citasMedicasDatos []CitaMedicaDatos = []CitaMedicaDatos{}

	err := r.db.With(ctx).
		Select().
		From().
		Where(dbx.NewExp("(date(fecha) between date(now()) and date_add(date(now()),interval 3 day)) and estado_notificacion = 'NO'")).
		All(&citasMedica)

	if err != nil {
		return []CitaMedicaDatos{}, err
	}

	for i := 0; i < len(citasMedica); i++ {
		var duenioNombre, duenioApellido, telefono, nombreMascota string
		duenioNombre = ""
		duenioApellido = ""
		telefono = ""
		nombreMascota = ""
		err := r.db.With(ctx).
			Select("c.nombres", "c.apellidos", "c.telefono", "m.nombre").
			From("mascotas m").
			InnerJoin("clientes c", dbx.NewExp("c.id_cliente = m.id_cliente")).
			Where(dbx.HashExp{"m.id_mascota": citasMedica[i].IdMascota}).
			Row(&duenioNombre, &duenioApellido, &telefono, &nombreMascota)

		if err != nil {
			return []CitaMedicaDatos{}, err
		}

		citasMedicasDatos = append(citasMedicasDatos, CitaMedicaDatos{
			citasMedica[i],
			duenioApellido + " " + duenioNombre,
			telefono,
			nombreMascota,
		})
	}

	return citasMedicasDatos, err
}

// Create saves a new CitaMedica record in the database.
// It returns the ID of the newly inserted citaMedica record.
func (r repository) CrearCitaMedica(ctx context.Context, citaMedica entity.CitaMedica) (entity.CitaMedica, error) {
	err := r.db.With(ctx).Model(&citaMedica).Insert()
	if err != nil {
		return entity.CitaMedica{}, err
	}
	return citaMedica, nil
}

// Create saves a new CitaMedica record in the database.
// It returns the ID of the newly inserted citaMedica record.
func (r repository) ActualizarCitaMedica(ctx context.Context, citaMedica entity.CitaMedica) (entity.CitaMedica, error) {
	var err error
	if citaMedica.IdCitaMedica != 0 {
		err = r.db.With(ctx).Model(&citaMedica).Update()
	} else {
		err = r.db.With(ctx).Model(&citaMedica).Insert()
	}
	if err != nil {
		return entity.CitaMedica{}, err
	}
	return citaMedica, nil
}

// Get reads the citaMedica with the specified ID from the database.
func (r repository) GetCitaMedicaPorId(ctx context.Context, idCitaMedica int) (entity.CitaMedica, error) {
	var citaMedica entity.CitaMedica
	err := r.db.With(ctx).Select().Model(idCitaMedica, &citaMedica)
	return citaMedica, err
}
