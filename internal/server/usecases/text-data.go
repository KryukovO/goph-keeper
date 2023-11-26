package usecases

import (
	"context"
	"time"

	"github.com/KryukovO/goph-keeper/internal/server/entities"
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
	return nil
}

// UpdateTextData выполняет обновление текстовых данных.
func (uc *TextDataUseCase) UpdateTextData(ctx context.Context, data entities.TextData) error {
	return nil
}

// DeleteTextData выполняет удаление текстовых данных.
func (uc *TextDataUseCase) DeleteTextData(ctx context.Context, data entities.TextData) error {
	return nil
}

// TextLabelsList возвращает список заголовков сохранённых текстовых данных.
func (uc *TextDataUseCase) TextLabelsList(ctx context.Context, userID int64) ([]string, error) {
	return []string{}, nil
}

// TextData возвращает сохранённые текстовые данные по заголовку.
func (uc *TextDataUseCase) TextData(ctx context.Context, data *entities.TextData) error {
	return nil
}
