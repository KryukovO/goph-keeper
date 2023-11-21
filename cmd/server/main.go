package main

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sirupsen/logrus"

	"github.com/KryukovO/goph-keeper/internal/server/app"
	"github.com/KryukovO/goph-keeper/internal/server/config"
)

var (
	// buildVersion представляет собой хранилище значения ldflag - версия сборки.
	buildVersion = "N/A"
	// buildDate представляет собой хранилище значения ldflag - дата сборки.
	buildDate = "N/A"
)

func main() {
	fmt.Printf(
		"Build version: %s\nBuild date: %s\n",
		buildVersion, buildDate,
	)

	cfg := config.NewConfig()

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05 Z07:00",
	})

	srv := app.NewApp(cfg, logger)
	if err := srv.Run(context.Background()); err != nil {
		logger.Fatalf("Server error: %s. Exit(1)", err.Error())
	}
}
