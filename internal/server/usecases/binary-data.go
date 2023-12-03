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
	ctx context.Context, repo repository.SubscriptionRepository,
	fs filestorage.FileStorage, timeout time.Duration,
) (*BinaryDataUseCase, error) {
	uc := &BinaryDataUseCase{
		repo:    repo,
		fs:      fs,
		timeout: timeout,
	}

	err := uc.initFileStorage(ctx)
	if err != nil {
		return nil, err
	}

	return uc, nil
}

// initFileStorage инициализирует файловое хранилище.
func (uc *BinaryDataUseCase) initFileStorage(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	subscriptions, err := uc.repo.Sunbsriptions(ctx)
	if err != nil {
		return err
	}

	uc.fs.SetSubscriptions(subscriptions)

	return nil
}

// AddBinaryData выполняет сохранение бинарных данных.
func (uc *BinaryDataUseCase) AddBinaryData(_ context.Context, data entities.File) error {
	return uc.fs.Save(data)
}

// DeleteBinaryData выполняет удаление бинарных данных.
func (uc *BinaryDataUseCase) DeleteBinaryData(_ context.Context, data entities.File) error {
	return uc.fs.Delete(data)
}

// FileNamesList возвращает список сохранённых файлов.
func (uc *BinaryDataUseCase) FileNamesList(_ context.Context, userID int64) ([]string, error) {
	return uc.fs.List(userID)
}

// BinaryData возвращает сохранённые бинарные данные по имени файла.
func (uc *BinaryDataUseCase) BinaryData(_ context.Context, data *entities.File) error {
	return uc.fs.Load(data)
}
