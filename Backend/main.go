package main

import (
	"github.com/MatheusOberziner/Monitoramento_Pacientes/configs"
	"github.com/MatheusOberziner/Monitoramento_Pacientes/router"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	router.Initialize()
}
