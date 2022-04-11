package entity

type Servicio struct {
	IdServicio  int     `json:"id_servicio" db:"pk,id_servicio"`
	Descripcion string  `json:"descripcion" db:"descripcion"`
	Valor       float32 `json:"valor" db:"valor"`
}

func (c Servicio) TableName() string {
	return "servicio"
}
