// Package pgstorage описывает репозиторий сервера, реализованный в PostgreSQL.
package pgrepo

import (
	"context"

	"github.com/KryukovO/goph-keeper/internal/server/entities"
	"github.com/KryukovO/goph-keeper/pkg/postgres"
)

// PgRepo - репозиторий PostgreSQL.
type PgRepo struct {
	db *postgres.Postgres
}

// NewPgRepo создаёт новый новый объект PgRepo.
func NewPgRepo(db *postgres.Postgres) *PgRepo {
	return &PgRepo{
		db: db,
	}
}

// CreateUser выполняет создание пользователя user.
func (repo *PgRepo) CreateUser(ctx context.Context, user entities.User) (int64, error) {
	return 0, nil
}

// User выполняет заполнение полей EncryptedPassword и Salt значениями из репозитория.
func (repo *PgRepo) User(ctx context.Context, user *entities.User) error {
	return nil
}
