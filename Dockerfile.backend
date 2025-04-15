# Dockerfile.backend

# ビルドステージ
FROM golang:1.23.4 AS builder

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/main.go

# 実行ステージ
FROM gcr.io/distroless/base

WORKDIR /app

COPY --from=builder /app/app .

CMD ["./app", "server", "start"]