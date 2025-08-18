package server

import (
	"context"
	"log"
	"log/slog"
	"strings"
	"time"

	"github.com/asaycle/routiq/pkg/lib/session"
	"github.com/asaycle/routiq/pkg/lib/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ctxKeyUID struct{}

func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// メタデータを取得
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		ctx = session.SetGuestRole(ctx)
		return handler(ctx, req)
	}

	// Authorization ヘッダーを取得
	authHeader := md["authorization"]
	if len(authHeader) == 0 {
		ctx = session.SetGuestRole(ctx)
		return handler(ctx, req)
	}

	// トークン部分を取得
	tokenStr := strings.TrimPrefix(authHeader[0], "Bearer ")
	claims, err := token.VerifyToken(tokenStr)
	if err != nil {
		slog.Info("failed verify token", slog.Any("err", err))
		return nil, status.Error(codes.Unauthenticated, "failed verify token")
	}
	ctx = session.SetUserRole(ctx, claims.Subject)
	return handler(ctx, req)
}

// UnaryInterceptor logs unary gRPC requests and responses.
func UnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()

	// Log method and request payload
	log.Printf("RPC method: %s", info.FullMethod)
	log.Printf("Request: %v", req)

	// Handle the request
	resp, err := handler(ctx, req)

	// Log response and error (if any)
	log.Printf("Response: %v", resp)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	// Log elapsed time
	log.Printf("Elapsed time: %v", time.Since(start))
	return resp, err
}

// StreamInterceptor logs stream gRPC requests and responses.
func StreamInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	start := time.Now()

	// Log method
	log.Printf("RPC method: %s", info.FullMethod)

	// Handle the stream
	err := handler(srv, ss)

	// Log error (if any)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	// Log elapsed time
	log.Printf("Elapsed time: %v", time.Since(start))
	return err
}
