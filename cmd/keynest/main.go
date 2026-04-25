package main

import (
	"log/slog"

	"github.com/light-devteam/keynest/internal/config"
	"github.com/light-devteam/keynest/internal/logger"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)
	log.Info(
		"Startup KeyNest Application",
		slog.String("env", cfg.Env),
	)
}
