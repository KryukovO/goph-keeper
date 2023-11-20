// Package repository содержит описание интерфейса репозитория сервера.
package repository

import (
	"context"

	"github.com/KryukovO/goph-keeper/internal/server/entities"
)

// Repository описывает репозиторий сервера.
type Repository interface {
	// CreateUser выполняет создание пользователя user.
	CreateUser(ctx context.Context, user entities.User) (userID int64, err error)
	// User выполняет заполнение полей EncryptedPassword и Salt значениями из репозитория.
	User(ctx context.Context, user *entities.User) (err error)
	// Close закрывает соединение с репозиторием.
	Close() error
}
