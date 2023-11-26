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

// NewPgRepo создаёт новый объект PgRepo.
func NewPgRepo(db *postgres.Postgres) *PgRepo {
	return &PgRepo{
		db: db,
	}
}

// CreateUser выполняет создание пользователя user.
func (repo *PgRepo) CreateUser(
	ctx context.Context, user entities.User, subscription entities.Subscription,
) (int64, error) {
	return 0, nil
}

// User выполняет заполнение полей EncryptedPassword и Salt значениями из репозитория.
func (repo *PgRepo) User(ctx context.Context, user *entities.User) error {
	return nil
}

// AddAuthData выполняет сохранение пары логин/пароль в репозитории.
func (repo *PgRepo) AddAuthData(ctx context.Context, data entities.AuthData) error {
	return nil
}

// AddTextData выполняет сохранение текстовых данных в репозитории.
func (repo *PgRepo) AddTextData(ctx context.Context, data entities.TextData) error {
	return nil
}

// AddBankData выполняет сохранение данных банковских карт в репозитории.
func (repo *PgRepo) AddBankData(ctx context.Context, data entities.BankData) error {
	return nil
}

// UpdateAuthData выполняет обновление пары логин/пароль в репозитории.
func (repo *PgRepo) UpdateAuthData(ctx context.Context, data entities.AuthData) error {
	return nil
}

// UpdateTextData выполняет обновление текстовых данных в репозитории.
func (repo *PgRepo) UpdateTextData(ctx context.Context, data entities.TextData) error {
	return nil
}

// UpdateBankData выполняет обновление данных банковских карт в репозитории.
func (repo *PgRepo) UpdateBankData(ctx context.Context, data entities.BankData) error {
	return nil
}

// DeleteAuthData выполняет удаление пары логин/пароль из репозитория.
func (repo *PgRepo) DeleteAuthData(ctx context.Context, data entities.AuthData) error {
	return nil
}

// DeleteTextData выполняет удаление текстовых данных из репозитория.
func (repo *PgRepo) DeleteTextData(ctx context.Context, data entities.TextData) error {
	return nil
}

// DeleteBankData выполняет удаление данных банковских карт из репозитория.
func (repo *PgRepo) DeleteBankData(ctx context.Context, data entities.BankData) error {
	return nil
}

// AuthDataList возвращает список сохраненных пар логин/пароль из репозитория.
func (repo *PgRepo) AuthDataList(ctx context.Context, userID int64) ([]entities.AuthData, error) {
	return []entities.AuthData{}, nil
}

// TextLabelsList возвращает список заголовков сохранённых текстовых данных из репозитория.
func (repo *PgRepo) TextLabelsList(ctx context.Context, userID int64) ([]string, error) {
	return []string{}, nil
}

// TextData возвращает сохранённые текстовые данные по заголовку из репозитория.
func (repo *PgRepo) TextData(ctx context.Context, data *entities.TextData) error {
	return nil
}

// FileNamesList возвращает список сохранённых файлов из репозитория.
func (repo *PgRepo) FileNamesList(ctx context.Context, userID int64) ([]string, error) {
	return []string{}, nil
}

// BankCardList возвращает список номеров банковских карт из репозитория.
func (repo *PgRepo) BankCardNumbersList(ctx context.Context, userID int64) ([]string, error) {
	return []string{}, nil
}

// BankCardList возвращает данные банковской карты по номеру из репозитория.
func (repo *PgRepo) BankCard(ctx context.Context, data *entities.BankData) error {
	return nil
}

// Sunbsriptions возвращает информацию о подписке пользователей из репозитория.
func (repo *PgRepo) Sunbsriptions() (map[int64]entities.Subscription, error) {
	return map[int64]entities.Subscription{}, nil
}
