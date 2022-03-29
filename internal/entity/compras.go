package entity

import "time"

type Compras struct {
	IdCompra    int       `json:"id_compra" db:"pk,id_compra"`
	IdUsuario   int       `json:"id_usuario" db:"id_usuario"`
	Fecha       time.Time `json:"fecha" db:"fecha"`
	Valor       float32   `json:"valor" db:"valor"`
	Descripcion *string   `json:"descripcion" db:"descripcion"`
}

func (c Compras) TableName() string {
	return "compras"
}
