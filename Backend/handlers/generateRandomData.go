package handlers

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/MatheusOberziner/Monitoramento_Pacientes/db"
)

// 2b função responsável por gerar dados aleatoriamente
func GenerateRandomData() {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	var idPaciente int64
	sql := `SELECT id FROM pacientes ORDER BY RANDOM() LIMIT 1`

	err = conn.QueryRow(sql).Scan(&idPaciente)
	if err != nil {
		log.Printf("Erro ao selecionar um id_paciente aleatório: %v", err)
	}

	// random entre 20 e 130 bpm
	frequenciaCardiaca := rand.Intn(111) + 20

	// random entre 100 e 180
	pressaoSistolica := rand.Intn(80) + 100
	// random entre 70 e 110
	pressaoDiastolica := rand.Intn(40) + 70
	pressaoArterial := fmt.Sprintf("%d/%d", pressaoSistolica, pressaoDiastolica)

	// random entre 33 e 43 °C
	temperatura := 33.0 + rand.Float64()*(43.0-33.0)

	// random entre 80 e 99 %
	saturacaoOxigenio := 80.0 + rand.Float64()*(99.0-80.0)

	dataHoraRegistro := time.Now().Format("2006/01/02 15:04:05")

	_, err = conn.Exec(`INSERT INTO sinais_vitais (id_paciente, frequencia_cardiaca, pressao_arterial, temperatura, saturacao_oxigenio, data_hora_registro) VALUES ($1, $2, $3, $4, $5, $6)`,
		idPaciente, frequenciaCardiaca, pressaoArterial, temperatura, saturacaoOxigenio, dataHoraRegistro)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Dados inseridos com sucesso para o paciente de ID:", idPaciente)
}
