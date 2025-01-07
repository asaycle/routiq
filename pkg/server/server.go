package server

import (
	"fmt"
	"log"
	"net"

	"github.com/asaycle/routiq.git/pkg/handler"
	"github.com/asaycle/routiq.git/pkg/lib/config"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Start() error {
	return s.startGrpcServer()
}

func (s *Server) startGrpcServer() error {
	addr := fmt.Sprintf(":%s", s.cfg.Server.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return xerrors.Errorf("failed to listen on %s: %w", addr, err)
	}

	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			authInterceptor,
			UnaryInterceptor,
		),
		grpc.ChainStreamInterceptor(
			StreamInterceptor,
		),
	)
	// ハンドラを登録
	handler.ActivateAuthHandler(srv, s.cfg)
	handler.ActivateRouteHandler(srv, s.cfg)
	handler.ActivateTagHandler(srv, s.cfg)
	handler.ActivateRidingHandler(srv, s.cfg)
	reflection.Register(srv)

	log.Printf("gRPC server is running on %s", addr)
	return srv.Serve(listener)
}
