package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// GetDB はデータベースに接続し、*sql.DBを返す
func GetDB(user, pass, host, dbname string) (*sql.DB, error) {

	// 環境変数が空でないことを確認するバリデーション
	if user == "" || pass == "" || dbname == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	// データベース接続文字列の生成
	dbconn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", user, pass, host, dbname)

	// データベース接続の確立
	db, err := sql.Open("mysql", dbconn)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect database:%w", err)
	}

	var pingErr error
	for i := 0; i < 10; i++ {
		pingErr = db.Ping()
		if pingErr == nil {
			return db, nil
		}
		log.Printf("DB not ready... retrying in 2 seconds (attempt %d/10): %v", i+1, pingErr)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("could not connect to DB after retries: %w", pingErr)
}
