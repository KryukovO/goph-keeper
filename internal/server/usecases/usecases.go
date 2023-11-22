// Package usecases описывает слой логики взамодействия с репозиторием.
package usecases

import (
	"context"

	"github.com/KryukovO/goph-keeper/internal/server/entities"
)

// UseCases описывает интерфейс логики взамодействия с репозиторием.
type UseCases interface {
	// Registration выполняет регистрацию пользователя.
	Registration(ctx context.Context, user entities.User) (token string, err error)
	// Authorization выполняет авторизацию пользователя.
	Authorization(ctx context.Context, user entities.User) (token string, err error)

	// AddAuthData выполняет сохранение пары логин/пароль.
	AddAuthData(ctx context.Context, data entities.AuthData) error
	// AddTextData выполняет сохранение текстовых данных.
	AddTextData(ctx context.Context, data entities.TextData) error
	// AddBinaryData выполняет сохранение бинарных данных.
	AddBinaryData(ctx context.Context, data entities.File) error
	// AddBankData выполняет сохранение данных банковских карт.
	AddBankData(ctx context.Context, data entities.BankData) error

	// UpdateAuthData выполняет обновление пары логин/пароль.
	UpdateAuthData(ctx context.Context, data entities.AuthData) error
	// UpdateTextData выполняет обновление текстовых данных.
	UpdateTextData(ctx context.Context, data entities.TextData) error
	// UpdateBinaryData выполняет обновление бинарных данных.
	UpdateBinaryData(ctx context.Context, data entities.File) error
	// UpdateBankData выполняет обновление данных банковских карт.
	UpdateBankData(ctx context.Context, data entities.BankData) error

	// DeleteAuthData выполняет удаление пары логин/пароль.
	DeleteAuthData(ctx context.Context, data entities.AuthData) error
	// DeleteTextData выполняет удаление текстовых данных.
	DeleteTextData(ctx context.Context, data entities.TextData) error
	// DeleteBinaryData выполняет удаление бинарных данных.
	DeleteBinaryData(ctx context.Context, data entities.File) error
	// DeleteBankData выполняет удаление данных банковских карт.
	DeleteBankData(ctx context.Context, data entities.BankData) error

	// AuthDataList возвращает список сохраненных пар логин/пароль.
	AuthDataList(ctx context.Context, userID int64) (data []entities.AuthData, err error)
	// TextLabelsList возвращает список заголовков сохранённых текстовых данных.
	TextLabelsList(ctx context.Context, userID int64) (labels []string, err error)
	// TextData возвращает сохранённые текстовые данные по заголовку.
	TextData(ctx context.Context, data *entities.TextData) error
	// FileNamesList возвращает список сохранённых файлов.
	FileNamesList(ctx context.Context, userID int64) (fileNames []string, err error)
	// BinaryData возвращает сохранённые бинарные данные по имени файла.
	BinaryData(ctx context.Context, data *entities.File) error
	// BankCardList возвращает список номеров банковских карт.
	BankCardNumbersList(ctx context.Context, userID int64) (numbers []string, err error)
	// BankCardList возвращает данные банковской карты по номеру.
	BankCard(ctx context.Context, data *entities.BankData) error
}
