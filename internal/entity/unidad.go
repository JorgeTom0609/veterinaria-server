package entity

type Unidad struct {
	IdUnidad    int    `json:"id_unidad" db:"pk,id_unidad"`
	IdMedida    int    `json:"id_medida" db:"id_medida"`
	Descripcion string `json:"descripcion" db:"descripcion"`
}

func (p Unidad) TableName() string {
	return "unidad"
}
