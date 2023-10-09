package handlers

import (
	"github.com/MatheusOberziner/Monitoramento_Pacientes/db"
	"github.com/MatheusOberziner/Monitoramento_Pacientes/models"
	"github.com/labstack/echo/v4"
)

func Post(paciente models.Paciente) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO pacientes (nome, cpf, sexo, data_nascimento, cidade)
				VALUES (%1, %2, %3, %4, %5)
				RETURNING id`

	err = conn.QueryRow(sql, paciente.Nome, paciente.Cpf, paciente.Sexo, paciente.Data_nascimento, paciente.Cidade).Scan(&id)
	return
}

func CreatePaciente(c echo.Context) error {

}
