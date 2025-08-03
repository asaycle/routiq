# Dockerfile.backend

# ビルドステージ
FROM golang:1.23.4 AS builder

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/main.go

# 実行ステージ
FROM alpine

WORKDIR /app

COPY --from=builder /app/app /app/app
COPY ./db/migration /app/migrations
COPY ./entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]