package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rainbow777/todolist/api/middlewares"
	"github.com/rainbow777/todolist/controllers"
	"github.com/rainbow777/todolist/repository"
	"github.com/rainbow777/todolist/services"
)

// muxルーターの生成とルーティング
func NewRouter(db *sql.DB) *mux.Router {
	// 依存関係を注入（DI）レポジトリ→サービス→コントローラー層

	repository := repository.NewMyAppRepository(db)
	service := services.NewMyAppService(repository)
	controller := controllers.NewMyAppController(service)

	// muxルーターを生成
	r := mux.NewRouter()

	// ルーティング登録
	r.HandleFunc("/todo/insert", controller.InsertTaskHandler).Methods(http.MethodPost)
	r.HandleFunc("/todo/gettask/{id:[0-9]+}", controller.GetTaskHandler).Methods(http.MethodGet)
	r.HandleFunc("/todo/getlist", controller.GetListHandler).Methods(http.MethodGet)
	r.HandleFunc("/todo/update/{id:[0-9]+}", controller.UpdateTaskHandler).Methods(http.MethodPatch)
	r.HandleFunc("/todo/delete/{id:[0-9]+}", controller.DeleteTaskHandler).Methods(http.MethodDelete)

	// ミドルウェアを登録（登録順に実行される）
	r.Use(middlewares.LoggingMiddleware) // ロギング
	r.Use(middlewares.AuthHandle)        // ユーザー認証

	return r
}
