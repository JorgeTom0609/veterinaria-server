package entity

import "time"

// Rol represents a rol.
type ExamenMascota struct {
	IdExamenMascota int        `json:"id_examen_mascota" db:"pk,id_examen_mascota"`
	IdUsuario       int        `json:"id_usuario" db:"id_usuario"`
	IdMascota       int        `json:"id_mascota" db:"id_mascota"`
	IdTipoExamen    int        `json:"id_tipo_examen" db:"id_tipo_examen"`
	FechaSolicitud  time.Time  `json:"fecha_solicitud" db:"fecha_solicitud"`
	FechaLlenado    *time.Time `json:"fecha_llenado" db:"fecha_llenado"`
	Estado          string     `json:"estado" db:"estado"`
}

func (em ExamenMascota) TableName() string {
	return "examenes_mascota"
}
