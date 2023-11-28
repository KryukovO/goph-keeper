// Package repository содержит описание интерфейса репозитория сервера.
package repository

import (
	"context"

	"github.com/KryukovO/goph-keeper/internal/entities"
)

// UserRepository описывает интерфейс репозитория сервера
// для управления пользователями.
type UserRepository interface {
	// CreateUser выполняет запись пользователя user в репозиторий.
	CreateUser(ctx context.Context, user entities.User) (userID int64, err error)
	// User выполняет заполнение полей EncryptedPassword и Salt значениями из репозитория.
	User(ctx context.Context, user *entities.User) (err error)
}

// AuthDataRepository описывает интерфейс репозитория сервера
// для взаимодействия с парами логин/пароль пользователя.
type AuthDataRepository interface {
	// AddAuthData выполняет сохранение пары логин/пароль в репозитории.
	AddAuthData(ctx context.Context, data entities.AuthData) error
	// UpdateAuthData выполняет обновление пары логин/пароль в репозитории.
	UpdateAuthData(ctx context.Context, data entities.AuthData) error
	// DeleteAuthData выполняет удаление пары логин/пароль из репозитория.
	DeleteAuthData(ctx context.Context, data entities.AuthData) error
	// AuthDataList возвращает список сохраненных пар логин/пароль из репозитория.
	AuthDataList(ctx context.Context, userID int64) (data []entities.AuthData, err error)
}

// TextDataRepository описывает интерфейс репозитория сервера
// для взаимодействия с текстовыми пользователя.
type TextDataRepository interface {
	// AddTextData выполняет сохранение текстовых данных в репозитории.
	AddTextData(ctx context.Context, data entities.TextData) error
	// UpdateTextData выполняет обновление текстовых данных в репозитории.
	UpdateTextData(ctx context.Context, data entities.TextData) error
	// DeleteTextData выполняет удаление текстовых данных из репозитория.
	DeleteTextData(ctx context.Context, data entities.TextData) error
	// TextLabelsList возвращает список заголовков сохранённых текстовых данных из репозитория.
	TextLabelsList(ctx context.Context, userID int64) (labels []string, err error)
	// TextData возвращает сохранённые текстовые данные по заголовку из репозитория.
	TextData(ctx context.Context, data *entities.TextData) error
}

// BankDataRepository описывает интерфейс репозитория сервера
// для взаимодействия с данными банковских карт пользователя.
type BankDataRepository interface {
	// AddBankData выполняет сохранение данных банковских карт в репозитории.
	AddBankData(ctx context.Context, data entities.BankData) error
	// UpdateBankData выполняет обновление данных банковских карт в репозитории.
	UpdateBankData(ctx context.Context, data entities.BankData) error
	// DeleteBankData выполняет удаление данных банковских карт из репозитория.
	DeleteBankData(ctx context.Context, data entities.BankData) error
	// BankCardList возвращает список номеров банковских карт из репозитория.
	BankCardNumbersList(ctx context.Context, userID int64) (numbers []string, err error)
	// BankCardList возвращает данные банковской карты по номеру из репозитория.
	BankCard(ctx context.Context, data *entities.BankData) error
}

// SubscriptionRepository описывает интерфейс репозитория сервера
// для взаимодействия с подписками пользователя.
type SubscriptionRepository interface {
	// Sunbsriptions возвращает информацию о подписке пользователей из репозитория.
	Sunbsriptions() (subscriptions map[int64]entities.Subscription, err error)
}
