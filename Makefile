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
test:
	sudo docker compose exec $(SERVICE_APP) go test -v ./...

test-repo:
	sudo docker compose exec $(SERVICE_APP) go test -v ./repository/...

test-svc:
	sudo docker compose exec $(SERVICE_APP) go test -v ./services/...

test-cr:
	ssudo docker compose exec $(SERVICE_APP) go test -v ./controllers/...

# DBのログを確認
logs:
	sudo docker compose logs -f $(SERVICE_APP)

# DBの中身を覗く（パスワード入力が必要）
db:
	sudo docker exec -it $(SERVICE_DB) mysql -u rainbow777 -p tododb