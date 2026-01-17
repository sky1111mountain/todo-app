package repository_test

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func prepareTestDB(t *testing.T, filename string) {
	cleanupDB(t)

	setupTestData(t, filename)

	t.Cleanup(func() {
		cleanupDB(t)
	})
}

func cleanupDB(t *testing.T) {
	sqlCommands := readSqlFile(t, "cleanupDB.sql")
	executeSql(t, sqlCommands)
}

func setupTestData(t *testing.T, filename string) {
	sqlCommands := readSqlFile(t, filename)
	executeSql(t, sqlCommands)
}

func readSqlFile(t *testing.T, filename string) []string {
	filepath := fmt.Sprintf("./testdata/%s", filename)
	sqlScript, err := os.ReadFile(filepath)
	if err != nil {
		t.Fatalf("failed to read %s: %v", filename, err)
	}

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
