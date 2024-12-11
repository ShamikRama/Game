package main

import (
	"Game/internal/api"
	"Game/internal/config"
	"Game/internal/db"
	"Game/internal/repository"
	"Game/internal/service"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	db, err := db.New(*cfg)
	log := setUpLogger()
	if err != nil {
		log.Error("failed to connect storage", "error", err)
		os.Exit(1)
	}

	repos := repository.NewRepository(db)
	serv := service.NewService(repos)
	handlers := api.NewHandlers(serv)

	router := handlers.InitRoutes()

	srv := http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	go func() {
		err = srv.ListenAndServe()
		if err != nil {
			log.Error("failde to start the server")
		}
	}()

	log.Error("server running")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Error("server stoped")

}

func setUpLogger() *slog.Logger {
	var log *slog.Logger
	log = slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
	)
	return log
}
