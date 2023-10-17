package router

import (
	"fmt"
	"log"
	"time"

	"github.com/MatheusOberziner/Monitoramento_Pacientes/configs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Initialize() {
	router := echo.New()

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: configs.GetCORS().AllowedOrigins,
		AllowMethods: configs.GetCORS().AllowedMethods,
	}))

	initialeRoutes(router)

	// Iniciada gorotina
	go func() {
		// Define loop para que a cada 20 seg execute a função de geração de dados
		ticker := time.NewTicker(20 * time.Second)

		for range ticker.C {
			log.Println("Loop executado a cada 20 segundos")
			// handlers.GenerateRandomData()
		}
	}()

	port := configs.GetServerPort()
	address := fmt.Sprintf(":%s", port)
	router.Logger.Fatal(router.Start(address))
}
