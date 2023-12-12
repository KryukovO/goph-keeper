// Package grpc содержит реализацию gRPC-компоненты сервера.
package grpc

import (
	"context"

	api "github.com/KryukovO/goph-keeper/api/serverpb"
	"github.com/KryukovO/goph-keeper/internal/server/usecases"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sirupsen/logrus"
)

// KeeperServer - реализация gRPC компоненты сервера.
type KeeperServer struct {
	api.UnimplementedKeeperServer

	userUC   usecases.User
	authUC   usecases.AuthData
	textUC   usecases.TextData
	bankUC   usecases.BankData
	binaryUC usecases.BinaryData

	log *logrus.Logger
}

// NewKeeperServer возвращает новый объект KeeperServer.
func NewKeeperServer(
	userUC usecases.User, authUC usecases.AuthData,
	textUC usecases.TextData, bankUC usecases.BankData,
	binaryUC usecases.BinaryData,
	log *logrus.Logger,
) (*KeeperServer, error) {
	return &KeeperServer{
		userUC:   userUC,
		authUC:   authUC,
		textUC:   textUC,
		bankUC:   bankUC,
		binaryUC: binaryUC,
		log:      log,
	}, nil
}

// Registration выполняет регистрацию пользователя.
// Возвращает JSON Web Token, который необходимо использовать для аутентификации.
func (s KeeperServer) Registration(
	ctx context.Context, req *api.RegistrationRequest,
) (*api.RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registration not implemented")
}

// Authorization выполняет авторизацию пользователя.
// Возвращает JSON Web Token, который необходимо использовать для аутентификации.
func (s KeeperServer) Authorization(
	ctx context.Context, req *api.AuthorizationRequest,
) (*api.AuthorizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorization not implemented")
}

// AddAuthData выполняет сохранение пары логин/пароль в репозиторий.
func (s KeeperServer) AddAuthData(ctx context.Context, req *api.AddAuthDataRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAuthData not implemented")
}

// AddTextData выполняет сохранение текстовых данных в репозиторий.
func (s KeeperServer) AddTextData(ctx context.Context, req *api.AddTextDataRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTextData not implemented")
}

// AddBinaryData выполняет сохранение бинарных данных в хранилище.
func (s KeeperServer) AddBinaryData(stream api.Keeper_AddBinaryDataServer) error {
	return status.Errorf(codes.Unimplemented, "method AddBinaryData not implemented")
}

// AddBankData выполняет сохранение данных банковских карт в репозиторий.
func (s KeeperServer) AddBankData(ctx context.Context, req *api.AddBankDataRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBankData not implemented")
}

// UpdateAuthData выполняет обновление пары логин/пароль в репозитории.
func (s KeeperServer) UpdateAuthData(ctx context.Context, req *api.UpdateAuthDataRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuthData not implemented")
}

// UpdateTextData выполняет обновление текстовых данных в репозитории.
func (s KeeperServer) UpdateTextData(ctx context.Context, req *api.UpdateTextDataRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTextData not implemented")
}

// UpdateBankData выполняет обновление данных банковских карт в репозитории.
func (s KeeperServer) UpdateBankData(ctx context.Context, req *api.UpdateBankDataRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBankData not implemented")
}

// DeleteAuthData выполняет удаление пары логин/пароль из репозитория.
func (s KeeperServer) DeleteAuthData(ctx context.Context, req *api.DeleteAuthDataRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuthData not implemented")
}

// DeleteTextData выполняет удаление текстовых данных из репозитория.
func (s KeeperServer) DeleteTextData(ctx context.Context, req *api.DeleteTextDataRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTextData not implemented")
}

// DeleteBinaryData выполняет удаление бинарных данных из хранилища.
func (s KeeperServer) DeleteBinaryData(ctx context.Context, req *api.DeleteBinaryDataRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBinaryData not implemented")
}

// DeleteBankData выполняет удаление данных банковских карт из репозитория.
func (s KeeperServer) DeleteBankData(ctx context.Context, req *api.DeleteBankDataRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBankData not implemented")
}

// AuthDataList возвращает список сохраненных пар логин/пароль.
func (s KeeperServer) AuthDataList(ctx context.Context, _ *empty.Empty) (*api.AuthDataListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthDataList not implemented")
}

// TextLabelsList возвращает список заголовков сохранённых текстовых данных.
func (s KeeperServer) TextLabelsList(ctx context.Context, _ *empty.Empty) (*api.TextLabelsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TextLabelsList not implemented")
}

// TextData возвращает сохранённые текстовые данные по заголовку.
func (s KeeperServer) TextData(ctx context.Context, req *api.TextDataRequest) (*api.TextDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TextData not implemented")
}

// FileNamesList возвращает список сохранённых файлов.
func (s KeeperServer) FileNamesList(ctx context.Context, _ *empty.Empty) (*api.FileNamesListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileNamesList not implemented")
}

// BinaryData возвращает сохранённые бинарные данные по имени файла.
func (s KeeperServer) BinaryData(req *api.BinaryDataRequest, stream api.Keeper_BinaryDataServer) error {
	return status.Errorf(codes.Unimplemented, "method BinaryData not implemented")
}

// BankCardList возвращает список номеров банковских карт.
func (s KeeperServer) BankCardNumbersList(ctx context.Context, _ *empty.Empty) (*api.BankCardListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BankCardNumbersList not implemented")
}

// BankCardList возвращает данные банковской карты по номеру.
func (s KeeperServer) BankCard(ctx context.Context, req *api.BankCardRequest) (*api.BankCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BankCard not implemented")
}
