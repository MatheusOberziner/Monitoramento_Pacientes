package router

import (
	"github.com/MatheusOberziner/Monitoramento_Pacientes/handlers"
	"github.com/labstack/echo/v4"
)

func initialeRoutes(router *echo.Echo) {
	basePath := "/api"
	pacientes := router.Group(basePath + "/pacientes")
	{
		pacientes.GET("", handlers.ListPacientes)
		// pacientes.GET("/:id", handlers.ShowPaciente)
		// pacientes.POST("", handlers.CreatePaciente)
	}
}
