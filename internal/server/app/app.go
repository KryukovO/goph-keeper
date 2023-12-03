// Package server содержит реализацию сервера.
package app

import (
	"context"
	"net"
	"os/signal"
	"syscall"

	api "github.com/KryukovO/goph-keeper/api/serverpb"
	"github.com/KryukovO/goph-keeper/internal/server/config"
	localstorage "github.com/KryukovO/goph-keeper/internal/server/filestorage/local-storage"
	sgrpc "github.com/KryukovO/goph-keeper/internal/server/grpc"
	"github.com/KryukovO/goph-keeper/internal/server/repository/pgrepo"
	"github.com/KryukovO/goph-keeper/internal/server/usecases"
	"github.com/KryukovO/goph-keeper/pkg/postgres"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

// Server - реализация сервера.
type App struct {
	cfg        *config.Config
	grpcServer *grpc.Server
	log        *logrus.Logger
}

// NewServer возвращает объект Server.
func NewApp(cfg *config.Config, log *logrus.Logger) *App {
	return &App{
		cfg: cfg,
		log: log,
	}
}

// Run выполняет запуск сервера.
func (a *App) Run(ctx context.Context) error {
	sigCtx, sigCancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer sigCancel()

	a.log.Info("Connecting to the database...")

	repoCtx, cancel := context.WithTimeout(ctx, a.cfg.RepositoryTimeout)
	defer cancel()

	db, err := postgres.NewPostgres(repoCtx, a.cfg.DSN)
	if err != nil {
		return err
	}

	a.log.Info("Database connection established")

	defer func() {
		db.Close()

		a.log.Info("Database connection closed")
	}()

	a.log.Info("Initializing file storage...")

	fs, err := localstorage.NewLocalStorage(a.cfg.FSFolder)
	if err != nil {
		return err
	}

	a.log.Info("File storage is initialized")

	defer func() {
		fs.Close()

		a.log.Info("File storage closed")
	}()

	repo := pgrepo.NewPgRepo(db)

	user := usecases.NewUserUseCase(repo, a.cfg.RepositoryTimeout)
	auth := usecases.NewAuthDataUseCase(repo, a.cfg.RepositoryTimeout)
	txt := usecases.NewTextDataUseCase(repo, a.cfg.RepositoryTimeout)
	bank := usecases.NewBankDataUseCase(repo, a.cfg.RepositoryTimeout)

	binary, err := usecases.NewBinaryDataUseCase(ctx, repo, fs, a.cfg.RepositoryTimeout)
	if err != nil {
		return err
	}

	itcManager := sgrpc.NewManager([]byte(a.cfg.SecretKey), a.log)

	a.grpcServer = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			itcManager.LoggingUnaryInterceptor,
			itcManager.AuthUnaryInterceptor,
		),
		grpc.ChainStreamInterceptor(
			itcManager.LoggingStreamInterceptor,
			itcManager.AuthStreamInterceptor,
		),
	)

	keeperServer, err := sgrpc.NewKeeperServer(
		user, auth, txt, bank, binary,
		a.log,
	)
	if err != nil {
		return err
	}

	group, groupCtx := errgroup.WithContext(ctx)

	group.Go(func() error { return a.runGRPCServer(keeperServer) })

	group.Go(func() error {
		select {
		case <-groupCtx.Done():
			return nil
		case <-sigCtx.Done():
		}

		a.shutdown()

		return nil
	})

	return group.Wait()
}

func (a *App) runGRPCServer(storageServer *sgrpc.KeeperServer) error {
	a.log.Infof("Run gRPC-server at %s...", a.cfg.Address)

	listen, err := net.Listen("tcp", a.cfg.Address)
	if err != nil {
		return err
	}

	api.RegisterKeeperServer(a.grpcServer, storageServer)

	if err := a.grpcServer.Serve(listen); err != nil {
		return err
	}

	return nil
}

func (a *App) shutdown() {
	a.log.Info("Stopping server...")

	a.grpcServer.GracefulStop()

	a.log.Info("gRPC-server stopped gracefully")
}
