package api

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/rainbow777/todolist/api/middlewares"
	"github.com/rainbow777/todolist/controllers"
	"github.com/rainbow777/todolist/repository"
	"github.com/rainbow777/todolist/services"
)

// muxルーターの生成とルーティング
func NewRouter(db *sql.DB) *chi.Mux {

	// 依存関係を注入（DI）レポジトリ→サービス→コントローラー層
	repository := repository.NewMyAppRepository(db)
	service := services.NewMyAppService(repository)
	controller := controllers.NewMyAppController(service)

	r := chi.NewRouter()

	r.Use(middlewares.LoggingMiddleware) // ロギング

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/*", fs)

	r.Route("/todo", func(r chi.Router) {
		r.Use(middlewares.AuthHandle) // このブロック内だけに認証を適用！

		r.Post("/insert", controller.InsertTaskHandler)
		r.Get("/gettask/{id:[0-9]+}", controller.GetTaskHandler)
		r.Get("/getlist", controller.GetListHandler)
		r.Patch("/update/{id:[0-9]+}", controller.UpdateTaskHandler)
		r.Delete("/delete/{id:[0-9]+}", controller.DeleteTaskHandler)
	})

	return r
}
