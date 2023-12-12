package main

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"

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

	helpFlag := false

	pflag.BoolVarP(&helpFlag, "help", "h", false, "Shows usage")

	pflag.StringVarP(&cfg.Address, "address", "a", cfg.Address, "Address to run HTTP server")
	pflag.StringVarP(&cfg.DSN, "dsn", "d", cfg.DSN, "URI to database")
	pflag.StringVarP(&cfg.FSFolder, "fsfolder", "f", cfg.FSFolder, "File storage folder")
	pflag.StringVar(&cfg.SecretKey, "secret", cfg.SecretKey, "Authorization token encryption key")
	pflag.DurationVar(&cfg.UserTokenTTL, "tokenttl", cfg.UserTokenTTL, "User token lifetime")
	pflag.DurationVar(&cfg.RepositoryTimeout, "repotimeout", cfg.RepositoryTimeout, "Repository connection timeout")

	pflag.Parse()

	if helpFlag {
		pflag.Usage()

		return
	}

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
