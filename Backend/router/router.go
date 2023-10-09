package router

import (
	"fmt"

	"github.com/MatheusOberziner/Monitoramento_Pacientes/configs"
	"github.com/labstack/echo/v4"
)

func Initialize() {
	router := echo.New()

	initialeRoutes(router)

	port := configs.GetServerPort()
	address := fmt.Sprintf(":%s", port)
	router.Logger.Fatal(router.Start(address))
}
