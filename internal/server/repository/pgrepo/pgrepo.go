// Package pgstorage описывает репозиторий сервера, реализованный в PostgreSQL.
package pgrepo

import (
	"context"
	"database/sql"

	"github.com/KryukovO/goph-keeper/internal/server/entities"
)

// PgRepo - репозиторий PostgreSQL.
type PgRepo struct {
	db *sql.DB
}

// NewPgRepo - создаёт новое подключение к репозиторию PostgreSQL.
func NewPgRepo(ctx context.Context, dsn string) (*PgRepo, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	s := &PgRepo{
		db: db,
	}

	err = s.db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// CreateUser выполняет создание пользователя user.
func (repo *PgRepo) CreateUser(ctx context.Context, user entities.User) (int64, error) {
	return 0, nil
}

// User выполняет заполнение полей EncryptedPassword и Salt значениями из репозитория.
func (repo *PgRepo) User(ctx context.Context, user *entities.User) error {
	return nil
}

// Close закрывает соединение с репозиторием.
func (repo *PgRepo) Close() error {
	return repo.db.Close()
}
