package repository_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	_ "embed"

	_ "github.com/go-sql-driver/mysql"
)

func TestMain(m *testing.M) {

	setDB()

	m.Run()

	teardown()
}

var (
	dbUser     string
	dbPass     string
	dbHost     string
	testDBname string
	dbconn     string
	testDB     *sql.DB
)

func setDB() {
	getEnv()

	connectDB()
}

func teardown() {
	testDB.Close()
}

func getEnv() {
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbHost = os.Getenv("TEST_DB_HOST")
	testDBname = os.Getenv("TEST_DB_NAME")

	if dbUser == "" || dbPass == "" || testDBname == "" {
		log.Fatalf("missing required environment variables : dbUser, dbPass, testDBname")
	}
}

func connectDB() {
	var err error

	// Dockerネットワーク内では、サービス名(dbHost)と内部ポート(3306)を使用して接続する
	dbconn = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPass, dbHost, testDBname)
	testDB, err = sql.Open("mysql", dbconn)
	if err != nil {
		log.Fatalf("fail to connect DB %v", err)
	}

	if err = testDB.Ping(); err != nil {
		log.Fatalf("fail to ping DB %v", err)
	}
}

func prepareTestDB(t *testing.T, filename string) {
	cleanupDB(t)

	setupTestData(t, filename)

	t.Cleanup(func() {
		cleanupDB(t)
	})
}

//go:embed testdata/cleanupDB.sql
var cleanupDBSQL string

//go:embed testdata/insertTestData.sql
var insertTestDataSQL string

//go:embed testdata/gettaskTestData.sql
var gettaskTestDataSQL string

//go:embed testdata/getlistTestData.sql
var getlistTestDataSQL string

//go:embed testdata/updataTestData.sql
var updataTestDataSQL string

//go:embed testdata/deleteTestData.sql
var deleteTestDataSQL string

func cleanupDB(t *testing.T) {
	sqlCommands := readSqlFile(t, "cleanupDB.sql")
	executeSql(t, sqlCommands)
}

func setupTestData(t *testing.T, filename string) {
	sqlCommands := readSqlFile(t, filename)
	executeSql(t, sqlCommands)
}

func getSqlContent(t *testing.T, filename string) string {
	switch filename {
	case "cleanupDB.sql":
		return cleanupDBSQL
	case "deleteTestData.sql":
		return deleteTestDataSQL
	case "getlistTestData.sql":
		return getlistTestDataSQL
	case "gettaskTestData.sql":
		return gettaskTestDataSQL
	case "insertTestData.sql":
		return insertTestDataSQL
	case "updataTestData.sql":
		return updataTestDataSQL
	default:
		t.Fatalf("Unknown SQL file: %s", filename)
		return ""
	}
}

func readSqlFile(t *testing.T, filename string) []string {

	sqlScript := getSqlContent(t, filename)

	return strings.Split(string(sqlScript), ";")
}

func executeSql(t *testing.T, commands []string) {
	for _, command := range commands {
		command = strings.TrimSpace(command)
		if command == "" {
			continue
		}

		_, err := testDB.Exec(command)
		if err != nil {
			t.Fatalf("failed to execute sql: %s err: %v", command, err)
		}
	}
}
