package repository_test

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rainbow777/todolist/repository"
	"github.com/rainbow777/todolist/repository/testdata"
	"github.com/rainbow777/todolist/structure"
	"github.com/stretchr/testify/assert"
)

// 以下、Insert系
func TestInsertTaskDB(t *testing.T) {
	prepareTestDB(t, "insertTestData.sql")

	for _, test := range testdata.InsertTaskTestCases {
		t.Run(test.TestName, func(t *testing.T) {
			repo := repository.NewMyAppRepository(testDB)
			returnedTask, err := repo.InsertTaskDB(test.InsertTask)
			if err == nil {
				assert.Equal(t, test.ExpectedTask.ID, returnedTask.ID, "TaskID is mismatched")
				assert.Equal(t, test.ExpectedTask.Task, returnedTask.Task, "Task Content is mismatched")
				assert.Equal(t, test.ExpectedTask.Priority, returnedTask.Priority, "Priority is mismatched")
				assert.Equal(t, test.ExpectedTask.Status, returnedTask.Status, "Status is mismatched")
				assert.Equal(t, test.ExpectedTask.UserName, returnedTask.UserName, "UserName is mismatched")
			} else {
				t.Errorf("Unexpected Error from InsertTaskDB %v", err)
			}
		})
	}
}

func TestInsertTaskDB_ExecError(t *testing.T) {
	t.Run("Case:Error occured in Exec in InsertTaskDB", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("Failed to create sqlmock: %v", err)
		}
		defer db.Close()

		mock.ExpectExec("insert into todolist").
			WillReturnError(errors.New("SIMULATED: OrizinalErr from DB.Exec"))

		repo := repository.NewMyAppRepository(db)

		_, err = repo.InsertTaskDB(testdata.TestTaskForInsertErr)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "R001")
		assert.Contains(t, err.Error(), "Error occurred in db.Exec in InsertTaskDB")
		assert.Contains(t, err.Error(), "SIMULATED: OrizinalErr from DB.Exec")
	})

}

func TestInsertTaskDB_RowsAffectedErr(t *testing.T) {
	t.Run("Case:Error occured in result.RowsAffected", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("Failed to create sqlmock: %v", err)
		}
		defer db.Close()

		mock.ExpectExec("insert into todolist").
			WillReturnResult(sqlmock.NewErrorResult(errors.New("SIMULATED: RowAffected Err")))

		repo := repository.NewMyAppRepository(db)

		_, err = repo.InsertTaskDB(testdata.TestTaskForInsertErr)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "R002")
		assert.Contains(t, err.Error(), "Failed to check database changes were complete")
		assert.Contains(t, err.Error(), "SIMULATED: RowAffected Err")
	})
}

func TestInsertTaskDB_NoRowsAffected(t *testing.T) {
	t.Run("Case:DB RowsAffected returns 0", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("Failed to create sqlmock: %v", err)
		}
		defer db.Close()

		mock.ExpectExec("insert into todolist").WillReturnResult(sqlmock.NewResult(0, 0))

		repo := repository.NewMyAppRepository(db)

		_, err = repo.InsertTaskDB(testdata.TestTaskForInsertErr)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "R009")
		assert.Contains(t, err.Error(), "Database has not changed")
		assert.Contains(t, err.Error(), "No affectedRows")
	})
}

func TestGetInsertedTask_LastInsertIdError(t *testing.T) {
	t.Run("Case:LastInsertId returns Error in GetInsertedTask in InsertTaskDB", func(t *testing.T) {
		db, _, err := sqlmock.New()
		if err != nil {
			t.Fatalf("Failed to create sqlmock: %v", err)
		}
		defer db.Close()

		repo := repository.NewMyAppRepository(db)
		result := sqlmock.NewErrorResult(errors.New("SIMULATED:LastInsertID Error"))

		_, err = repository.GetInsertedTask(repo, result, "rainbow777")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "R001")
		assert.Contains(t, err.Error(), "Error occurred in LastInsertId")
		assert.Contains(t, err.Error(), "SIMULATED:LastInsertID Error")
	})
}

func TestGetInsertedTask_FailedGetInsertedTask(t *testing.T) {
	t.Run("Case:LastInsertId returns Error in GetInsertedTask in InsertTaskDB", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := repository.NewMyAppRepository(db)

		mock.ExpectExec("SELECT").WillReturnError(errors.New("SIMULATED: Error occurred in GetTaskDB"))
		result := sqlmock.NewResult(1, 1)

		_, err = repository.GetInsertedTask(repo, result, "rainbow777")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "R004")
		assert.Contains(t, err.Error(), "Error occurred in row.Scan in GetTaskDB")
		assert.Contains(t, err.Error(), "SIMULATED: Error occurred in GetTaskDB")
	})
}

func TestGetInsertedTask_ReturnNoRows(t *testing.T) {
	t.Run("Case:GetTaskDB returns NoRows", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec("insert into todolist").WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectQuery("SELECT").
			WillReturnRows(sqlmock.NewRows(
				[]string{"id", "task", "priority", "status", "username", "created_at"}))

		repo := repository.NewMyAppRepository(db)

		_, err = repo.InsertTaskDB(testdata.TestTaskForInsertErr)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "R003")
		assert.Contains(t, err.Error(), "Failed to get last inserted task")
		assert.Contains(t, err.Error(), "R004")
		assert.Contains(t, err.Error(), "Error occurred in row.Scan in GetTaskDB")
		assert.Contains(t, err.Error(), "sql: no rows in result set")
	})
}

// Insert系テストここまで------------------------------------------------------------------------------------

// 以下、GetTask系テスト
func TestGetTaskDB(t *testing.T) {
	prepareTestDB(t, "gettaskTestData.sql")

	for _, test := range testdata.GetTaskTestCases {
		t.Run(test.TestName, func(t *testing.T) {
			repo := repository.NewMyAppRepository(testDB)
			got, err := repo.GetTaskDB(test.TaskID, test.AuthUserName)
			assert.NoError(t, err)
			assert.Equal(t, test.ExpectedTask.ID, got.ID, "TaskID is mismatched")
			assert.Equal(t, test.ExpectedTask.Task, got.Task, "Task content is mismatched")
			assert.Equal(t, test.ExpectedTask.Priority, got.Priority, "Priority is mismatched")
			assert.Equal(t, test.ExpectedTask.UserName, got.UserName, "UserName is mismatched")
		})
	}

}

func TestGetTaskDB_ErrQueryRow(t *testing.T) {
	t.Run("Case:Error occurred in QueryRow", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery("select").WillReturnError(errors.New("SIMULATED: QueryRow Error"))
		repo := repository.NewMyAppRepository(db)
		_, err = repo.GetTaskDB(1, "rainbow777")
		assert.Error(t, err)
		// assert.Contains(t, err.Error(), "SIMULATED: QueryRow Error")
	})
}

func TestGetTaskDB_ErrorScanRows(t *testing.T) {
	t.Run("Case: ScanRows Error in GetTaskDB", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery("select").
			WillReturnRows(sqlmock.NewRows([]string{"id", "task", "priority", "status", "rainbow777"}))

		repo := repository.NewMyAppRepository(db)
		_, err = repo.GetTaskDB(1, "rainbow777")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "R004")
		assert.Contains(t, err.Error(), "Error occurred in row.Scan in GetTaskDB")
		// assert.Contains(t, err.Error(), "sql: no rows in result set")
	})
}

func TestGetTaskDB_Error_NoData(t *testing.T) {
	t.Run("Case: NoData", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery("SELECT").
			WillReturnError(sql.ErrNoRows)

		repo := repository.NewMyAppRepository(db)
		_, err = repo.GetTaskDB(999, "rainbow777")
		assert.Error(t, err)
		fmt.Println(err)
		fmt.Println(err.Error())
		assert.Contains(t, err.Error(), "R004")
		assert.Contains(t, err.Error(), "Error occurred in row.Scan in GetTaskDB")
		assert.Contains(t, err.Error(), "sql: no rows in result set")
	})
}

// GetTask系ここまで-------------------------------------------------------------------------------------------

// 以下、GetList系テスト
func TestGetTodolistDB(t *testing.T) {
	prepareTestDB(t, "getlistTestData.sql")
	for _, test := range testdata.GetListTestCases {
		t.Run(test.TestName, func(t *testing.T) {
			repo := repository.NewMyAppRepository(testDB)
			list, err := repo.GetTodolistDB(test.Request)
			assert.NoError(t, err)
			for i, task := range list {
				assert.Equal(t, test.ExpectedList[i].ID, task.ID, "TaskID is unmatched")
				assert.Equal(t, test.ExpectedList[i].Task, task.Task, "Task content is unmatched")
				assert.Equal(t, test.ExpectedList[i].Priority, task.Priority, "Priority is unmatched")
				assert.Equal(t, test.ExpectedList[i].Status, task.Status, "Status is unmatched")
				assert.Equal(t, test.ExpectedList[i].UserName, task.UserName, "UserName is unmatched")
			}
		})
	}
}

func TestGetTodolistDB_Error_GetRows(t *testing.T) {
	requestSlice := []structure.GetListRequest{
		{
			AuthUserName: "rainbow777",
		},
		{
			AuthUserName: "rainbow777",
			Priorities:   []string{"high"},
		},
		{
			AuthUserName: "rainbow777",
			Priorities:   []string{"medium"},
		},
		{
			AuthUserName: "rainbow777",
			Priorities:   []string{"low"},
		},
	}
	for _, request := range requestSlice {
		t.Run("Case: Error Occurred in GetRows", func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)

			mock.ExpectQuery("SELECT").WillReturnError(errors.New("SIMULATED: db.Query Error"))

			repo := repository.NewMyAppRepository(db)

			_, err = repo.GetTodolistDB(request)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "R005")
			assert.Contains(t, err.Error(), "Error occured in db.Query in GetRows func")
			assert.Contains(t, err.Error(), "SIMULATED: db.Query Error")
		})
	}
}

func TestGetTodolistDB_Error_ScanRows(t *testing.T) {
	t.Run("Case: Error Occurred in ScanRows", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery("SELECT").
			WillReturnRows(
				// int型のコラムidにstring型の値を代入してscanエラーを起こす
				sqlmock.NewRows([]string{"id", "task", "priority", "status", "username", "created_at"}).
					AddRow("invalid_id", "task1", "high", "done", "rainbow777", time.Now()))

		repo := repository.NewMyAppRepository(db)

		_, err = repo.GetTodolistDB(structure.GetListRequest{AuthUserName: "rainbow777"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "R004")
		assert.Contains(t, err.Error(), "Error occurred in rows.Scan in ScanRows")
		assert.Contains(t, err.Error(), " sql: Scan error on column index 0")
	})
}

// GetList系テストここまで----------------------------------------------------------------------------------

// 以下 Update系テスト
func TestUpdateTaskDB(t *testing.T) {
	prepareTestDB(t, "updataTestData.sql")

	for _, test := range testdata.UpdateTaskTestCases {
		t.Run(test.TestName, func(t *testing.T) {
			repo := repository.NewMyAppRepository(testDB)
			updatedTask, err := repo.UpdateTaskDB(&test.UpdateData)
			assert.NoError(t, err)
			assert.Equal(t, test.ExpectedTask.ID, updatedTask.ID, "TaskID is unmatched")
			assert.Equal(t, test.ExpectedTask.Task, updatedTask.Task, "Task content is unmatched")
			assert.Equal(t, test.ExpectedTask.Priority, updatedTask.Priority, "Priority is unmatched")
			assert.Equal(t, test.ExpectedTask.Status, updatedTask.Status, "Status is unmatched")
			assert.Equal(t, test.ExpectedTask.UserName, updatedTask.UserName, "UserName is unmatched")
		})
	}
}

func TestUpdateTaskDB_Error_Exec(t *testing.T) {
	t.Run("Case:Error occurred in Exec", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE todolist SET").
			WillReturnError(errors.New("SIMULATED: Exec Error"))

		repo := repository.NewMyAppRepository(db)

		_, err = repo.UpdateTaskDB(&testdata.UpdateData)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "R007")
		assert.Contains(t, err.Error(), "Error occurred in tx.Exec")
		assert.Contains(t, err.Error(), "SIMULATED: Exec Error")

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestUpdateTaskDB_Error_AffectedRows(t *testing.T) {
	t.Run("Case: Error occurred in RowsAffected", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE todolist SET").
			WillReturnResult(
				sqlmock.NewErrorResult(errors.New("SIMULATED: RowsAffected Error")))

		repo := repository.NewMyAppRepository(db)
		_, err = repo.UpdateTaskDB(&testdata.UpdateData)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "R002")
		assert.Contains(t, err.Error(), "Failed to check database changes were complete")
		assert.Contains(t, err.Error(), "SIMULATED: RowsAffected Error")

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestUpdateTaskDB_Error_UnchangeRows(t *testing.T) {
	t.Run("Case: Data is Unchanged", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE todolist SET").
			WillReturnResult(
				sqlmock.NewResult(0, 0))

		repo := repository.NewMyAppRepository(db)
		_, err = repo.UpdateTaskDB(&testdata.UpdateData)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "R009")
		assert.Contains(t, err.Error(), "Database has not changed")
		assert.Contains(t, err.Error(), "No affectedRows")

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestUpdateTaskDB_Error_GetUpdatedTask(t *testing.T) {
	t.Run("Case: Failed to Get updated task", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE todolist SET").
			WillReturnResult(sqlmock.NewResult(0, 1))

		mock.ExpectQuery("SELECT").
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "task", "priority", "status", "username", "created_at"}))

		repo := repository.NewMyAppRepository(db)
		_, err = repo.UpdateTaskDB(&testdata.UpdateData)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "R004")
		assert.Contains(t, err.Error(), "Error occurred in row.Scan")
		assert.Contains(t, err.Error(), "sql: no rows in result set")

	})
}

// Update系テストここまで-------------------------------------------------------------------------------------

func TestDeleteTaskDB(t *testing.T) {
	prepareTestDB(t, "deleteTestData.sql")

	for _, test := range testdata.DeleteTaskTestData {
		t.Run(test.TestName, func(t *testing.T) {

			repo := repository.NewMyAppRepository(testDB)
			err := repo.DeleteTaskDB(test.DataInDB.ID, "rainbow777")
			assert.NoError(t, err)

			getQuery := `select * from todolist where id = ?;`
			row := testDB.QueryRow(getQuery, test.DataInDB.ID)
			err = row.Scan()
			assert.Equal(t, sql.ErrNoRows, err, "The data hasn't been deleted")
		})
	}

}

func TestDeleteTaskDB_Error_Exec(t *testing.T) {
	t.Run("Case:Error occurred in Exec", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		mock.ExpectExec("delete from todolist").
			WillReturnError(errors.New("SIMULATED: Exec Error in DeleteTaskDB"))

		repo := repository.NewMyAppRepository(db)
		err = repo.DeleteTaskDB(1, "rainbow777")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "R008")
		assert.Contains(t, err.Error(), "Failed to delete task")
		assert.Contains(t, err.Error(), "SIMULATED: Exec Error in DeleteTaskDB")
	})
}

func TestTestDeleteTaskDB_Error_RowAffected(t *testing.T) {
	t.Run("Error occurred in RowAffected", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		mock.ExpectExec("delete from todolist").
			WillReturnResult(sqlmock.NewErrorResult(errors.New("SIMULATED: RowAffected Error")))

		repo := repository.NewMyAppRepository(db)
		err = repo.DeleteTaskDB(1, "rainbow777")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "R002")
		assert.Contains(t, err.Error(), "Failed to check database changes were complete")
		assert.Contains(t, err.Error(), "SIMULATED: RowAffected Error")
	})
}

func TestTestDeleteTaskDB_Error_UnchangeRow(t *testing.T) {
	t.Run("Case: The row is unchanged", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		mock.ExpectExec("delete from todolist").
			WillReturnResult(sqlmock.NewResult(0, 0))

		repo := repository.NewMyAppRepository(db)
		err = repo.DeleteTaskDB(1, "rainbow777")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "R009")
		assert.Contains(t, err.Error(), "Database has not changed")
		assert.Contains(t, err.Error(), "No affectedRows")
	})
}
