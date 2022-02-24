package entity

type ResultadoDetalleInformativo struct {
	IdResultadoDetalleInformativo int    `json:"id_resultado_detalle_informativo" db:"pk,id_resultado_detalle_informativo"`
	IdExamenMascota               int    `json:"id_examen_mascota" db:"id_examen_mascota"`
	IdDetalleExamenInformativo    int    `json:"id_detalle_examen_informativo" db:"id_detalle_examen_informativo"`
	Resultado                     string `json:"resultado" db:"resultado"`
}

func (r ResultadoDetalleInformativo) TableName() string {
	return "resultados_detalle_informativo"
}
