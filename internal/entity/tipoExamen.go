package entity

type TipoExamen struct {
	IdTipoExamen int     `json:"id_tipo_examen" db:"pk,id_tipo_examen"`
	IdEspecie    int     `json:"id_especie" db:"id_especie"`
	Titulo       string  `json:"titulo" db:"titulo"`
	Descripcion  string  `json:"descripcion" db:"descripcion"`
	Muestra      string  `json:"muestra" db:"muestra"`
	Valor        float32 `json:"valor" db:"valor"`
}

func (t TipoExamen) TableName() string {
	return "tipos_examenes"
}
