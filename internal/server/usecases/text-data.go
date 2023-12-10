package usecases

import (
	"context"
	"time"

	"github.com/KryukovO/goph-keeper/internal/entities"
	"github.com/KryukovO/goph-keeper/internal/server/repository"
)

// TextDataUseCase реализует логику взаимодействия
// с репозиторием для управления текстовыми данными пользователя.
type TextDataUseCase struct {
	repo repository.TextDataRepository

	timeout time.Duration
}

// NewTextDataUseCase возвращает новый объект TextDataUseCase.
func NewTextDataUseCase(repo repository.TextDataRepository, timeout time.Duration) *TextDataUseCase {
	return &TextDataUseCase{
		repo:    repo,
		timeout: timeout,
	}
}

// AddTextData выполняет сохранение текстовых данных.
func (uc *TextDataUseCase) AddTextData(ctx context.Context, data entities.TextData) error {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	return uc.repo.AddTextData(ctx, data)
}

// UpdateTextData выполняет обновление текстовых данных.
func (uc *TextDataUseCase) UpdateTextData(ctx context.Context, oldLabel string, data entities.TextData) error {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	return uc.repo.UpdateTextData(ctx, oldLabel, data)
}

// DeleteTextData выполняет удаление текстовых данных.
func (uc *TextDataUseCase) DeleteTextData(ctx context.Context, data entities.TextData) error {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	return uc.repo.DeleteTextData(ctx, data)
}

// TextLabelsList возвращает список заголовков сохранённых текстовых данных.
func (uc *TextDataUseCase) TextLabelsList(ctx context.Context, userID int64) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	return uc.repo.TextLabelsList(ctx, userID)
}

// TextData возвращает сохранённые текстовые данные по заголовку.
func (uc *TextDataUseCase) TextData(ctx context.Context, data *entities.TextData) error {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	return uc.repo.TextData(ctx, data)
}
