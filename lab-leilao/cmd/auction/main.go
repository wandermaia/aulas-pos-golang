package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/wandermaia/aulas-pos-golang/lab-leilao/configuration/database/mongodb"
)

func main() {

	ctx := context.Background()
	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Eror trying to load env variables")
		return
	}

	_, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

}
