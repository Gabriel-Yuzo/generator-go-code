package main

import (
	"generator-go-code/internal/config"
	"generator-go-code/internal/database"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitConfig()
	database.Init()

	e := echo.New()

	log.Printf("Starting server on port %s...", config.AppConfig.Server.Port)
	if err := e.Start(":" + config.AppConfig.Server.Port); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
