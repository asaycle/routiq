package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/asaycle/routiq.git/api/proto/v1"
	_ "github.com/asaycle/routiq.git/pkg/handler"
	"github.com/asaycle/routiq.git/pkg/server/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/proto"
)

type Ports struct {
	Grpc    int
	Gateway int
}

type Server struct {
	Ports           *Ports
	EnabledHandlers []handlers.Handler
}

func NewServer() *Server {
	return &Server{
		Ports: &Ports{
			Grpc:    50051,
			Gateway: 8080,
		},
		EnabledHandlers: make([]handlers.Handler, 0),
	}
}

func (s *Server) Start() error {
	go func() {
		if err := s.startGatewayServer(); err != nil {
			log.Fatalf("Failed Start GatewayServer: %v", err)
		}
	}()
	return s.startGrpcServer()
}

func (s *Server) startGrpcServer() error {
	addr := fmt.Sprintf(":%d", s.Ports.Grpc)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %v", addr, err)
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
	log.Printf("test: %d", len(handlers.GetHandlers()))
	for _, handler := range handlers.GetHandlers() {
		handler.Register(srv)
	}
	reflection.Register(srv)

	log.Printf("gRPC server is running on %s", addr)
	return srv.Serve(listener)
}

func (s *Server) startGatewayServer() error {
	grpcAddr := "localhost:50051"
	httpAddr := "localhost:8081"
	mux := runtime.NewServeMux(
		runtime.WithForwardResponseOption(customHTTPResponseModifier),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterAuthServiceHandlerFromEndpoint(context.Background(), mux, grpcAddr, opts)
	if err != nil {
		log.Fatalf("Failed to start gRPC-Gateway: %v", err)
	}

	log.Printf("HTTP server listening on %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}
	return nil
}

func customHTTPResponseModifier(ctx context.Context, w http.ResponseWriter, p proto.Message) error {
	if resp, ok := p.(*pb.ExchangeOAuthCodeResponse); ok && resp.RedirectUrl != "" {
		// リダイレクトレスポンスを作成
		http.Redirect(w, nil, resp.RedirectUrl, http.StatusFound)
		return nil
	}
	return nil
}
