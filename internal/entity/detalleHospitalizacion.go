package entity

import "time"

type DetalleHospitalizacion struct {
	IdDetalleHospitalizacion int       `json:"id_detalle_hospitalizacion" db:"pk,id_detalle_hospitalizacion"`
	IdHospitalizacion        int       `json:"id_hospitalizacion" db:"id_hospitalizacion"`
	IdUsuario                int       `json:"id_usuario" db:"id_usuario"`
	Descripcion              string    `json:"descripcion" db:"descripcion"`
	Fecha                    time.Time `json:"fecha" db:"fecha"`
}

func (dh DetalleHospitalizacion) TableName() string {
	return "detalles_hospitalizacion"
}
