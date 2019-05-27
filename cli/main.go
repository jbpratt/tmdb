package main

import (
	"log"

	"github.com/jbpratt78/tmdb/cli/cmd"
	"github.com/joho/godotenv"
)

func main() {
	cmd.Execute()
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
