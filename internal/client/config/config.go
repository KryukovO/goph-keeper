// Package config содержит конфигурацию клиента.
package config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	address        = "localhost:8080"
	requestTimeout = 15 * time.Second
	fileStorage    = "./"
)

// Config - конфигурация сервера.
type Config struct {
	// Address - адрес эндпоинта gRPC-сервера (host:port).
	Address string
	// RequestTimeout - ограничение времени выполнения запроса.
	RequestTimeout time.Duration
	// FileStorage - путь до директории в которую будут сохраняться файлы.
	FileStorage string
}

// NewConfig возвращает новый объект Config.
func NewConfig() *Config {
	vpr := viper.New()

	vpr.AllowEmptyEnv(false)

	vpr.BindEnv("server_address")
	vpr.BindEnv("req_timeout")
	vpr.BindEnv("file_storage")

	vpr.SetDefault("server_address", address)
	vpr.SetDefault("req_timeout", requestTimeout)
	vpr.SetDefault("file_storage", fileStorage)

	return &Config{
		Address:        vpr.GetString("server_address"),
		RequestTimeout: vpr.GetDuration("req_timeout"),
		FileStorage:    vpr.GetString("file_storage"),
	}
}
