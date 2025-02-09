package main

import (
	"fmt"
	"github.com/Miroslovelife/whareflow/internal/config"
	"github.com/Miroslovelife/whareflow/pkg/database"
	"github.com/Miroslovelife/whareflow/pkg/server"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	slog.Info("Success loading config")

	log := setupLogger(cfg.Env)

	log.Info("Success setup logger")

	db := database.NewPostgres(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.StoragePath.Postgres.Host,
		"postgres",
		cfg.StoragePath.Postgres.Password,
		cfg.StoragePath.Postgres.Database,
		cfg.StoragePath.Postgres.Port,
		"disable",
	))

	echoServer := server.NewEchoServer(*log, db, cfg)
	echoServer.Start()

}

func setupLogger(env string) *slog.Logger {
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
