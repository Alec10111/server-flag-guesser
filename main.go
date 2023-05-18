package main

import (
	"fgapi/src/models"
	"fgapi/src/routes"
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

// type application struct {
// 	config config
// 	logger *log.Logger
// }

func main() {

	dsn := "host=flag-db user=flag password=flag dbname=flag port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Country{})

	srv := &http.Server{
		// Addr:         fmt.Sprint(":" + os.Getenv("PORT")),
		Addr:         ":8080",
		Handler:      routes.Router(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	logger.Println("Starting server on port " + os.Getenv("PORT"))

	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
