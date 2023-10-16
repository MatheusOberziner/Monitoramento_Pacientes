package router

import (
	"github.com/MatheusOberziner/Monitoramento_Pacientes/handlers"
	"github.com/labstack/echo/v4"
)

// Endpoints que serão consumidos no Frontend
func initialeRoutes(router *echo.Echo) {
	basePath := "/api"
	pacientes := router.Group(basePath + "/pacientes")
	{
		pacientes.GET("", handlers.ListPacientes)
		pacientes.GET("/:id", handlers.GetPaciente)
		pacientes.POST("", handlers.CreatePaciente)
	}
	sinais_vitais := router.Group(basePath + "/sinais_vitais")
	{
		sinais_vitais.GET("/:id_paciente", handlers.ListSinaisVitais)
	}
}
