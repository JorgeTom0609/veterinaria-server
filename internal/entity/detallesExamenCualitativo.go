package entity

type DetallesExamenCualitativo struct {
	IdDetalleExamenCualitativo int    `json:"id_detalle_examen_cualitativo" db:"pk,id_detalle_examen_cualitativo"`
	IdTipoExamen               int    `json:"id_tipo_examen" db:"id_tipo_examen"`
	Parametro                  string `json:"parametro" db:"parametro"`
}

func (d DetallesExamenCualitativo) TableName() string {
	return "detalles_examen_cualitativo"
}
