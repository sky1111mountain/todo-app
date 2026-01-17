package repository_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rainbow777/todolist/envconfig"
)

func TestMain(m *testing.M) {
	// 環境変数ファイル(.env)の読み込み
	envconfig.LoadEnvConfig()

	setDB()

	m.Run()

	teardown()
}
