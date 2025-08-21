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

protoc-gen:
	protoc \
		--proto_path=$(PROTO_DIR) \
		--proto_path=$(VENDOR_DIR) \
		--go_out=${PROTO_DIR} \
		--go_opt=paths=source_relative \
		--go-grpc_out=${PROTO_DIR} \
		--go-grpc_opt=paths=source_relative \
		--go-resourcename_out=$(PROTO_DIR) --go-resourcename_opt=paths=source_relative \
		--grpc-gateway_out=${PROTO_DIR} \
		--grpc-gateway_opt=paths=source_relative \
		--grpc-web_out=import_style=typescript,mode=grpcwebtext:${JS_PROTO_OUT_DIR} \
		$(PROTO_DIR)/v1/*.proto

protoc-gen-ts:
	npx grpc_tools_node_protoc \
		--proto_path=./api/proto/vendor \
		--proto_path=./api/proto \
		--js_out=import_style=commonjs,binary:./public/lib/proto \
		--grpc-web_out=import_style=typescript,mode=grpcwebtext:./public/lib/proto \
		./api/proto/v1/*.proto \
		./api/proto/vendor/google/api/*.proto \
		./api/proto/vendor/google/type/*.proto

proto:
	make protoc-gen
	make protoc-gen-ts

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor
