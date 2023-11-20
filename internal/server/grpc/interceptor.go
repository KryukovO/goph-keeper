package grpc

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Manager предназначен для управления interceptors.
type Manager struct {
	secret []byte
	log    *logrus.Logger
}

// NewManager возвращает новый объект Manager.
func NewManager(secret []byte, log *logrus.Logger) *Manager {
	return &Manager{
		secret: secret,
		log:    log,
	}
}

// LoggingInterceptor - выполняет логгирование входящего gRPC запроса.
func (itc *Manager) LoggingInterceptor(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	uuid := uuid.New()

	uuidCtx := metadata.AppendToOutgoingContext(ctx, "uuid", uuid.String())

	itc.log.Infof("[%s] received gRPC request: %s", uuid, info.FullMethod)

	ts := time.Now()
	resp, err := handler(uuidCtx, req)

	if err != nil {
		st, _ := status.FromError(err)

		itc.log.Infof(
			"[%s] query response status: %d; duration: %s",
			uuid, st.Code(), time.Since(ts),
		)
	} else {
		itc.log.Infof(
			"[%s] query response status: OK; duration: %s",
			uuid, time.Since(ts),
		)
	}

	return resp, err
}

// AuthInterceptor - выполняет аутентификацию пользователя из входящего gRPC запроса.
func (itc *Manager) AuthInterceptor(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	// NOTE: JWT будет находится в метаданных по ключу 'token'
	return handler(ctx, req)
}
