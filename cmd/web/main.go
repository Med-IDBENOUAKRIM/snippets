package main

import (
	"context"
	"database/sql"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Application struct {
	logger *slog.Logger
}

func main() {

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

	db, err := sql.Open("postgres", dbSource)

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
