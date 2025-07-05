package main

import (
	"log"
	"os"
	"path/filepath"
	"restservice/internal/exceptions"
	"restservice/internal/server"
	"strconv"

	"github.com/joho/godotenv"

	"restservice/internal/infra"
)

func loadEnv() {
	exePath, _ := os.Executable()
	dir := filepath.Dir(exePath)

	candidates := []string{
		filepath.Join(dir, ".env"),
		filepath.Join(dir, "cmd", ".env"),
		filepath.Join(dir, "..", ".env"),
		filepath.Join(dir, "..", "cmd", ".env"),
	}

	for _, p := range candidates {
		if err := godotenv.Load(p); err == nil {
			log.Printf("✅ .env loaded from: %s", p)
			return
		}
	}

	exceptions.HandleError(&exceptions.CustomException{
		Field:   "DotEnv",
		Message: "Failed to load env file",
	})
	log.Fatal("Failed to load .env file")
}

func main() {
	loadEnv()

	port := os.Getenv("SERVER_PORT")
	if _, err := strconv.Atoi(port); err != nil {
		exceptions.HandleError(&exceptions.CustomException{
			Field:   "DotEnv",
			Message: "Incorrect type of SERVER_PORT",
		})
	}

	go func() {
		if err := <-infra.StartPrometheus(":9090"); err != nil {
			log.Println(err)
		}
	}()

	dbLink := os.Getenv("DB_LINK")
	log.Printf("Running server on port %s", port)
	server.Setup(port, dbLink)
}
