package usecases

import (
	"context"
	"time"

	"github.com/KryukovO/goph-keeper/internal/server/entities"
	"github.com/KryukovO/goph-keeper/internal/server/repository"
)

// KeeperUseCases реализует логику взаимодействия с репозиторием.
type KeeperUseCases struct {
	repo    repository.Repository
	timeout time.Duration
}

// NewKeeperUseCases возвращает новый объект KeeperUseCases.
func NewKeeperUseCases(repo repository.Repository, timeout time.Duration) *KeeperUseCases {
	return &KeeperUseCases{
		repo:    repo,
		timeout: timeout,
	}
}

// Registration выполняет регистрацию пользователя.
func (uc *KeeperUseCases) Registration(ctx context.Context, user entities.User) (string, error) {
	return "", nil
}

// Authorization выполняет авторизацию пользователя.
func (uc *KeeperUseCases) Authorization(ctx context.Context, user entities.User) (string, error) {
	return "", nil
}

// AddAuthData выполняет сохранение пары логин/пароль.
func (uc *KeeperUseCases) AddAuthData(ctx context.Context, data entities.AuthData) error {
	return nil
}

// AddTextData выполняет сохранение текстовых данных.
func (uc *KeeperUseCases) AddTextData(ctx context.Context, data entities.TextData) error {
	return nil
}

// AddBinaryData выполняет сохранение бинарных данных.
func (uc *KeeperUseCases) AddBinaryData(ctx context.Context, data entities.File) error {
	return nil
}

// AddBankData выполняет сохранение данных банковских карт.
func (uc *KeeperUseCases) AddBankData(ctx context.Context, data entities.BankData) error {
	return nil
}

// UpdateAuthData выполняет обновление пары логин/пароль.
func (uc *KeeperUseCases) UpdateAuthData(ctx context.Context, data entities.AuthData) error {
	return nil
}

// UpdateTextData выполняет обновление текстовых данных.
func (uc *KeeperUseCases) UpdateTextData(ctx context.Context, data entities.TextData) error {
	return nil
}

// UpdateBinaryData выполняет обновление бинарных данных.
func (uc *KeeperUseCases) UpdateBinaryData(ctx context.Context, data entities.File) error {
	return nil
}

// UpdateBankData выполняет обновление данных банковских карт.
func (uc *KeeperUseCases) UpdateBankData(ctx context.Context, data entities.BankData) error {
	return nil
}

// DeleteAuthData выполняет удаление пары логин/пароль.
func (uc *KeeperUseCases) DeleteAuthData(ctx context.Context, data entities.AuthData) error {
	return nil
}

// DeleteTextData выполняет удаление текстовых данных.
func (uc *KeeperUseCases) DeleteTextData(ctx context.Context, data entities.TextData) error {
	return nil
}

// DeleteBinaryData выполняет удаление бинарных данных.
func (uc *KeeperUseCases) DeleteBinaryData(ctx context.Context, data entities.File) error {
	return nil
}

// DeleteBankData выполняет удаление данных банковских карт.
func (uc *KeeperUseCases) DeleteBankData(ctx context.Context, data entities.BankData) error {
	return nil
}

// AuthDataList возвращает список сохраненных пар логин/пароль.
func (uc *KeeperUseCases) AuthDataList(ctx context.Context, userID int64) ([]entities.AuthData, error) {
	return nil, nil
}

// TextLabelsList возвращает список заголовков сохранённых текстовых данных.
func (uc *KeeperUseCases) TextLabelsList(ctx context.Context, userID int64) ([]string, error) {
	return nil, nil
}

// TextData возвращает сохранённые текстовые данные по заголовку.
func (uc *KeeperUseCases) TextData(ctx context.Context, data *entities.TextData) error {
	return nil
}

// FileNamesList возвращает список сохранённых файлов.
func (uc *KeeperUseCases) FileNamesList(ctx context.Context, userID int64) ([]string, error) {
	return nil, nil
}

// BinaryData возвращает сохранённые бинарные данные по имени файла.
func (uc *KeeperUseCases) BinaryData(ctx context.Context, data *entities.File) error {
	return nil
}

// BankCardList возвращает список номеров банковских карт.
func (uc *KeeperUseCases) BankCardNumbersList(ctx context.Context, userID int64) ([]string, error) {
	return nil, nil
}

// BankCardList возвращает данные банковской карты по номеру.
func (uc *KeeperUseCases) BankCard(ctx context.Context, data *entities.BankData) error {
	return nil
}
