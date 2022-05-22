package entity

type Receta struct {
	IdReceta     int    `json:"id_receta" db:"pk,id_receta"`
	IdProducto   int    `json:"id_producto" db:"id_producto"`
	IdConsulta   int    `json:"id_consulta" db:"id_consulta"`
	Prescripcion string `json:"prescripcion" db:"prescripcion"`
}

func (r Receta) TableName() string {
	return "receta"
}
