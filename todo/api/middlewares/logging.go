package middlewares

import (
	"log"
	"net/http"
)

type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

func NewMyResposeWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

// ハンドラが実行されると、WriteHeadeメソッドが自動で実行されて、ステータスコードがcodeフィールドに格納される
func (rlw *resLoggingWriter) WriteHeader(code int) {
	// レシーバーをポインタにしないとコピー先に値が格納されてしまう

	rlw.code = code
	rlw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		log.Println(req.RequestURI, req.Method)

		// レスポンス内容を確認するために自作レスポンスを作成
		rlw := NewMyResposeWriter(w)

		// 自作レスポンスを渡して、次のハンドラを実行
		next.ServeHTTP(rlw, req)

		log.Printf("response code:%d", rlw.code)
	})
}
