package entity

import "database/sql"

type ResultadoDetalleCualitativo struct {
	IdResultadoDetalleCualitativo int          `json:"id_resultado_detalle_cualitativo" db:"pk,id_resultado_detalle_cualitativo"`
	IdExamenMascota               int          `json:"id_examen_mascota" db:"id_examen_mascota"`
	IdDetalleExamenCualitativo    int          `json:"id_detalle_examen_cualitativo" db:"id_detalle_examen_cualitativo"`
	Resultado                     sql.NullBool `json:"resultado" db:"resultado"`
}

func (r ResultadoDetalleCualitativo) TableName() string {
	return "resultados_detalle_cualitativo"
}
