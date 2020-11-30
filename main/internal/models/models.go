package models

import (
// "fmt"
// "net/http"
)

// Config of Database
type Config struct {
	AppPort    string `yaml:"app_port" env:"PORT"`
	DBPort     string `yaml:"db_port" env:"DBPORT" env-default:"5432"`
	DBHost     string `yaml:"db_host" env:"DBHOST" env-default:"localhost"`
	DBName     string `yaml:"db_name" env:"DBNAME" env-default:"postgres"`
	DBUser     string `yaml:"db_user" env:"DBUSER" env-default:"postgres"`
	DBPassword string `yaml:"db_password" env:"DBPASSWORD"`
}

// PointLocation data
type PointLocation struct {
	pointID    int
	latitude   float64
	longitude  float64
	title      string
	address    string
	wasteTypes string
}
