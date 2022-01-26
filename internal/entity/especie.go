package entity

type Especie struct {
	IdEspecie   int    `json:"id_especie" db:"pk,id_especie"`
	Descripcion string `json:"descripcion" db:"descripcion"`
}

func (e Especie) TableName() string {
	return "especies"
}
