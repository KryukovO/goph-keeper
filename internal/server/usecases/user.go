package usecases

import (
	"context"
	"time"

	"github.com/KryukovO/goph-keeper/internal/entities"
	"github.com/KryukovO/goph-keeper/internal/server/repository"
)

// UserUseCase реализует логику взаимодействия
// с репозиторием для управления пользователями.
type UserUseCase struct {
	repo repository.UserRepository

	timeout time.Duration
}

// NewAuthDataUseCase возвращает новый объект AuthDataUseCase.
func NewUserUseCase(repo repository.UserRepository, timeout time.Duration) *UserUseCase {
	return &UserUseCase{
		repo:    repo,
		timeout: timeout,
	}
}

// Registration выполняет регистрацию пользователя.
func (uc *UserUseCase) Registration(ctx context.Context, user entities.User) (string, error) {
	return "", nil
}

// Authorization выполняет авторизацию пользователя.
func (uc *UserUseCase) Authorization(ctx context.Context, user entities.User) (string, error) {
	return "", nil
}
