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
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	return uc.repo.AddBankData(ctx, data)
}

// UpdateBankData выполняет обновление данных банковских карт.
func (uc *BankDataUseCase) UpdateBankData(ctx context.Context, oldNumber string, data entities.BankData) error {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	return uc.repo.UpdateBankData(ctx, oldNumber, data)
}

// DeleteBankData выполняет удаление данных банковских карт.
func (uc *BankDataUseCase) DeleteBankData(ctx context.Context, data entities.BankData) error {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	return uc.repo.DeleteBankData(ctx, data)
}

// BankCardList возвращает список номеров банковских карт.
func (uc *BankDataUseCase) BankCardNumbersList(ctx context.Context, userID int64) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	return uc.repo.BankCardNumbersList(ctx, userID)
}

// BankCard возвращает данные банковской карты по номеру.
func (uc *BankDataUseCase) BankCard(ctx context.Context, data *entities.BankData) error {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	return uc.repo.BankCard(ctx, data)
}
