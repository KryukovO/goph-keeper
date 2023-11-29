package main

import (
	"fmt"

	"github.com/KryukovO/goph-keeper/internal/client/app"
	"github.com/KryukovO/goph-keeper/internal/client/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

var (
	// buildVersion представляет собой хранилище значения ldflag - версия сборки.
	buildVersion = "N/A"
	// buildDate представляет собой хранилище значения ldflag - дата сборки.
	buildDate = "N/A"
)

func main() {
	cfg := config.NewConfig()

	helpFlag := false
	versionFlag := false

	pflag.BoolVarP(&helpFlag, "help", "h", false, "Shows usage")
	pflag.BoolVarP(&versionFlag, "version", "v", false, "Shows version info")

	pflag.StringVarP(&cfg.Address, "address", "a", cfg.Address, "Address to run HTTP server")
	pflag.DurationVarP(&cfg.RequestTimeout, "timeout", "t", cfg.RequestTimeout, "Request timeout")
	pflag.StringVarP(&cfg.FileStorage, "files", "f", cfg.FileStorage, "File storage directory")

	pflag.Parse()

	if helpFlag {
		pflag.Usage()

		return
	}

	if versionFlag {
		fmt.Printf(
			"Build version: %s\nBuild date: %s\n",
			buildVersion, buildDate,
		)

		return
	}

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05 Z07:00",
	})

	a := app.NewApp(cfg)
	if err := a.Run(); err != nil {
		logger.Fatalf("Client error: %s. Exit(1)", err.Error())
	}
}
