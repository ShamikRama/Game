package main

import (
	"Game/internal/config"
	"Game/internal/db"
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustLoad()
	db, err := db.New(*cfg)
	log := setUpLogger()
	if err != nil {
		log.Error("failed to connect storage", "error", err)
		os.Exit(1)
	}

}

func setUpLogger() *slog.Logger {
	var log *slog.Logger
	log = slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
	)
	return log
}
