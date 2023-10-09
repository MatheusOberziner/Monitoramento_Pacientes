package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/MatheusOberziner/Monitoramento_Pacientes/db"
	"github.com/MatheusOberziner/Monitoramento_Pacientes/models"
	"github.com/labstack/echo/v4"
)

type CreatePacienteRequest struct {
	Nome            string `json:"nome" query:"nome"`
	Cpf             string `json:"cpf" query:"cpf"`
	Sexo            string `json:"sexo" query:"sexo"`
	Data_nascimento string `json:"data_nascimento" query:"data_nascimento"`
	Cidade          string `json:"cidade" query:"cidade"`
}

func post(paciente models.Paciente) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO pacientes (nome, cpf, sexo, data_nascimento, cidade)
				VALUES ($1, $2, $3, $4, $5)
				RETURNING id`

	err = conn.QueryRow(sql, paciente.Nome, paciente.Cpf, paciente.Sexo, paciente.Data_nascimento, paciente.Cidade).Scan(&id)
	return
}

// Cria um endpoint responsável por adicionar um paciente
func CreatePaciente(c echo.Context) error {
	request := new(CreatePacienteRequest)
	if err := c.Bind(request); err != nil {
		log.Println("Erro no bind da requisição:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Erro ao fazer bind da requisição"})
	}

	dataNascimento, err := time.Parse("02/01/2006", request.Data_nascimento)
	if err != nil {
		log.Println("Erro ao converter a data:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Formato de data inválido"})
	}

	paciente := models.Paciente{
		Nome:            request.Nome,
		Cpf:             request.Cpf,
		Sexo:            request.Sexo,
		Cidade:          request.Cidade,
		Data_nascimento: dataNascimento,
	}

	id, err := post(paciente)
	if err != nil {
		log.Println("Erro: ", err)
		return c.String(http.StatusInternalServerError, "Erro ao adicionar um paciente")
	}

	paciente.ID = id

	return c.JSON(http.StatusOK, paciente)
}
