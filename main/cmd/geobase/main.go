package main

import (
	"flag"
	"fmt"
	"geobase/internal/database"
	models "geobase/internal/models"
	"geobase/internal/server"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

const address = ":1234"

func main() {
	cfg := prepareConfig()
	storage := database.Initialize(cfg)
	srv := server.New(storage)
	log.Println("starting server")
	err := srv.Run(address)
	if err != nil {
		log.Fatal(err)
	}
}

func prepareConfig() *models.Config {
	var cfg models.Config
	configFile := getConfigFile()

	if err := cleanenv.ReadConfig(configFile, &cfg); err != nil {
		fmt.Printf("Unable to get app configuration due to: %s\n", err.Error())
	}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		fmt.Printf("Unable to retrieve app configuration due to: %s\n", err.Error())
		os.Exit(1)
	}
	return &cfg
}

func getConfigFile() string {
	configFile := flag.String("config", "config.yml", "config file")
	return *configFile
}
