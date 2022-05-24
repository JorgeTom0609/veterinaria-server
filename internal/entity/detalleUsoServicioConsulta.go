package entity

type DetalleUsoServicioConsulta struct {
	IdDetalleUsoServicioConsulta int     `json:"id_detalle_uso_servicio_consulta" db:"pk,id_detalle_uso_servicio_consulta"`
	IdDetalleServicioConsulta    int     `json:"id_detalle_servicio_consulta" db:"id_detalle_servicio_consulta"`
	IdReferencia                 int     `json:"id_referencia" db:"id_referencia"`
	Tabla                        string  `json:"tabla" db:"tabla"`
	Cantidad                     float32 `json:"cantidad" db:"cantidad"`
}

func (dus DetalleUsoServicioConsulta) TableName() string {
	return "detalle_usos_servicio_consulta"
}
