
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
FROM golang:1.25-alpine
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# バイナリだけではなく、builderステージの /app（ソース一式）をまるごとコピーする
COPY --from=builder /app .

# アプリがポート8080（仮）を使う場合
EXPOSE 8080

# 実行
CMD ["./todo-app"]