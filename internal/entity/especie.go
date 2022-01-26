package entity

type Especie struct {
	IdEspecie   int    `json:"id_especie" db:"id_especie"`
	Descripcion string `json:"descripcion" db:"descripcion"`
}
