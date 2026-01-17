package services_test

import (
	"testing"

	"github.com/rainbow777/todolist/services"
	"github.com/rainbow777/todolist/services/testdata"
)

func TestInsertTaskService(t *testing.T) {

	for _, test := range testdata.InsertTaskTestCases {
		t.Run(test.TestName, func(t *testing.T) {
			repo := MakeRepoMockInsert(test)

			ser := services.NewMyAppService(repo)

			returnTask, err := ser.InsertTaskService(test.InsertTask)

			CheckReturnedDataInsertTask(t, test, returnTask, err)
		})
	}

}

func TestGetTaskService(t *testing.T) {

	for _, test := range testdata.GetTaskTestCases {
		t.Run(test.TestName, func(t *testing.T) {
			repoMock := MakeRepoMockGetTask(test)

			ser := services.NewMyAppService(repoMock)
			returnedTask, err := ser.GetTaskService(test.RequestID, test.UserName)

			CheckReturnedDataGetTask(t, test, returnedTask, err)
		})
	}
}

func TestGetListService(t *testing.T) {

	for _, test := range testdata.GetListTestCases {
		t.Run(test.TestName, func(t *testing.T) {

			repoMock := MakeRepoMockGetList(test)
			ser := services.NewMyAppService(repoMock)

			returnedList, err := ser.GetListService(test.Request)

			CheckReturnedDataGetList(t, test, returnedList, err)

		})
	}
}

func TestUpdateTaskService(t *testing.T) {

	for _, test := range testdata.UpdateTaskTestCases {
		t.Run(test.TestName, func(t *testing.T) {

			repoMock := MakeRepoMockUpdate(test)
			ser := services.NewMyAppService(repoMock)

			returnedData, err := ser.UpdateTaskService(&test.UpdateData)

			CheckReturnedDataUpdate(t, test, returnedData, err)
		})
	}
}

func TestDeleteTaskService(t *testing.T) {

	for _, test := range testdata.DleteTaskTestCases {
		t.Run(test.TestName, func(t *testing.T) {
			repoMock := MakeRepoMockDelete(test)
			ser := services.NewMyAppService(repoMock)

			err := ser.DeleteTaskService(test.TaskID, test.AuthUserName)

			CheckReturnedDataDelete(t, test, err)
		})
	}
}
