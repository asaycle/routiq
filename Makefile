DB_NAME := routiq
DB_URL := postgres://root:root@localhost:5432/routiq?sslmode=disable

MIGRATION_DIR := db/migration

# マイグレーションツールの実行パス（インストール済み前提）
MIGRATE_CMD := migrate

# マイグレーションの適用
db-migration-up:
	$(MIGRATE_CMD) -database "$(DB_URL)" -path $(MIGRATION_DIR) up

# マイグレーションのロールバック（1段階）
db-migration-down:
	$(MIGRATE_CMD) -database "$(DB_URL)" -path $(MIGRATION_DIR) down 1

# マイグレーションのステータス確認
db-migration-status:
	$(MIGRATE_CMD) -database "$(DB_URL)" -path $(MIGRATION_DIR) version

# マイグレーションのリセット（注意: 全てダウンした後にアップ）
db-migration-reset:
	$(MIGRATE_CMD) -database "$(DB_URL)" -path $(MIGRATION_DIR) down
	$(MIGRATE_CMD) -database "$(DB_URL)" -path $(MIGRATION_DIR) up

server-reset:
	make docker-restart
	sleep 3
	make db-migration-up

server-start:
	go run cmd/main.go server start

server-reset-start:
	make server-reset
	make server-start

docker-restart:
	docker compose down
	docker compose up -d

setup-tools:
	sh scripts/setup-tools.sh

PROTO_DIR = ./api/proto
VENDOR_DIR = ${PROTO_DIR}/vendor
JS_PROTO_OUT_DIR = ./public/lib/proto

proto:
	buf dep update
	PATH="${PWD}/bin:$$PATH" buf generate
