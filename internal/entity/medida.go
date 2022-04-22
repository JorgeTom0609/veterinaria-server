package entity

type Medida struct {
	IdMedida    int    `json:"id_medida" db:"pk,id_medida"`
	Descripcion string `json:"descripcion" db:"descripcion"`
}

func (p Medida) TableName() string {
	return "medida"
}
