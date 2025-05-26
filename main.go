package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	var v = godotenv.Load(".env")
	println("Loading environment variables from .env file:", v)

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable is not set")

	}
	fmt.Printf("The PORT is set to: %s\n", portString)
}
