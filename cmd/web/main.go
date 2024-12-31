package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
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

	err := http.ListenAndServe(":4000", app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
