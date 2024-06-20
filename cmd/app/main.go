package main

import (
	"generator-go-code/internal/config"
	"generator-go-code/internal/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.InitConfig()
	database.Init()

	r := mux.NewRouter()
	// Adicione suas rotas aqui

	log.Printf("Starting server on port %s...", config.AppConfig.Server.Port)
	if err := http.ListenAndServe(":"+config.AppConfig.Server.Port, r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
