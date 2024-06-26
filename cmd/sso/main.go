package main

import (

	"SSO/internal/config"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev = "dev"
	envProd = " prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting application",
		slog.String("env", cfg.Env),
		slog.Any("cfg", cfg),
		slog.Int("port", cfg.GRPC.Port),
		)

	log.Debug("debug message")

	log.Warn("warn message")

	log.Error("error message")

	//TODO: инициализировать логгер

	//TODO: инициализировать приложение (app) - точку входа в основной сервис

	//TODO: запустить gRPC-сервер приложения
}

func setupLogger(env string) *slog.Logger{
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
			)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
			)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
			)
	}

	return log
}