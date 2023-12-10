package usecases

import (
	"context"
	"time"

	"github.com/KryukovO/goph-keeper/internal/entities"
	"github.com/KryukovO/goph-keeper/internal/server/repository"
)

// AuthDataUseCase реализует логику взаимодействия
// с репозиторием для управления парами логин/пароль пользователя.
type AuthDataUseCase struct {
	repo repository.AuthDataRepository

	timeout time.Duration
}

// NewAuthDataUseCase возвращает новый объект AuthDataUseCase.
func NewAuthDataUseCase(repo repository.AuthDataRepository, timeout time.Duration) *AuthDataUseCase {
	return &AuthDataUseCase{
		repo:    repo,
		timeout: timeout,
	}
}

// AddAuthData выполняет сохранение пары логин/пароль.
func (uc *AuthDataUseCase) AddAuthData(ctx context.Context, data entities.AuthData) error {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	return uc.repo.AddAuthData(ctx, data)
}

// UpdateAuthData выполняет обновление пары логин/пароль.
func (uc *AuthDataUseCase) UpdateAuthData(
	ctx context.Context, oldResource, oldLogin string, data entities.AuthData,
) error {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	return uc.repo.UpdateAuthData(ctx, oldResource, oldLogin, data)
}

// DeleteAuthData выполняет удаление пары логин/пароль.
func (uc *AuthDataUseCase) DeleteAuthData(ctx context.Context, data entities.AuthData) error {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	return uc.repo.DeleteAuthData(ctx, data)
}

// AuthDataList возвращает список сохраненных пар логин/пароль.
func (uc *AuthDataUseCase) AuthDataList(ctx context.Context, userID int64) ([]entities.AuthData, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	return uc.repo.AuthDataList(ctx, userID)
}
