package entity

type DetallesExamenInformativo struct {
	IdDetalleExamenInformativo int    `json:"id_detalle_examen_informativo" db:"pk,id_detalle_examen_informativo"`
	IdTipoExamen               int    `json:"id_tipo_examen" db:"id_tipo_examen"`
	Parametro                  string `json:"parametro" db:"parametro"`
}

func (d DetallesExamenInformativo) TableName() string {
	return "detalles_examen_informativo"
}
