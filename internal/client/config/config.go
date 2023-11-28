// Package config содержит конфигурацию клиента.
package config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	address        = "localhost:8080"
	requestTimeout = 5 * time.Second
)

// Config - конфигурация сервера.
type Config struct {
	// Address - адрес эндпоинта gRPC-сервера (host:port).
	Address string
	// RequestTimeout - ограничение времени выполнения запроса.
	RequestTimeout time.Duration
}

// NewConfig возвращает новый объект Config.
func NewConfig() *Config {
	vpr := viper.New()

	vpr.AllowEmptyEnv(false)

	vpr.BindEnv("server_address")
	vpr.BindEnv("req_timeout")

	vpr.SetDefault("server_address", address)
	vpr.SetDefault("req_timeout", requestTimeout)

	return &Config{
		Address:        vpr.GetString("server_address"),
		RequestTimeout: vpr.GetDuration("req_timeout"),
	}
}
