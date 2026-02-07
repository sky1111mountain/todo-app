
# --- Stage 1: ビルド用 ---
FROM golang:1.25-alpine AS builder

# 作業ディレクトリ
WORKDIR /app

# 先に依存関係をコピーしてダウンロード
COPY go.mod go.sum ./
RUN go mod download

# すべてのソースコードをコピー
COPY . .

# main.go がルートにあるので "." でビルド可能
# 静的バイナリとしてコンパイル
RUN CGO_ENABLED=0 GOOS=linux go build -o todo-app .

# --- Stage 2: 実行用 ---
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app

# 
COPY --from=builder /app/todo-app .
# static フォルダもコピーが必要！
COPY --from=builder /app/static ./static

# アプリがポート8080（仮）を使う場合
EXPOSE 8080

# 実行
CMD ["./todo-app"]