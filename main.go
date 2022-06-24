package main

import (
	"github.com/Tambarie/evolve-credit/application/server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env with godotenv: %s", err)
	}
	server.Start()
}
