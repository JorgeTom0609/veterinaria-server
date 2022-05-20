package entity

type Servicio struct {
	IdServicio  int     `json:"id_servicio" db:"pk,id_servicio"`
	IdUsuario   int     `json:"id_usuario" db:"id_usuario"`
	IdEspecie   int     `json:"id_especie" db:"id_especie"`
	Descripcion string  `json:"descripcion" db:"descripcion"`
	Valor       float32 `json:"valor" db:"valor"`
}

func (c Servicio) TableName() string {
	return "servicios"
}
