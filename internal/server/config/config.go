// Package config содержит конфигурацию сервера.
package config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	address           = ":8080"
	dsn               = ""
	secretKey         = ""
	userTokenTTL      = 30 * time.Minute
	shutdownTimeout   = 10 * time.Second
	repositoryTimeout = 3 * time.Second
)

// Config - конфигурация сервера.
type Config struct {
	// Address - адрес эндпоинта gRPC-сервера (host:port).
	Address string
	// DSN - адрес подключения к БД.
	DSN string
	// RepositoryTimeout - таймаут выполнения операций с репозиторием.
	RepositoryTimeout time.Duration
	// SecretKey - ключ шифрования токена аутентификации.
	SecretKey string
	// UserTokenTTL - время жизни токена пользователя
	UserTokenTTL time.Duration
}

// NewConfig возвращает новый объект Config.
func NewConfig() *Config {
	vpr := viper.New()

	vpr.AllowEmptyEnv(false)

	vpr.BindEnv("server_address")
	vpr.BindEnv("database_uri")
	vpr.BindEnv("jwt_secret")
	vpr.BindEnv("jwt_ttl")
	vpr.BindEnv("server_shutdown")
	vpr.BindEnv("repository_timeout")

	vpr.SetDefault("server_address", address)
	vpr.SetDefault("database_uri", dsn)
	vpr.SetDefault("jwt_secret", secretKey)
	vpr.SetDefault("jwt_ttl", userTokenTTL)
	vpr.SetDefault("server_shutdown", shutdownTimeout)
	vpr.SetDefault("repository_timeout", repositoryTimeout)

	return &Config{
		Address:      vpr.GetString("server_address"),
		DSN:          vpr.GetString("database_uri"),
		SecretKey:    vpr.GetString("jwt_secret"),
		UserTokenTTL: vpr.GetDuration("jwt_ttl"),
	}
}
