package repository_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMain(m *testing.M) {

	setDB()

	m.Run()

	teardown()
}
