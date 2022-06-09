package entity

import "time"

type CitaMedica struct {
	IdCitaMedica       int       `json:"id_cita_medica" db:"pk,id_cita_medica"`
	IdMascota          int       `json:"id_mascota" db:"id_mascota"`
	Motivo             string    `json:"motivo" db:"motivo"`
	Fecha              time.Time `json:"fecha" db:"fecha"`
	EstadoNotificacion string    `json:"estado_notificacion" db:"estado_notificacion"`
}

func (c CitaMedica) TableName() string {
	return "citas_medicas"
}
