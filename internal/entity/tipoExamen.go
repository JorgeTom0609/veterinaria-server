package entity

type TipoExamen struct {
	IdTipoExamen int     `json:"id_tipo_examen" db:"pk,id_tipo_examen"`
	IdEspecie    int     `json:"id_especie" db:"id_especie"`
	Descripcion  *string `json:"descripcion" db:"descripcion"`
	Muestra      *string `json:"muestra" db:"muestra"`
}

func (t TipoExamen) TableName() string {
	return "tipos_examenes"
}