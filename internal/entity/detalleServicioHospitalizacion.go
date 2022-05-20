package entity

import "time"

type DetalleServicioHospitalizacion struct {
	IdDetalleServicioHospitalizacion int       `json:"id_detalle_servicio_hospitalizacion" db:"pk,id_detalle_servicio_hospitalizacion"`
	IdHospitalizacion                int       `json:"id_hospitalizacion" db:"id_hospitalizacion"`
	IdUsuario                        int       `json:"id_usuario" db:"id_usuario"`
	IdServicio                       int       `json:"id_servicio" db:"id_servicio"`
	Descripcion                      string    `json:"descripcion" db:"descripcion"`
	Fecha                            time.Time `json:"fecha" db:"fecha"`
}

func (dsh DetalleServicioHospitalizacion) TableName() string {
	return "detalles_servicios_hospitalizacion"
}
