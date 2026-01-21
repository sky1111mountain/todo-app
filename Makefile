.PHONY: up build down test logs db

SERVICE_APP = app
SERVICE_DB = db

# docker compose up を短縮
up:
	sudo docker compose up -d

# ビルドして起動
build:
	sudo docker compose up -d --build

# 停止
down:
	sudo docker compose down

# 【本命】テスト実行
test-local:
	DB_USER=rainbow777 \
	DB_PASS=klehorha \
	TEST_DB_NAME=test_db \
	TEST_DB_HOST=127.0.0.1 \
	TEST_DB_PORT=3307 \
	go test -v ./controllers/... ./services/... ./repository/...

test:
	go test -v ./...

test-repo:
	go test -v ./repository/...

test-svc:
	go test -v ./services/...

test-cr:
	go test -v ./controllers/...

# DBのログを確認
logs:
	sudo docker compose logs -f $(SERVICE_APP)

# DBの中身を覗く（パスワード入力が必要）
db:
	sudo docker exec -it $(SERVICE_DB) mysql -u rainbow777 -p tododb