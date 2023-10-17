package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/MatheusOberziner/Monitoramento_Pacientes/db"
	"github.com/MatheusOberziner/Monitoramento_Pacientes/models"
	"github.com/labstack/echo/v4"
)

func getSinais(idPaciente int64, dataFiltro string) (lista []models.Sinais_Vitais, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	sql := `SELECT * FROM sinais_vitais WHERE id_paciente = $1`
	params := []interface{}{idPaciente}

	if dataFiltro != "" {
		sql += " AND DATE(data_hora_registro) = $2"
		params = append(params, dataFiltro)
	}

	rows, err := conn.Query(sql, params...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var sinais_vitais models.Sinais_Vitais

		err = rows.Scan(&sinais_vitais.ID_Paciente, &sinais_vitais.Frequencia_Cardiaca, &sinais_vitais.Pressao_Arterial, &sinais_vitais.Temperatura, &sinais_vitais.Saturacao_Oxigenio, &sinais_vitais.Data_Hora_Registro)
		if err != nil {
			log.Printf("Erro no scan: %v", err)
			continue
		}

		lista = append(lista, sinais_vitais)
	}

	return lista, nil
}

// 2a e 5 Cria um endpoint responsável por listar o histórico de sinais vitais do paciente
func ListSinaisVitais(c echo.Context) error {
	idPaciente, err := strconv.Atoi(c.Param("id_paciente"))
	if err != nil {
		return c.String(http.StatusBadRequest, "ID de paciente inválido")
	}

	dataFiltro := c.QueryParam("data")

	lista, err := getSinais(int64(idPaciente), dataFiltro)
	if err != nil {
		log.Println("Error", err)
		return c.String(http.StatusInternalServerError, "Erro ao listar pacientes")
	}

	return c.JSON(http.StatusOK, lista)
}
