package usecases

import (
	"context"
	"time"

	"github.com/KryukovO/goph-keeper/internal/entities"
	"github.com/KryukovO/goph-keeper/internal/server/repository"
)

// BankDataUseCase реализует логику взаимодействия
// с репозиторием для управления данными банковских карт пользователя.
type BankDataUseCase struct {
	repo repository.BankDataRepository

	timeout time.Duration
}

// NewBankDataUseCase возвращает новый объект BankDataUseCase.
func NewBankDataUseCase(repo repository.BankDataRepository, timeout time.Duration) *BankDataUseCase {
	return &BankDataUseCase{
		repo:    repo,
		timeout: timeout,
	}
}

// AddBankData выполняет сохранение данных банковских карт.
func (uc *BankDataUseCase) AddBankData(ctx context.Context, data entities.BankData) error {
	return nil
}

// UpdateBankData выполняет обновление данных банковских карт.
func (uc *BankDataUseCase) UpdateBankData(ctx context.Context, data entities.BankData) error {
	return nil
}

// DeleteBankData выполняет удаление данных банковских карт.
func (uc *BankDataUseCase) DeleteBankData(ctx context.Context, data entities.BankData) error {
	return nil
}

// BankCardList возвращает список номеров банковских карт.
func (uc *BankDataUseCase) BankCardNumbersList(ctx context.Context, userID int64) ([]string, error) {
	return []string{}, nil
}

// BankCard возвращает данные банковской карты по номеру.
func (uc *BankDataUseCase) BankCard(ctx context.Context, data *entities.BankData) error {
	return nil
}
