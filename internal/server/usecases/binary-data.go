package usecases

import (
	"context"
	"time"

	"github.com/KryukovO/goph-keeper/internal/entities"
	"github.com/KryukovO/goph-keeper/internal/server/filestorage"
	"github.com/KryukovO/goph-keeper/internal/server/repository"
)

// BinaryDataUseCase реализует логику взаимодействия
// с репозиторием для управления бинарными данными пользователя.
type BinaryDataUseCase struct {
	repo repository.SubscriptionRepository
	fs   filestorage.FileStorage

	timeout time.Duration
}

// NewBinaryDataUseCase возвращает новый объект BinaryDataUseCase.
func NewBinaryDataUseCase(
	repo repository.SubscriptionRepository, fs filestorage.FileStorage, timeout time.Duration,
) (*BinaryDataUseCase, error) {
	uc := &BinaryDataUseCase{
		repo:    repo,
		fs:      fs,
		timeout: timeout,
	}

	err := uc.initFileStorage()
	if err != nil {
		return nil, err
	}

	return uc, nil
}

// initFileStorage инициализирует файловое хранилище.
func (uc *BinaryDataUseCase) initFileStorage() error {
	subscriptions, err := uc.repo.Sunbsriptions()
	if err != nil {
		return err
	}

	uc.fs.SetSubscriptions(subscriptions)

	return nil
}

// AddBinaryData выполняет сохранение бинарных данных.
func (uc *BinaryDataUseCase) AddBinaryData(ctx context.Context, data entities.File) error {
	return nil
}

// DeleteBinaryData выполняет удаление бинарных данных.
func (uc *BinaryDataUseCase) DeleteBinaryData(ctx context.Context, data entities.File) error {
	return nil
}

// FileNamesList возвращает список сохранённых файлов.
func (uc *BinaryDataUseCase) FileNamesList(ctx context.Context, userID int64) ([]string, error) {
	return []string{}, nil
}

// BinaryData возвращает сохранённые бинарные данные по имени файла.
func (uc *BinaryDataUseCase) BinaryData(ctx context.Context, data *entities.File) error {
	return nil
}
