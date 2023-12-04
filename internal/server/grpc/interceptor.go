package grpc

import (
	"context"
	"strconv"
	"time"

	"github.com/KryukovO/goph-keeper/pkg/utils"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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

// LoggingUnaryInterceptor - выполняет логгирование входящего unary gRPC запроса.
func (m *Manager) LoggingUnaryInterceptor(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	uuid := uuid.New()

	uuidCtx := metadata.AppendToOutgoingContext(ctx, "uuid", uuid.String())

	m.log.Infof("[%s] received gRPC request: %s", uuid, info.FullMethod)

	ts := time.Now()
	resp, err := handler(uuidCtx, req)

	if err != nil {
		st, _ := status.FromError(err)

		m.log.Infof(
			"[%s] query response status: %s: %s; duration: %s",
			uuid, st.Code(), st.Message(), time.Since(ts),
		)
	} else {
		m.log.Infof(
			"[%s] query response status: OK; duration: %s",
			uuid, time.Since(ts),
		)
	}

	return resp, err
}

// LoggingStreamInterceptor - выполняет логгирование входящего stream gRPC запроса.
func (m *Manager) LoggingStreamInterceptor(
	srv interface{}, stream grpc.ServerStream,
	info *grpc.StreamServerInfo, handler grpc.StreamHandler,
) error {
	uuid := uuid.New()
	sw := newStreamWrapper(stream)

	uuidCtx := metadata.AppendToOutgoingContext(sw.Context(), "uuid", uuid.String())
	sw.SetContext(uuidCtx)

	m.log.Infof("[%s] received gRPC request: %s", uuid, info.FullMethod)

	ts := time.Now()
	err := handler(srv, sw)

	if err != nil {
		st, _ := status.FromError(err)

		m.log.Infof(
			"[%s] query response status: %s: %s; duration: %s",
			uuid, st.Code(), st.Message(), time.Since(ts),
		)
	} else {
		m.log.Infof(
			"[%s] query response status: OK; duration: %s",
			uuid, time.Since(ts),
		)
	}

	return err
}

// AuthUnaryInterceptor - выполняет аутентификацию пользователя из входящего unary gRPC запроса.
func (m *Manager) AuthUnaryInterceptor(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	// Пропускаем запросы на регистрацию и авторизацию
	if _, ok := req.(RequestWithLogin); ok {
		return handler(ctx, req)
	}

	var (
		token  string
		userID int64
	)

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get("token")
		if len(values) > 0 {
			token = values[0]
		}
	}

	err := utils.ParseTokenString(&userID, token, m.secret)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	userCtx := metadata.AppendToOutgoingContext(ctx, "userID", strconv.FormatInt(userID, 10))

	return handler(userCtx, req)
}

// AuthStreamInterceptor - выполняет аутентификацию пользователя из входящего stream gRPC запроса.
func (m *Manager) AuthStreamInterceptor(
	srv interface{}, stream grpc.ServerStream,
	info *grpc.StreamServerInfo, handler grpc.StreamHandler,
) error {
	var (
		token  string
		userID int64
	)

	sw := newStreamWrapper(stream)

	if md, ok := metadata.FromIncomingContext(sw.Context()); ok {
		values := md.Get("token")
		if len(values) > 0 {
			token = values[0]
		}
	}

	err := utils.ParseTokenString(&userID, token, m.secret)
	if err != nil {
		return status.Error(codes.Unauthenticated, err.Error())
	}

	userCtx := metadata.AppendToOutgoingContext(sw.Context(), "userID", strconv.FormatInt(userID, 10))
	sw.SetContext(userCtx)

	return handler(srv, sw)
}
