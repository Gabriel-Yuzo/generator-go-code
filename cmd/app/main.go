package main

import (
	"generator-go-code/internal/config"
	"generator-go-code/internal/infra/database"
	"generator-go-code/internal/repository"
	"generator-go-code/internal/usecase"

	"generator-go-code/internal/server/handler"
	"generator-go-code/internal/server/routes"

	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitConfig()
	db := database.Init()
	e := echo.New()

	// Inicializa o repositório, serviço e handler
	userRepository := repository.NewUserRepositoryImpl(db)
	userUsecase := usecase.NewUserUsecaseImpl(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)
	routes.SetupUserRoutes(e, userHandler)

	log.Printf("Starting server on port %s...", config.AppConfig.Server.Port)
	if err := e.Start(":" + config.AppConfig.Server.Port); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
