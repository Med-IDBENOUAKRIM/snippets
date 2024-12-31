package main

import (
	"context"
	"database/sql"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Application struct {
	logger *slog.Logger
}

func main() {
	godotenv.Load(".env")

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &Application{
		logger: logger,
	}

	log.Println("starting server on :4000")

	db, err := connectDB()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("db connection is established")

	err = http.ListenAndServe(":4000", app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func connectDB() (*sql.DB, error) {
	dbSource := os.Getenv("DB_SOURCE")
	driverName := os.Getenv("DRIVER_NAME")

	log.Println("dbSource >>> ", dbSource)
	log.Println("driverName >>> ", driverName)

	db, err := sql.Open(driverName, dbSource)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil

}
