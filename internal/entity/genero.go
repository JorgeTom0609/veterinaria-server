package entity

type Genero struct {
	IdGenero    int    `json:"id_genero" db:"id_genero"`
	Descripcion string `json:"descripcion" db:"descripcion"`
}
