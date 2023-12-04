package app

import (
	"context"

	"github.com/KryukovO/goph-keeper/api/serverpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// RequestWithLogin служит для определения принадлежности gRPC-запроса
// к группе управления пользователями.
type RequestWithLogin interface {
	GetLogin() string
	GetPassword() string
}

func (a *App) unaryInterceptor(ctx context.Context, method string, req interface{},
	reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	// Пропускаем запросы на регистрацию и авторизацию
	if _, ok := req.(RequestWithLogin); ok {
		return invoker(ctx, method, req, reply, cc, opts...)
	}

	tokenCtx := metadata.AppendToOutgoingContext(ctx, "token", a.token)

	err := invoker(tokenCtx, method, req, reply, cc, opts...)
	if err != nil {
		if status.Code(err) == codes.Unauthenticated {
			resp, err := a.client.Authorization(ctx, &serverpb.AuthorizationRequest{
				Login:    a.user.Login,
				Password: a.user.Password,
			})
			if err != nil {
				return err
			}

			a.token = resp.GetToken()

			tokenCtx = metadata.AppendToOutgoingContext(ctx, "token", a.token)

			return invoker(tokenCtx, method, req, reply, cc, opts...)
		}
	}

	return nil
}

func (a *App) streamInterceptor(
	ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn,
	method string, streamer grpc.Streamer, opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	tokenCtx := metadata.AppendToOutgoingContext(ctx, "token", a.token)

	stream, err := streamer(tokenCtx, desc, cc, method, opts...)
	if err != nil {
		if status.Code(err) == codes.Unauthenticated {
			resp, err := a.client.Authorization(ctx, &serverpb.AuthorizationRequest{
				Login:    a.user.Login,
				Password: a.user.Password,
			})
			if err != nil {
				return nil, err
			}

			a.token = resp.GetToken()

			tokenCtx = metadata.AppendToOutgoingContext(ctx, "token", a.token)

			return streamer(tokenCtx, desc, cc, method, opts...)
		}
	}

	return stream, nil
}
