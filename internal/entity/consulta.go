package entity

import "time"

type Consulta struct {
	IdConsulta             int       `json:"id_consulta" db:"pk,id_consulta"`
	IdMascota              int       `json:"id_mascota" db:"id_mascota"`
	IdUsuario              int       `json:"id_usuario" db:"id_usuario"`
	Fecha                  time.Time `json:"fecha" db:"fecha"`
	Valor                  float32   `json:"valor" db:"valor"`
	Motivo                 *string   `json:"motivo" db:"motivo"`
	Temperatura            *float32  `json:"temperatura" db:"temperatura"`
	Peso                   *float32  `json:"peso" db:"peso"`
	Tamaño                 *float32  `json:"tamanio" db:"tamaño"`
	CondicionCorporal      *string   `json:"condicion_corporal" db:"condicion_corporal"`
	NivelesDeshidratacion  *string   `json:"niveles_deshidratacion" db:"niveles_deshidratacion"`
	Diagnostico            *string   `json:"diagnostico" db:"diagnostico"`
	Edad                   *string   `json:"edad" db:"edad"`
	TiempoLlenadoCapilar   int       `json:"tiempo_llenado_capilar" db:"tiempo_llenado_capilar"`
	FrecuenciaCardiaca     int       `json:"frecuencia_cardiaca" db:"frecuencia_cardiaca"`
	FrecuenciaRespiratoria int       `json:"frecuencia_respiratoria" db:"frecuencia_respiratoria"`
	EstadoConsulta         string    `json:"estado_consulta" db:"estado_consulta"`
}

func (c Consulta) TableName() string {
	return "consulta"
}
