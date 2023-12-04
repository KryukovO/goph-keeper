// Package usecases описывает слой логики взамодействия с репозиторием.
package usecases

import (
	"context"

	"github.com/KryukovO/goph-keeper/internal/entities"
)

// User описывает интерфейс логики взамодействия
// с репозиторием для управления пользователями.
type User interface {
	// Registration выполняет регистрацию пользователя.
	Registration(ctx context.Context, user entities.User) (token string, err error)
	// Authorization выполняет авторизацию пользователя.
	Authorization(ctx context.Context, user entities.User) (token string, err error)
}

// AuthData описывает интерфейс логики взамодействия
// с репозиторием для управления парами логин/пароль пользователя.
type AuthData interface {
	// AddAuthData выполняет сохранение пары логин/пароль.
	AddAuthData(ctx context.Context, data entities.AuthData) error
	// UpdateAuthData выполняет обновление пары логин/пароль.
	UpdateAuthData(ctx context.Context, oldResource, oldLogin string, data entities.AuthData) error
	// DeleteAuthData выполняет удаление пары логин/пароль.
	DeleteAuthData(ctx context.Context, data entities.AuthData) error
	// AuthDataList возвращает список сохраненных пар логин/пароль.
	AuthDataList(ctx context.Context, userID int64) (data []entities.AuthData, err error)
}

// TextData описывает интерфейс логики взамодействия
// с репозиторием для управления текстовыми данными пользователя.
type TextData interface {
	// AddTextData выполняет сохранение текстовых данных.
	AddTextData(ctx context.Context, data entities.TextData) error
	// UpdateTextData выполняет обновление текстовых данных.
	UpdateTextData(ctx context.Context, oldLabel string, data entities.TextData) error
	// DeleteTextData выполняет удаление текстовых данных.
	DeleteTextData(ctx context.Context, data entities.TextData) error
	// TextLabelsList возвращает список заголовков сохранённых текстовых данных.
	TextLabelsList(ctx context.Context, userID int64) (labels []string, err error)
	// TextData возвращает сохранённые текстовые данные по заголовку.
	TextData(ctx context.Context, data *entities.TextData) error
}

// BankData описывает интерфейс логики взамодействия
// с репозиторием для управления данными банковских карт пользователя.
type BankData interface {
	// AddBankData выполняет сохранение данных банковских карт.
	AddBankData(ctx context.Context, data entities.BankData) error
	// UpdateBankData выполняет обновление данных банковских карт.
	UpdateBankData(ctx context.Context, oldNumber string, data entities.BankData) error
	// DeleteBankData выполняет удаление данных банковских карт.
	DeleteBankData(ctx context.Context, data entities.BankData) error
	// BankCardList возвращает список номеров банковских карт.
	BankCardNumbersList(ctx context.Context, userID int64) (numbers []string, err error)
	// BankCard возвращает данные банковской карты по номеру.
	BankCard(ctx context.Context, data *entities.BankData) error
}

// BinaryData описывает интерфейс логики взамодействия
// с репозиторием для управления бинарными данными пользователя.
type BinaryData interface {
	// AddBinaryData выполняет сохранение бинарных данных.
	AddBinaryData(ctx context.Context, data entities.File) error
	// DeleteBinaryData выполняет удаление бинарных данных.
	DeleteBinaryData(ctx context.Context, data entities.File) error
	// FileNamesList возвращает список сохранённых файлов.
	FileNamesList(ctx context.Context, userID int64) (fileNames []string)
	// BinaryData возвращает сохранённые бинарные данные по имени файла.
	BinaryData(ctx context.Context, data *entities.File) error
	// UpdateSubscription обновляет информацию о подписке пользователя.
	UpdateSubscription(ctx context.Context, userID int64, subscription entities.Subscription)
}
