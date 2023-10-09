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
