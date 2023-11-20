// Package server содержит реализацию сервера.
package server

import (
	"context"
	"net"
	"os/signal"
	"syscall"

	api "github.com/KryukovO/goph-keeper/api/serverpb"
	"github.com/KryukovO/goph-keeper/internal/server/config"
	sgrpc "github.com/KryukovO/goph-keeper/internal/server/grpc"
	"github.com/KryukovO/goph-keeper/internal/server/repository/pgrepo"
	"github.com/KryukovO/goph-keeper/internal/server/usecases"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

// Server - реализация сервера.
type Server struct {
	cfg        *config.Config
	grpcServer *grpc.Server
	log        *logrus.Logger
}

// NewServer возвращает объект Server.
func NewServer(cfg *config.Config, log *logrus.Logger) *Server {
	return &Server{
		cfg: cfg,
		log: log,
	}
}

// Run выполняет запуск сервера.
func (s *Server) Run(ctx context.Context) error {
	sigCtx, sigCancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer sigCancel()

	repoCtx, cancel := context.WithTimeout(sigCtx, s.cfg.RepositoryTimeout)
	defer cancel()

	s.log.Info("Connecting to the repository...")

	repo, err := pgrepo.NewPgRepo(repoCtx, s.cfg.DSN)
	if err != nil {
		return err
	}

	s.log.Info("Repository connection established")

	keeper := usecases.NewKeeperUseCases(repo, s.cfg.RepositoryTimeout)
	defer func() {
		keeper.Close()

		s.log.Info("Repository connection closed")
	}()

	itcManager := sgrpc.NewManager([]byte(s.cfg.SecretKey), s.log)
	s.grpcServer = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			itcManager.LoggingInterceptor,
			itcManager.AuthInterceptor, // NOTE: нужно игнорировать методы Registration и Authorization
		),
	)

	keeperServer, err := sgrpc.NewKeeperServer(keeper, s.log)
	if err != nil {
		return err
	}

	group, groupCtx := errgroup.WithContext(ctx)

	group.Go(func() error { return s.runGRPCServer(keeperServer) })

	group.Go(func() error {
		select {
		case <-groupCtx.Done():
			return nil
		case <-sigCtx.Done():
		}

		s.log.Info("Stopping server...")

		s.shutdown()

		return nil
	})

	return group.Wait()
}

func (s *Server) runGRPCServer(storageServer *sgrpc.KeeperServer) error {
	s.log.Infof("Run gRPC-server at %s...", s.cfg.Address)

	listen, err := net.Listen("tcp", s.cfg.Address)
	if err != nil {
		return err
	}

	api.RegisterKeeperServer(s.grpcServer, storageServer)

	if err := s.grpcServer.Serve(listen); err != nil {
		return err
	}

	return nil
}

func (s *Server) shutdown() {
	s.grpcServer.GracefulStop()

	s.log.Info("gRPC-server stopped gracefully")
}
