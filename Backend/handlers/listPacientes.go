package handlers

import (
	"log"
	"net/http"

	"github.com/MatheusOberziner/Monitoramento_Pacientes/db"
	"github.com/MatheusOberziner/Monitoramento_Pacientes/models"
	"github.com/labstack/echo/v4"
)

func getAll(nomeFilter string, cidadeFilter string) (pacientes []models.Paciente, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	sql := `SELECT * FROM pacientes`
	params := []interface{}{}

	if nomeFilter != "" {
		sql += ` WHERE nome ILIKE $1`
		params = append(params, "%"+nomeFilter+"%")
	}

	if cidadeFilter != "" {
		sql += ` WHERE cidade ILIKE $1`
		params = append(params, "%"+cidadeFilter+"%")
	}

	rows, err := conn.Query(sql, params...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var paciente models.Paciente

		err = rows.Scan(&paciente.ID, &paciente.Nome, &paciente.Cpf, &paciente.Sexo, &paciente.Data_nascimento, &paciente.Cidade)
		if err != nil {
			log.Printf("Erro no scan: %v", err)
			continue
		}

		pacientes = append(pacientes, paciente)
	}

	return pacientes, nil
}

// Cria um endpoint responsável por listar os pacientes
func ListPacientes(c echo.Context) error {
	nomeFilter := c.QueryParam("nome")
	cidadeFilter := c.QueryParam("cidade")

	pacientes, err := getAll(nomeFilter, cidadeFilter)
	if err != nil {
		log.Println("Error", err)
		return c.String(http.StatusInternalServerError, "Erro ao listar pacientes")
	}

	return c.JSON(http.StatusOK, pacientes)
}
