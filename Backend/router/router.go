package router

import (
	"fmt"
	"time"

	"github.com/MatheusOberziner/Monitoramento_Pacientes/configs"
	"github.com/MatheusOberziner/Monitoramento_Pacientes/handlers"
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

	// 2b A cada X tempo executará o evento
	// Iniciada gorotina
	go func() {
		// Define loop para que a cada 30 seg execute a função de geração de dados
		ticker := time.NewTicker(30 * time.Second)

		for range ticker.C {
			handlers.GenerateRandomData()
		}
	}()

	port := configs.GetServerPort()
	address := fmt.Sprintf(":%s", port)
	router.Logger.Fatal(router.Start(address))
}
