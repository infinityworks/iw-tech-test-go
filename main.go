package main

import (
	"log/slog"
	"os"

	"github.com/aviva-verde/tech-test-backend-go/api"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	server := api.Server{
		Logger: logger,
	}
	server.Start()
}
