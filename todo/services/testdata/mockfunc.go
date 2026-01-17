package testdata

import (
	"database/sql"
	"errors"

	"github.com/rainbow777/todolist/structure"
)

type RepositoryMock struct {
	DB                   *sql.DB
	InsertTaskDBClosure  func(insertData structure.InsertData) (structure.Todo, error)
	GetTaskDBClosure     func(taskID int, username string) (structure.Todo, error)
	GetTodolistDBClosure func(getListRequest structure.GetListRequest) ([]structure.Todo, error)
	UpdateTodoDBClosure  func(updateData *structure.UpdateData) (structure.Todo, error)
	DeleteTaskDBClosure  func(taskID int, username string) error
}

func NewRepositoryMock(db *sql.DB) *RepositoryMock {
	return &RepositoryMock{DB: db}
}

func (r *RepositoryMock) InsertTaskDB(insertData structure.InsertData) (structure.Todo, error) {
	if r.InsertTaskDBClosure != nil {
		return r.InsertTaskDBClosure(insertData)
	} else {
		return structure.Todo{}, errors.New("InsertTaskDBClosure isn't set up")
	}
}

func (r *RepositoryMock) GetTaskDB(taskID int, username string) (structure.Todo, error) {
	if r.GetTaskDBClosure != nil {
		return r.GetTaskDBClosure(taskID, username)
	} else {
		return structure.Todo{}, errors.New("GetTaskDBClosure isn't set up")
	}
}

func (r *RepositoryMock) GetTodolistDB(getListRequest structure.GetListRequest) ([]structure.Todo, error) {
	if r.GetTodolistDBClosure != nil {
		return r.GetTodolistDBClosure(getListRequest)
	} else {
		return []structure.Todo{}, errors.New("GetTodolistDBClosure isn't set up")
	}
}

func (r *RepositoryMock) UpdateTaskDB(updateData *structure.UpdateData) (structure.Todo, error) {
	if r.UpdateTodoDBClosure != nil {
		return r.UpdateTodoDBClosure(updateData)
	} else {
		return structure.Todo{}, errors.New("UpdateTodoDBClosure isn't set up")
	}
}

func (r *RepositoryMock) DeleteTaskDB(taskID int, username string) error {
	if r.DeleteTaskDBClosure != nil {
		return r.DeleteTaskDBClosure(taskID, username)
	} else {
		return errors.New("DeleteTaskDBClosure isn't set up")
	}
}
