package sentry

import (
	"log"
	"os"
	"time"

	sentry "github.com/getsentry/sentry-go"
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
