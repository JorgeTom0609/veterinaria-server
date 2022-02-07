package entity

type DetallesExamenCuantitativo struct {
	IdDetalleExamenCuantitativo int     `json:"id_detalle_examen_cuantitativo" db:"pk,id_detalle_examen_cuantitativo"`
	IdTipoExamen                int     `json:"id_tipo_examen" db:"id_tipo_examen"`
	Parametro                   string  `json:"parametro" db:"parametro"`
	RangoReferenciaInicial      float32 `json:"rango_referencia_inicial" db:"rango_referencia_inicial"`
	RangoReferenciaFinal        float32 `json:"rango_referencia_final" db:"rango_referencia_final"`
	Unidad                      *string `json:"unidad" db:"unidad"`
	AlertaMenor                 *string `json:"alerta_menor" db:"alerta_menor"`
	AlertaRango                 *string `json:"alerta_rango" db:"alerta_rango"`
	AlertaMayor                 *string `json:"alerta_mayor" db:"alerta_mayor"`
}

func (d DetallesExamenCuantitativo) TableName() string {
	return "detalles_examen_cuantitativo"
}
