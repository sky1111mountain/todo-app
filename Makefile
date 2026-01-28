.PHONY: up build down test logs db testdb

-include .env

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

# テストDBの起動
testdb:
	sudo docker compose up test_db

# テスト用の環境変数
TEST_ENV = 	DB_USER=$(DB_USER) \
			DB_PASS=$(DB_PASS) \
			TEST_DB_NAME=$(TEST_DB_NAME) \
			TEST_DB_HOST=$(TEST_DB_HOST) \
			TEST_DB_PORT=$(TEST_DB_PORT)

# テスト実行 「make testdb」でテストDBを起動してから実行してください
test:
	$(TEST_ENV) go test -v -count=1 ./controllers/... ./services/... ./repository/...

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
	sudo docker exec -it $(SERVICE_DB) mysql -u $(DB_USER) -p tododb