// Package repository содержит описание интерфейса репозитория сервера.
package repository

import (
	"context"

	"github.com/KryukovO/goph-keeper/internal/server/entities"
)

// Repository описывает репозиторий сервера.
type Repository interface {
	// CreateUser выполняет запись пользователя user в репозиторий.
	CreateUser(ctx context.Context, user entities.User) (userID int64, err error)
	// User выполняет заполнение полей EncryptedPassword и Salt значениями из репозитория.
	User(ctx context.Context, user *entities.User) (err error)

	// AddAuthData выполняет сохранение пары логин/пароль в репозитории.
	AddAuthData(ctx context.Context, data entities.AuthData) error
	// AddTextData выполняет сохранение текстовых данных в репозитории.
	AddTextData(ctx context.Context, data entities.TextData) error
	// AddBinaryData выполняет сохранение бинарных данных в репозитории.
	AddBinaryData(ctx context.Context, data entities.File) error
	// AddBankData выполняет сохранение данных банковских карт в репозитории.
	AddBankData(ctx context.Context, data entities.BankData) error

	// UpdateAuthData выполняет обновление пары логин/пароль в репозитории.
	UpdateAuthData(ctx context.Context, data entities.AuthData) error
	// UpdateTextData выполняет обновление текстовых данных в репозитории.
	UpdateTextData(ctx context.Context, data entities.TextData) error
	// UpdateBinaryData выполняет обновление бинарных данных в репозитории.
	UpdateBinaryData(ctx context.Context, data entities.File) error
	// UpdateBankData выполняет обновление данных банковских карт в репозитории.
	UpdateBankData(ctx context.Context, data entities.BankData) error

	// DeleteAuthData выполняет удаление пары логин/пароль из репозитория.
	DeleteAuthData(ctx context.Context, data entities.AuthData) error
	// DeleteTextData выполняет удаление текстовых данных из репозитория.
	DeleteTextData(ctx context.Context, data entities.TextData) error
	// DeleteBinaryData выполняет удаление бинарных данных из репозитория.
	DeleteBinaryData(ctx context.Context, data entities.File) error
	// DeleteBankData выполняет удаление данных банковских карт из репозитория.
	DeleteBankData(ctx context.Context, data entities.BankData) error

	// AuthDataList возвращает список сохраненных пар логин/пароль из репозитория.
	AuthDataList(ctx context.Context, userID int64) (data []entities.AuthData, err error)
	// TextLabelsList возвращает список заголовков сохранённых текстовых данных из репозитория.
	TextLabelsList(ctx context.Context, userID int64) (labels []string, err error)
	// TextData возвращает сохранённые текстовые данные по заголовку из репозитория.
	TextData(ctx context.Context, data *entities.TextData) error
	// FileNamesList возвращает список сохранённых файлов из репозитория.
	FileNamesList(ctx context.Context, userID int64) (fileNames []string, err error)
	// BinaryData возвращает сохранённые бинарные данные по имени файла из репозитория.
	BinaryData(ctx context.Context, data *entities.File) error
	// BankCardList возвращает список номеров банковских карт из репозитория.
	BankCardNumbersList(ctx context.Context, userID int64) (numbers []string, err error)
	// BankCardList возвращает данные банковской карты по номеру из репозитория.
	BankCard(ctx context.Context, data *entities.BankData) error
}
