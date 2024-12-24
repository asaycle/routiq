package handlers

import (
	"sync"

	"google.golang.org/grpc"
)

type Handler interface {
	Register(*grpc.Server)
}

var (
	enabledHandlers = []Handler{}
	handlersMu      sync.Mutex
)

func RegisterHandler(handler Handler) {
	handlersMu.Lock()
	defer handlersMu.Unlock()
	enabledHandlers = append(enabledHandlers, handler)
}

func GetHandlers() []Handler {
	handlersMu.Lock()
	defer handlersMu.Unlock()
	return enabledHandlers
}
