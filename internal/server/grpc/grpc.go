// Package grpc содержит реализацию gRPC-компоненты сервера.
package grpc

import (
	"bytes"
	"context"
	"errors"
	"io"
	"strconv"

	api "github.com/KryukovO/goph-keeper/api/serverpb"
	"github.com/KryukovO/goph-keeper/internal/entities"
	"github.com/KryukovO/goph-keeper/internal/server/usecases"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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
	user := entities.User{
		Login:        req.GetLogin(),
		Password:     req.GetPassword(),
		Subscription: entities.MakeSubscription(req.GetSubscription()),
	}

	userID, token, err := s.userUC.Registration(ctx, user)

	if errors.Is(err, entities.ErrUserAlreadyExists) {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	if errors.Is(err, entities.ErrInvalidLoginPassword) {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	s.binaryUC.UpdateSubscription(ctx, userID, entities.MakeSubscription(req.GetSubscription()))

	return &api.RegistrationResponse{Token: token}, nil
}

// Authorization выполняет авторизацию пользователя.
// Возвращает JSON Web Token, который необходимо использовать для аутентификации.
func (s KeeperServer) Authorization(
	ctx context.Context, req *api.AuthorizationRequest,
) (*api.AuthorizationResponse, error) {
	user := entities.User{
		Login:    req.GetLogin(),
		Password: req.GetPassword(),
	}

	token, err := s.userUC.Authorization(ctx, user)

	if errors.Is(err, entities.ErrInvalidLoginPassword) {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.AuthorizationResponse{Token: token}, nil
}

// AddAuthData выполняет сохранение пары логин/пароль в репозиторий.
func (s KeeperServer) AddAuthData(ctx context.Context, req *api.AddAuthDataRequest) (*empty.Empty, error) {
	userID, err := userFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	data := entities.AuthData{
		UserID:   userID,
		Resource: req.GetData().GetResource(),
		Login:    req.GetData().GetLogin(),
		Password: req.GetData().GetUserPassword(),
		Metadata: req.GetData().GetMetadata(),
	}

	err = s.authUC.AddAuthData(ctx, data)

	if errors.Is(err, entities.ErrAuthDataAlreadyExists) {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

// AddTextData выполняет сохранение текстовых данных в репозиторий.
func (s KeeperServer) AddTextData(ctx context.Context, req *api.AddTextDataRequest) (*empty.Empty, error) {
	userID, err := userFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	data := entities.TextData{
		UserID:   userID,
		Label:    req.GetData().GetLabel(),
		Data:     req.GetData().GetText(),
		Metadata: req.GetData().GetMetadata(),
	}

	err = s.textUC.AddTextData(ctx, data)

	if errors.Is(err, entities.ErrTextDataAlreadyExists) {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

// AddBinaryData выполняет сохранение бинарных данных в хранилище.
func (s KeeperServer) AddBinaryData(stream api.Keeper_AddBinaryDataServer) error {
	userID, err := userFromCtx(stream.Context())
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	resp, err := stream.Recv()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	file := entities.File{
		UserID:   userID,
		FileName: resp.GetData().GetFileName(),
	}

	data := bytes.Buffer{}

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		chunk := req.GetData().GetChunk()

		_, err = data.Write(chunk)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	file.Data = data

	err = s.binaryUC.AddBinaryData(stream.Context(), file)

	if errors.Is(err, entities.ErrFileIsTooBig) {
		return status.Error(codes.FailedPrecondition, err.Error())
	}

	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

// AddBankData выполняет сохранение данных банковских карт в репозиторий.
func (s KeeperServer) AddBankData(ctx context.Context, req *api.AddBankDataRequest) (*empty.Empty, error) {
	userID, err := userFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	data := entities.BankData{
		UserID:         userID,
		Number:         req.GetData().GetNumber(),
		CardholderName: req.GetData().GetCardholderName(),
		ExpiredAt:      req.GetData().GetExpirationDate(),
		CVV:            req.GetData().GetCVV(),
		Metadata:       req.GetData().GetMetadata(),
	}

	err = s.bankUC.AddBankData(ctx, data)

	if errors.Is(err, entities.ErrBankDataAlreadyExists) {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

// UpdateAuthData выполняет обновление пары логин/пароль в репозитории.
func (s KeeperServer) UpdateAuthData(ctx context.Context, req *api.UpdateAuthDataRequest) (*empty.Empty, error) {
	userID, err := userFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	data := entities.AuthData{
		UserID:   userID,
		Resource: req.GetData().GetResource(),
		Login:    req.GetData().GetLogin(),
		Password: req.GetData().GetUserPassword(),
		Metadata: req.GetData().GetMetadata(),
	}

	err = s.authUC.UpdateAuthData(ctx, req.GetOldResource(), req.GetOldLogin(), data)

	if errors.Is(err, entities.ErrAuthDataAlreadyExists) {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

// UpdateTextData выполняет обновление текстовых данных в репозитории.
func (s KeeperServer) UpdateTextData(ctx context.Context, req *api.UpdateTextDataRequest) (*empty.Empty, error) {
	userID, err := userFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	data := entities.TextData{
		UserID:   userID,
		Label:    req.GetData().GetLabel(),
		Data:     req.GetData().GetText(),
		Metadata: req.GetData().GetMetadata(),
	}

	err = s.textUC.UpdateTextData(ctx, req.GetOldLabel(), data)

	if errors.Is(err, entities.ErrTextDataAlreadyExists) {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

// UpdateBankData выполняет обновление данных банковских карт в репозитории.
func (s KeeperServer) UpdateBankData(ctx context.Context, req *api.UpdateBankDataRequest) (*empty.Empty, error) {
	userID, err := userFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	data := entities.BankData{
		UserID:         userID,
		Number:         req.GetData().GetNumber(),
		CardholderName: req.GetData().GetCardholderName(),
		ExpiredAt:      req.GetData().GetExpirationDate(),
		CVV:            req.GetData().GetCVV(),
		Metadata:       req.GetData().GetMetadata(),
	}

	err = s.bankUC.UpdateBankData(ctx, req.GetOldNumber(), data)

	if errors.Is(err, entities.ErrBankDataAlreadyExists) {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

// DeleteAuthData выполняет удаление пары логин/пароль из репозитория.
func (s KeeperServer) DeleteAuthData(ctx context.Context, req *api.DeleteAuthDataRequest) (*empty.Empty, error) {
	userID, err := userFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	data := entities.AuthData{
		UserID:   userID,
		Resource: req.GetResource(),
		Login:    req.GetLogin(),
	}

	err = s.authUC.DeleteAuthData(ctx, data)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

// DeleteTextData выполняет удаление текстовых данных из репозитория.
func (s KeeperServer) DeleteTextData(ctx context.Context, req *api.DeleteTextDataRequest) (*empty.Empty, error) {
	userID, err := userFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	data := entities.TextData{
		UserID: userID,
		Label:  req.GetLabel(),
	}

	err = s.textUC.DeleteTextData(ctx, data)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

// DeleteBinaryData выполняет удаление бинарных данных из хранилища.
func (s KeeperServer) DeleteBinaryData(ctx context.Context, req *api.DeleteBinaryDataRequest) (*empty.Empty, error) {
	userID, err := userFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	data := entities.File{
		UserID:   userID,
		FileName: req.GetFileName(),
	}

	err = s.binaryUC.DeleteBinaryData(ctx, data)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

// DeleteBankData выполняет удаление данных банковских карт из репозитория.
func (s KeeperServer) DeleteBankData(ctx context.Context, req *api.DeleteBankDataRequest) (*empty.Empty, error) {
	userID, err := userFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	data := entities.BankData{
		UserID: userID,
		Number: req.GetNumber(),
	}

	err = s.bankUC.DeleteBankData(ctx, data)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

// AuthDataList возвращает список сохраненных пар логин/пароль.
func (s KeeperServer) AuthDataList(ctx context.Context, _ *empty.Empty) (*api.AuthDataListResponse, error) {
	userID, err := userFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	list, err := s.authUC.AuthDataList(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := make([]*api.AuthData, 0, len(list))
	for _, item := range list {
		res = append(res, &api.AuthData{
			Resource:     item.Resource,
			Login:        item.Login,
			UserPassword: item.Password,
			Metadata:     item.Metadata,
		})
	}

	return &api.AuthDataListResponse{Data: res}, nil
}

// TextLabelsList возвращает список заголовков сохранённых текстовых данных.
func (s KeeperServer) TextLabelsList(ctx context.Context, _ *empty.Empty) (*api.TextLabelsListResponse, error) {
	userID, err := userFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	list, err := s.textUC.TextLabelsList(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.TextLabelsListResponse{Labels: list}, nil
}

// TextData возвращает сохранённые текстовые данные по заголовку.
func (s KeeperServer) TextData(ctx context.Context, req *api.TextDataRequest) (*api.TextDataResponse, error) {
	userID, err := userFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	data := entities.TextData{
		UserID: userID,
		Label:  req.GetLabel(),
	}

	err = s.textUC.TextData(ctx, &data)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.TextDataResponse{Data: &api.TextData{
		Label:    data.Label,
		Text:     data.Data,
		Metadata: data.Metadata,
	}}, nil
}

// FileNamesList возвращает список сохранённых файлов.
func (s KeeperServer) FileNamesList(ctx context.Context, _ *empty.Empty) (*api.FileNamesListResponse, error) {
	userID, err := userFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	list := s.binaryUC.FileNamesList(ctx, userID)

	return &api.FileNamesListResponse{FileNames: list}, nil
}

// BinaryData возвращает сохранённые бинарные данные по имени файла.
func (s KeeperServer) BinaryData(req *api.BinaryDataRequest, stream api.Keeper_BinaryDataServer) error {
	userID, err := userFromCtx(stream.Context())
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	data := entities.File{
		UserID:   userID,
		FileName: req.GetFileName(),
	}

	err = s.binaryUC.BinaryData(stream.Context(), &data)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	err = stream.Send(&api.BinaryDataResponse{
		Data: &api.BinaryData{Data: &api.BinaryData_FileName{FileName: data.FileName}},
	})
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	buffer := make([]byte, 1024)

	for {
		n, err := data.Data.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		err = stream.Send(&api.BinaryDataResponse{Data: &api.BinaryData{Data: &api.BinaryData_Chunk{Chunk: buffer[:n]}}})
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}

// BankCardList возвращает список номеров банковских карт.
func (s KeeperServer) BankCardNumbersList(ctx context.Context, _ *empty.Empty) (*api.BankCardListResponse, error) {
	userID, err := userFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	list, err := s.bankUC.BankCardNumbersList(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.BankCardListResponse{CardNumbers: list}, nil
}

// BankCardList возвращает данные банковской карты по номеру.
func (s KeeperServer) BankCard(ctx context.Context, req *api.BankCardRequest) (*api.BankCardResponse, error) {
	userID, err := userFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	data := entities.BankData{
		UserID: userID,
		Number: req.GetCardNumber(),
	}

	err = s.bankUC.BankCard(ctx, &data)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.BankCardResponse{
		Data: &api.BankData{
			Number:         data.Number,
			CardholderName: data.CardholderName,
			ExpirationDate: data.ExpiredAt,
			CVV:            data.CVV,
			Metadata:       data.CVV,
		},
	}, nil
}

// userFromCtx извлекает из метаданных в контексте значение userID.
func userFromCtx(ctx context.Context) (int64, error) {
	var (
		userID int64
		err    error
	)

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get("userID")
		if len(values) > 0 {
			userID, err = strconv.ParseInt(values[0], 10, 64)
			if err != nil {
				return 0, err
			}
		}
	}

	return userID, nil
}
