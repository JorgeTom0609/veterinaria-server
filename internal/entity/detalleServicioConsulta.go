package entity

import "time"

type DetalleServicioConsulta struct {
	IdDetalleServicioConsulta int       `json:"id_detalle_servicio_consulta" db:"pk,id_detalle_servicio_consulta"`
	IdConsulta                int       `json:"id_consulta" db:"id_consulta"`
	IdServicio                int       `json:"id_servicio" db:"id_servicio"`
	Valor                     float32   `json:"valor" db:"valor"`
	Fecha                     time.Time `json:"fecha" db:"fecha"`
}

func (dsc DetalleServicioConsulta) TableName() string {
	return "detalles_servicios_hospitalizacion"
}
