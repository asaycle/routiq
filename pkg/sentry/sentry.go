package sentry

import (
	"context"
	"log"
	"os"
	"time"

	sentry "github.com/getsentry/sentry-go"
	"google.golang.org/grpc"
)

func InitializeSentry() {
	dsn, ok := os.LookupEnv("SENTRY_DSN")
	if !ok {
		log.Fatal("SENTRY_DSN environment variable is not set")
		return
	}
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         dsn,
		Environment: os.Getenv("ENV"),
	})
	if err != nil {
		log.Fatalf("sentry.Init: %v", err)
		return
	}
}

func Flush() {
	sentry.Flush(2 * time.Second)
}

func SentryUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				sentry.CurrentHub().Recover(r)
				sentry.Flush(2 * time.Second)
				panic(r) // re-raise panic
			}
		}()

		resp, err = handler(ctx, req)
		if err != nil {
			sentry.CaptureException(err)
			sentry.Flush(2 * time.Second)
		}

		return resp, err
	}
}

func SentryStreamInterceptor() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		defer func() {
			if r := recover(); r != nil {
				sentry.CurrentHub().Recover(r)
				sentry.Flush(2 * time.Second)
				panic(r) // 再スロー
			}
		}()

		err := handler(srv, ss)
		if err != nil {
			sentry.CaptureException(err)
			sentry.Flush(2 * time.Second)
		}

		return err
	}
}
