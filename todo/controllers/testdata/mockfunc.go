package testdata

import (
	"errors"

	"github.com/rainbow777/todolist/structure"
)

type ServiceMock struct {
	InsertTaskServiceFunc func(structure.InsertData) (structure.Todo, error)
	GetTaskServiceFunc    func(int, string) (structure.Todo, error)
	GetListServiceFunc    func(structure.GetListRequest) ([]structure.Todo, error)
	UpdateTaskServiceFunc func(*structure.UpdateData) (structure.Todo, error)
}

func NewServiceMock() *ServiceMock {
	return &ServiceMock{}
}

func (s *ServiceMock) InsertTaskService(insertData structure.InsertData) (structure.Todo, error) {
	if s.InsertTaskServiceFunc != nil {
		return s.InsertTaskServiceFunc(insertData)
	}
	return structure.Todo{}, errors.New("InsertTaskServiceFunc isn't set up")
}

func (s *ServiceMock) GetTaskService(taskID int, username string) (structure.Todo, error) {

	if s.GetTaskServiceFunc != nil {
		return s.GetTaskServiceFunc(taskID, username)
	}

	return structure.Todo{}, errors.New("GetTaskServiceFunc isn't set up")
}

func (s *ServiceMock) GetListService(getListRequest structure.GetListRequest) ([]structure.Todo, error) {

	if s.GetListServiceFunc != nil {
		return s.GetListServiceFunc(getListRequest)
	}

	return []structure.Todo{}, errors.New("GetListServiceFunc isn't set up")
}

func (s *ServiceMock) UpdateTaskService(updateData *structure.UpdateData) (structure.Todo, error) {

	if s.UpdateTaskServiceFunc != nil {
		return s.UpdateTaskServiceFunc(updateData)
	}

	return structure.Todo{}, errors.New("UpdateTaskServiceFunc isn't set up")
}

func (s *ServiceMock) DeleteTaskService(taskID int, username string) error {
	return nil
}
