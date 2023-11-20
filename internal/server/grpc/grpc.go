// Package grpc содержит реализацию gRPC-компоненты сервера.
package grpc

import (
	"context"

	api "github.com/KryukovO/goph-keeper/api/serverpb"
	"github.com/KryukovO/goph-keeper/internal/server/usecases"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sirupsen/logrus"
)

// KeeperServer - реализация gRPC компоненты сервера.
type KeeperServer struct {
	api.UnimplementedKeeperServer

	uc  usecases.UseCases
	log *logrus.Logger
}

// NewKeeperServer возвращает новый объект KeeperServer.
func NewKeeperServer(uc usecases.UseCases, log *logrus.Logger) (*KeeperServer, error) {
	return &KeeperServer{
		uc:  uc,
		log: log,
	}, nil
}

// Registration выполняет регистрацию пользователя.
// Возвращает JSON Web Token, который необходимо использовать для аутентификации.
func (s KeeperServer) Registration(context.Context, *api.RegistrationRequest) (*api.RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registration not implemented")
}

// Authorization выполняет авторизацию пользователя.
// Возвращает JSON Web Token, который необходимо использовать для аутентификации.
func (s KeeperServer) Authorization(context.Context, *api.AuthorizationRequest) (*api.AuthorizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorization not implemented")
}
