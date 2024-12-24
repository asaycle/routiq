#!/bin/sh

brew install protobuf

echo "Installing required CLI tools..."

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

exec $SHELL -l
echo "All tools installed successfully!"