// package database

// import (
// 	"database/sql"
// 	// "fmt"
// 	"fmt"
// 	models "geobase/internal/models"
// 	_ "github.com/lib/pq" // required for PostgreSQL connection
// 	"log"
// )

// // const parameter database
// const (
// 	HOST = "host"
// 	PORT = "5432"
// )

// // Database connection
// type Database struct {
// 	db *sql.DB
// }

// //Initialize database connection
// func Initialize(config *models.Config) *Database {
// 	address := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
// 	db, err := sql.Open("postgres", address)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = db.Ping(); err != nil {
// 		log.Fatal(err)
// 	}

// 	pgStorage := Database{db: db}
// 	return &pgStorage
// }
