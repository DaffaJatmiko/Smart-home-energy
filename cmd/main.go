package main

import (
	"log"
	"os"

	"github.com/DaffaJatmiko/smart-home-energy/handlers"
	"github.com/DaffaJatmiko/smart-home-energy/service"
	"github.com/DaffaJatmiko/smart-home-energy/util"
	"github.com/gin-gonic/gin"
)

func main() {
	// Check current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	log.Printf("Current working directory: %s", wd)

	// Load environment variables
	err = util.LoadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	token := os.Getenv("HUGGINGFACE_TOKEN")
	if token == "" {
		log.Fatal("HUGGINGFACE_TOKEN is not set")
	}

	router := gin.Default()

	aiservice := service.NewAIService()
	aiHandler := handlers.NewAIHandler(aiservice)

	router.StaticFile("/", "./view/index.html")

	router.POST("/get-answer", aiHandler.GetAnswer)

	port := ":8080"
	log.Printf("Listening on port %s", port)
	err =router.Run(port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
