package services_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/rainbow777/todolist/myerrors"
	"github.com/rainbow777/todolist/services/testdata"
	"github.com/rainbow777/todolist/structure"
	"github.com/stretchr/testify/assert"
)

func MakeRepoMockInsert(test testdata.InsertTaskServiceData) *testdata.RepositoryMock {
	var db *sql.DB
	db = nil
	repo := testdata.NewRepositoryMock(db)
	repo.InsertTaskDBClosure = func(insertData structure.InsertData) (structure.Todo, error) {
		if test.ExpectedErr.ErrCode != "" {
			return structure.Todo{}, &test.ExpectedErr
		}
		return test.ExpectedTask, nil
	}
	return repo
}

func CheckReturnedDataInsertTask(t *testing.T, test testdata.InsertTaskServiceData, returnedTask structure.Todo, err error) {
	if err == nil {
		assert.Equal(t, test.ExpectedTask.ID, returnedTask.ID, "TaskID is unmatched")
		assert.Equal(t, test.ExpectedTask.Task, returnedTask.Task, "Task content is unmatched")
		assert.Equal(t, test.ExpectedTask.Priority, returnedTask.Priority, "Priority is unmatched")
		assert.Equal(t, test.ExpectedTask.Status, returnedTask.Status, "Status is unmatched")
		assert.Equal(t, test.ExpectedTask.UserName, returnedTask.UserName, "UserName is unmatched")
	} else {
		var returnedErr *myerrors.MyAppError
		if errors.As(err, &returnedErr) {
			assert.Equal(t, test.ExpectedErr.ErrCode, returnedErr.ErrCode, "Error code is unmatched")
			assert.Equal(t, test.ExpectedErr.Message, returnedErr.Message, "Error Message is unmatched")
		} else {
			t.Fatalf("Failed to change error to MyAppError %v", err)
		}
	}
}

func MakeRepoMockGetTask(test testdata.GetTaskTestData) *testdata.RepositoryMock {
	var db *sql.DB
	db = nil
	repoMock := testdata.NewRepositoryMock(db)
	repoMock.GetTaskDBClosure = func(taskID int, username string) (structure.Todo, error) {
		if test.ExpectedErr.ErrCode != "" {
			return structure.Todo{}, test.RepoReturnErr
		} else {
			return test.ExpectedTask, nil
		}
	}
	return repoMock
}

func CheckReturnedDataGetTask(t *testing.T, test testdata.GetTaskTestData, returnedTask structure.Todo, err error) {

	if err == nil {
		assert.Equal(t, test.ExpectedTask.ID, returnedTask.ID, "Task ID is unmatched")
		assert.Equal(t, test.ExpectedTask.Task, returnedTask.Task, "Task content is unmatched")
		assert.Equal(t, test.ExpectedTask.Priority, returnedTask.Priority, "Priority is unmatched")
		assert.Equal(t, test.ExpectedTask.Status, returnedTask.Status, "Status is unmatched")
		assert.Equal(t, test.ExpectedTask.UserName, returnedTask.UserName, "UserName is unmatched")
	} else {
		var returnedErr *myerrors.MyAppError
		if errors.As(err, &returnedErr) {
			assert.Equal(t, test.ExpectedErr.ErrCode, returnedErr.ErrCode, "ErrCode is unmatched")
			assert.Equal(t, test.ExpectedErr.Message, returnedErr.Message, "Err Message is unmatched")
			assert.Equal(t, test.ExpectedErr.Err.Error(), returnedErr.Err.Error(), "Err Message is unmatched")
		} else {
			t.Fatalf("Failed to change error to MyAppError %v", err)
		}
	}
}

func MakeRepoMockGetList(test testdata.GetListTestData) *testdata.RepositoryMock {
	var db *sql.DB
	db = nil
	repoMock := testdata.NewRepositoryMock(db)
	repoMock.GetTodolistDBClosure = func(getListRequest structure.GetListRequest) ([]structure.Todo, error) {
		if test.ExpectedErr.ErrCode != "" {
			return nil, test.RepoReturnErr
		} else {
			return test.ExpectedTask, nil
		}
	}
	return repoMock
}

func CheckReturnedDataGetList(t *testing.T, test testdata.GetListTestData,
	returnedList []structure.Todo, err error) {

	if err == nil {
		for i, returnedTask := range returnedList {
			assert.Equal(t, test.ExpectedTask[i].ID, returnedTask.ID, "Task ID is unmatched")
			assert.Equal(t, test.ExpectedTask[i].Task, returnedTask.Task, "Task content is unmatched")
			assert.Equal(t, test.ExpectedTask[i].Priority, returnedTask.Priority, "Priority is unmatched")
			assert.Equal(t, test.ExpectedTask[i].Status, returnedTask.Status, "Status is unmatched")
		}
	} else {
		var returnedErr *myerrors.MyAppError
		if errors.As(err, &returnedErr) {
			assert.Equal(t, test.ExpectedErr.ErrCode, returnedErr.ErrCode, "ErrCode is unmatched")
			assert.Equal(t, test.ExpectedErr.Message, returnedErr.Message, "ErrMessage is unmatched")
			assert.Equal(t, test.ExpectedErr.Err, returnedErr.Err, "wrappedErr is unmatched")
		} else {
			t.Fatalf("Failed to change error to MyAppError %v", err)
		}
	}

}

func MakeRepoMockUpdate(test testdata.UpdateTaskTestData) *testdata.RepositoryMock {
	var db *sql.DB = nil
	repoMock := testdata.NewRepositoryMock(db)
	repoMock.UpdateTodoDBClosure = func(updateData *structure.UpdateData) (structure.Todo, error) {
		if test.RepoReturnErr != nil {
			return structure.Todo{}, test.RepoReturnErr
		} else {
			return test.RepoReturnTask, nil
		}
	}
	return repoMock
}

func CheckReturnedDataUpdate(t *testing.T, test testdata.UpdateTaskTestData,
	returnedData structure.Todo, err error) {

	if err == nil {
		assert.Equal(t, test.ExpectedTask.ID, returnedData.ID, "TaskID is unmatched")
		assert.Equal(t, test.ExpectedTask.Task, returnedData.Task, "Task content is unmatched")
		assert.Equal(t, test.ExpectedTask.Priority, returnedData.Priority, "Priority is unmatched")
		assert.Equal(t, test.ExpectedTask.Status, returnedData.Status, "Status is unmatched")
		assert.Equal(t, test.ExpectedTask.UserName, returnedData.UserName, "UserName is unmatched")
	} else {
		var returnedErr *myerrors.MyAppError
		if errors.As(err, &returnedErr) {
			assert.Equal(t, test.ExpectedErr.ErrCode, returnedErr.ErrCode, "ErrCode is unmatched")
			assert.Equal(t, test.ExpectedErr.Message, returnedErr.Message, "Message is unmatched")
			assert.Equal(t, test.ExpectedErr.Err.Error(), returnedErr.Err.Error(), "WrappedErr is unmatched")
		} else {
			t.Fatalf("Failed to change error to MyAppError %v", err)
		}
	}

}

func MakeRepoMockDelete(test testdata.DeleteTaskTestData) *testdata.RepositoryMock {
	var db *sql.DB = nil
	repoMock := testdata.NewRepositoryMock(db)
	repoMock.DeleteTaskDBClosure = func(taskID int, username string) error {
		if test.RepoReturnErr != nil {
			return test.RepoReturnErr
		} else {
			return nil
		}
	}
	return repoMock
}

func CheckReturnedDataDelete(t *testing.T, test testdata.DeleteTaskTestData, err error) {

	if err != nil {
		var returnedErr *myerrors.MyAppError
		if errors.As(err, &returnedErr) {
			assert.Equal(t, test.ExpectedErr.ErrCode, returnedErr.ErrCode, "ErrCode is unmatched")
			assert.Equal(t, test.ExpectedErr.Message, returnedErr.Message, "Message is unmatched")
			assert.Equal(t, test.ExpectedErr.Err, returnedErr.Err, "WrappedErr is unmatched")
		} else {
			t.Fatalf("Failed to change error to MyAppError %v", err)
		}
	}

}
