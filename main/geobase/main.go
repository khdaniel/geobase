package main

import (
	"../internal/database"
	"../internal/server"
	"log"
)

const address = ":1234"

func main() {
	storage := database.NewRuntimeStorage()
	srv := server.New(storage)
	log.Println("starting server")
	err := srv.Run(address)
	if err != nil {
		log.Fatal(err)
	}
}
