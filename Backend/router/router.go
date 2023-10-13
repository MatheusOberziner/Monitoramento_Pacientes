package router

import (
	"fmt"

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

	port := configs.GetServerPort()
	address := fmt.Sprintf(":%s", port)
	router.Logger.Fatal(router.Start(address))
}
