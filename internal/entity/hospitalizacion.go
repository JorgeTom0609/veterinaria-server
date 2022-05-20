package entity

import (
	"database/sql"
	"time"
)

type Hospitalizacion struct {
	IdHospitalizacion     int          `json:"id_hospitalizacion" db:"pk,id_hospitalizacion"`
	IdConsulta            int          `json:"id_consulta" db:"id_consulta"`
	Motivo                string       `json:"motivo" db:"motivo"`
	FechaIngreso          time.Time    `json:"fecha_ingreso" db:"fecha_ingreso"`
	FechaSalida           *time.Time   `json:"fecha_salida" db:"fecha_salida"`
	Valor                 *float32     `json:"valor" db:"valor"`
	Abono                 float32      `json:"abono" db:"abono"`
	AuorizaExamenes       sql.NullBool `json:"autoriza_examenes" db:"autoriza_examenes"`
	EstadoHospitalizacion string       `json:"estado_hospitalizacion" db:"estado_hospitalizacion"`
}

func (h Hospitalizacion) TableName() string {
	return "hospitalizacion"
}
