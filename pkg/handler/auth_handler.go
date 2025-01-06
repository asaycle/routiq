package handler

import (
	"context"
	"log"

	pb "github.com/asaycle/routiq.git/api/proto/v1"
	"github.com/asaycle/routiq.git/pkg/domain/repository"
	"github.com/asaycle/routiq.git/pkg/domain/usecase"
	"github.com/asaycle/routiq.git/pkg/infrastructure/db"
	"github.com/asaycle/routiq.git/pkg/server/handlers"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

func init() {
	handlers.RegisterHandler(NewAuthHandler())
}

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	useCase usecase.AuthUsecase
}

func NewAuthHandler() *AuthHandler {
	pgdb, err := db.NewPgDB("localhost", 5432, "root", "root", "routiq")
	if err != nil {
		log.Panic("failed initialize pgdb", err)
	}
	txManager := repository.NewTransactionManager(pgdb)
	return &AuthHandler{
		useCase: usecase.NewAuthUsecaseImpl(
			repository.NewUserRepositoryImpl(),
			repository.NewUserProfileRepositoryImpl(),
			repository.NewSessionRepositoryImpl(),
			txManager,
		),
	}
}

func (h *AuthHandler) Register(s *grpc.Server) {
	pb.RegisterAuthServiceServer(s, h)
}

func (h *AuthHandler) ExchangeOAuthCode(ctx context.Context, req *pb.ExchangeOAuthCodeRequest) (*pb.ExchangeOAuthCodeResponse, error) {
	log.Printf("CreateUser: %v", req)
	resp, err := h.useCase.CreateUser(ctx, req.Code)
	if err != nil {
		return nil, xerrors.Errorf("Error CreateUser: %w", err)
	}
	return &pb.ExchangeOAuthCodeResponse{
		RedirectUrl:  resp.RedirectURL,
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		Role:         resp.Role,
	}, nil
}

func (h *AuthHandler) VerifyToken(ctx context.Context, req *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error) {
	return &pb.VerifyTokenResponse{
		UserId: "testUser",
		Role:   "user",
	}, nil
}

var _ pb.AuthServiceServer = &AuthHandler{}
