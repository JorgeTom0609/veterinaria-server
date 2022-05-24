package entity

import "database/sql"

type Servicio struct {
	IdServicio            int          `json:"id_servicio" db:"pk,id_servicio"`
	IdUsuario             int          `json:"id_usuario" db:"id_usuario"`
	IdEspecie             int          `json:"id_especie" db:"id_especie"`
	Descripcion           string       `json:"descripcion" db:"descripcion"`
	Valor                 float32      `json:"valor" db:"valor"`
	AplicaConsulta        sql.NullBool `json:"aplica_consulta" db:"aplica_consulta"`
	AplicaHospitalizacion sql.NullBool `json:"aplica_hospitalizacion" db:"aplica_hospitalizacion"`
}

func (c Servicio) TableName() string {
	return "servicios"
}
