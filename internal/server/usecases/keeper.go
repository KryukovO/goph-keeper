package usecases

import (
	"context"
	"time"

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
func (uc *KeeperUseCases) Registration(ctx context.Context, login string, pswd string) (string, error) {
	return "", nil
}

// Authorization выполняет авторизацию пользователя.
func (uc *KeeperUseCases) Authorization(ctx context.Context, login string, pswd string) (string, error) {
	return "", nil
}

// Close выполняет закрытие соединения с репозиторием.
func (uc *KeeperUseCases) Close() error {
	return uc.repo.Close()
}
