package main

import (
	"fmt"
	"log/slog"
	"os"
	"restfulApiGo/internal/config"
)

const (
	envLocal="local"
	envDev="dev"

)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting prod", slog.String("env", cfg.Env))
	log.Debug("debug messegaes are enable")

	fmt.Println(cfg)
}

func setupLogger(env string) *slog.Logger { 

	var log *slog.Logger

	switch env{
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}