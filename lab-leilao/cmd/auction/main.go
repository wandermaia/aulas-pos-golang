package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load() != nil {
		log.Fatal("Eror trying to load env variables")
		return
	}

}
