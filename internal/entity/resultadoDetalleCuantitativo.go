package entity

type ResultadoDetalleCuantitativo struct {
	IdResultadoDetalleCuantitativo int     `json:"id_resultado_detalle_cuantitativo" db:"pk,id_resultado_detalle_cuantitativo"`
	IdExamenMascota                int     `json:"id_examen_mascota" db:"id_examen_mascota"`
	IdDetalleExamenCuantitativo    int     `json:"id_detalle_examen_cuantitativo" db:"id_detalle_examen_cuantitativo"`
	Resultado                      float32 `json:"resultado" db:"resultado"`
}

func (r ResultadoDetalleCuantitativo) TableName() string {
	return "resultados_detalle_cuantitativo"
}
