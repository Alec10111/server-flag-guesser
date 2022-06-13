package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
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

type application struct {
	config config
	logger *log.Logger
}

func main() {
	srv := &http.Server{
		Addr:         fmt.Sprint(":3001"),
		Handler:      router(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	logger.Println("Starting server on port 3001")

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
