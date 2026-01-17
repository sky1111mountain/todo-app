package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rainbow777/todolist/api"
	"github.com/rainbow777/todolist/database"
	"github.com/rainbow777/todolist/envconfig"
)

func main() {
	// 環境変数ファイル(.env)の読み込み
	envconfig.LoadEnvConfig()

	// データベース接続
	db, err := database.GetDB(
		envconfig.AppConfig.DBUser,
		envconfig.AppConfig.DBPass,
		envconfig.AppConfig.DBHost,
		envconfig.AppConfig.DBName,
	)
	if err != nil {
		log.Fatalf("Error occurred in database.GetDB: %v", err)
	}
	defer db.Close()

	err = database.MakeTable(db)
	if err != nil {
		log.Fatalf("Failed to make todolist table %v", err)
	}

	// ルータの生成とルーティング
	r := api.NewRouter(db)

	log.Println("todo server start at port 8080")
	// サーバー起動
	log.Fatal(http.ListenAndServe(":8080", r))
}
