package handlers

import (
	"net/http"
	"strconv"

	"github.com/MatheusOberziner/Monitoramento_Pacientes/db"
	"github.com/MatheusOberziner/Monitoramento_Pacientes/models"
	"github.com/labstack/echo/v4"
)

func get(id int64) (*models.Paciente, error) {
	var paciente models.Paciente
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	sql := `SELECT * FROM pacientes WHERE id = $1`

	row := conn.QueryRow(sql, id)

	err = row.Scan(&paciente.ID, &paciente.Nome, &paciente.Cpf, &paciente.Sexo, &paciente.Data_nascimento, &paciente.Cidade)
	if err != nil {
		return nil, err
	}

	return &paciente, nil
}

func GetPaciente(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Erro ao pegar o id do paciente")
	}

	paciente, err := get(int64(id))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Erro em pegar dados do paciente")
	}

	return c.JSON(http.StatusOK, paciente)
}
