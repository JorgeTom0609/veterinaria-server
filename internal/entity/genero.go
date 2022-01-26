package entity

type Genero struct {
	IdGenero    int    `json:"id_genero" db:"pk,id_genero"`
	Descripcion string `json:"descripcion" db:"descripcion"`
}

func (g Genero) TableName() string {
	return "generos"
}
