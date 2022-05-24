package entity

type DetalleUsoServicio struct {
	IdDetalleUsoServicio             int     `json:"id_detalle_uso_servicio" db:"pk,id_detalle_uso_servicio"`
	IdDetalleServicioHospitalizacion int     `json:"id_detalle_servicio_hospitalizacion" db:"id_detalle_servicio_hospitalizacion"`
	IdReferencia                     int     `json:"id_referencia" db:"id_referencia"`
	Tabla                            string  `json:"tabla" db:"tabla"`
	Cantidad                         float32 `json:"cantidad" db:"cantidad"`
}

func (dus DetalleUsoServicio) TableName() string {
	return "detalle_usos_servicio"
}
