package usecases

import (
	"context"
	"time"

	"github.com/KryukovO/goph-keeper/internal/entities"
	"github.com/KryukovO/goph-keeper/internal/server/repository"
	"github.com/KryukovO/goph-keeper/pkg/utils"
)

// UserUseCase реализует логику взаимодействия
// с репозиторием для управления пользователями.
type UserUseCase struct {
	repo repository.UserRepository

	timeout  time.Duration
	secret   []byte
	tokenTTL time.Duration
}

// NewAuthDataUseCase возвращает новый объект AuthDataUseCase.
func NewUserUseCase(
	repo repository.UserRepository, timeout time.Duration,
	secret []byte, tokenTTL time.Duration,
) *UserUseCase {
	return &UserUseCase{
		repo:     repo,
		timeout:  timeout,
		secret:   secret,
		tokenTTL: tokenTTL,
	}
}

// Registration выполняет регистрацию пользователя.
func (uc *UserUseCase) Registration(ctx context.Context, user entities.User) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	if user.Login == "" || user.Password == "" {
		return "", entities.ErrInvalidLoginPassword
	}

	userID, err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}

	return utils.BuildJWTString(uc.secret, uc.tokenTTL, userID)
}

// Authorization выполняет авторизацию пользователя.
func (uc *UserUseCase) Authorization(ctx context.Context, user entities.User) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	if user.Login == "" || user.Password == "" {
		return "", entities.ErrInvalidLoginPassword
	}

	err := uc.repo.User(ctx, &user)
	if err != nil {
		return "", err
	}

	err = user.Validate()
	if err != nil {
		return "", err
	}

	return utils.BuildJWTString(uc.secret, uc.tokenTTL, user.ID)
}
