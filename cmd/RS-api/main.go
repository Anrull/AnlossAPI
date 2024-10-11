package main

import (
	"AnlossAPI/internal/bot"
	"AnlossAPI/internal/config"
	http_server "AnlossAPI/internal/http-server"
	"AnlossAPI/internal/storage/sqlite"
	"AnlossAPI/pkg/scheduler"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
	envDev   = "dev"
)

func main() {
	cfg := config.MustLoad()

	logger := setupLogger(cfg.Env)

	logger.Info("starting server")

	sqlite.New(*cfg)

	bot.New(logger)

	go http_server.New(logger, cfg)
	go autoSendDataBases(cfg, logger)

	select {}
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(
			slog.NewTextHandler(
				os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		logger = slog.New(
			slog.NewJSONHandler(
				os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(
				os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}

func autoSendDataBases(cfg *config.Config, logger *slog.Logger) {
	for i := 0; i < 7; i++ {
		scheduler.NewScheduler(i, 15, 0, func() {
			err := bot.SendFile(cfg.RecordsPath, "records.db", "time")
			if err != nil {
				logger.Info("error sending file", "file", cfg.RecordsPath, "error", err)
			}

			err = bot.SendFile(cfg.StudentsPath, "students.db", "time")
			if err != nil {
				logger.Info("error sending file", "file", cfg.StudentsPath, "error", err)
			}
		})
	}
}
