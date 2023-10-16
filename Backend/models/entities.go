package models

import (
	"time"
)

type Paciente struct {
	ID              int64     `json:"id" query:"id"`
	Nome            string    `json:"nome" query:"nome"`
	Cpf             string    `json:"cpf" query:"cpf"`
	Sexo            string    `json:"sexo" query:"sexo"`
	Data_nascimento time.Time `json:"data_nascimento" query:"data_nascimento"`
	Cidade          string    `json:"cidade" query:"cidade"`
}

type Sinais_Vitais struct {
	ID_Paciente         int64     `json:"id_paciente" query:"id_paciente"`
	Frequencia_Cardiaca int64     `json:"frequencia_cardiaca" query:"frequencia_cardiaca"`
	Pressao_Arterial    string    `json:"pressao_arterial" query:"pressao_arterial"`
	Temperatura         float64   `json:"temperatura" query:"temperatura"`
	Saturacao_Oxigenio  float64   `json:"saturacao_oxigenio" query:"saturacao_oxigenio"`
	Data_Hora_Registro  time.Time `json:"data_hora_registro" query:"data_hora_registro"`
}
