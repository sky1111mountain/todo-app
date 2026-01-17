package repository_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

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
	dbHost = os.Getenv("TESTDBHOST")
	testDBname = os.Getenv("TESTDB")

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
