package handlers

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/MatheusOberziner/Monitoramento_Pacientes/db"
)

var mensagemNotificacao string

// 2b função responsável por gerar dados aleatoriamente
func GenerateRandomData() {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	var idPaciente int64
	var nomePaciente string
	sql := `SELECT id, nome FROM pacientes ORDER BY RANDOM() LIMIT 1`

	err = conn.QueryRow(sql).Scan(&idPaciente, &nomePaciente)
	if err != nil {
		log.Printf("Erro ao selecionar um id_paciente aleatório: %v", err)
	}

	mensagemNotificacao = ""
	mensagemNotificacao += nomePaciente + " - "
	// random entre 20 e 130 bpm
	frequenciaCardiaca := rand.Intn(111) + 20
	if frequenciaCardiaca < 60 || frequenciaCardiaca > 100 {
		mensagemNotificacao += "Frequência Cardíaca anormal / "
	}

	// random entre 100 e 180
	pressaoSistolica := rand.Intn(80) + 100
	// random entre 70 e 110
	pressaoDiastolica := rand.Intn(40) + 70
	pressaoArterial := fmt.Sprintf("%d/%d", pressaoSistolica, pressaoDiastolica)
	if pressaoSistolica < 139 || pressaoDiastolica > 89 {
		mensagemNotificacao += "Pressão Arterial elevada / "
	}

	// random entre 33 e 43 °C
	temperatura := 33.0 + rand.Float64()*(43.0-33.0)
	if temperatura < 36 {
		mensagemNotificacao += "Temperatura corporal baixa / "
	} else if temperatura > 37.5 {
		mensagemNotificacao += "Temperatura corporal alta / "
	}

	// random entre 84 e 99 %
	saturacaoOxigenio := 84.0 + rand.Float64()*(99.0-84.0)
	if saturacaoOxigenio < 95 {
		mensagemNotificacao += "Saturação abaixo / "
	}

	dataHoraRegistro := time.Now().Format("2006/01/02 15:04:05")

	_, err = conn.Exec(`INSERT INTO sinais_vitais (id_paciente, frequencia_cardiaca, pressao_arterial, temperatura, saturacao_oxigenio, data_hora_registro) VALUES ($1, $2, $3, $4, $5, $6)`,
		idPaciente, frequenciaCardiaca, pressaoArterial, temperatura, saturacaoOxigenio, dataHoraRegistro)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Nome: %s", nomePaciente)
	log.Println("Dados inseridos com sucesso para o paciente de ID:", idPaciente)
}

func GetMensagem() string {
	return mensagemNotificacao
}
