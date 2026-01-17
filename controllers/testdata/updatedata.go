package testdata

import (
	"net/http"

	"github.com/rainbow777/todolist/myerrors"
	"github.com/rainbow777/todolist/structure"
)

type UpdateTaskData struct {
	Name          string
	Code          int
	TaskID        int
	UpdateRequest structure.UpdateTaskRequest
	ReturnValue   structure.Todo
	ExpectedError myerrors.MyAppError
}

func makePointerStr(s string) *string { return &s }

var UpdateTaskTestCases = []UpdateTaskData{
	{
		Name:   "case1",
		Code:   http.StatusOK,
		TaskID: 1,
		UpdateRequest: structure.UpdateTaskRequest{
			Task: makePointerStr("Updated Task 1"),
		},
		ReturnValue: structure.Todo{
			ID:       1,
			Task:     "Updated Task 1",
			Priority: "high",
			Status:   "not_done",
			UserName: "rainbow777",
		},
	}, {
		Name:   "case2",
		Code:   http.StatusOK,
		TaskID: 1,
		UpdateRequest: structure.UpdateTaskRequest{
			Priority: makePointerStr("medium"),
		},
		ReturnValue: structure.Todo{
			ID:       1,
			Task:     "Updated Task 1",
			Priority: "medium",
			Status:   "not_done",
			UserName: "rainbow777",
		},
	}, {
		Name:   "case3",
		Code:   http.StatusOK,
		TaskID: 1,
		UpdateRequest: structure.UpdateTaskRequest{
			Priority: makePointerStr("low"),
		},
		ReturnValue: structure.Todo{
			ID:       1,
			Task:     "Updated Task 1",
			Priority: "low",
			Status:   "not_done",
			UserName: "rainbow777",
		},
	}, {
		Name:   "case4",
		Code:   http.StatusOK,
		TaskID: 1,
		UpdateRequest: structure.UpdateTaskRequest{
			Priority: makePointerStr("high"),
		},
		ReturnValue: structure.Todo{
			ID:       1,
			Task:     "Updated Task 1",
			Priority: "high",
			Status:   "not_done",
			UserName: "rainbow777",
		},
	}, {
		Name:   "case5",
		Code:   http.StatusOK,
		TaskID: 1,
		UpdateRequest: structure.UpdateTaskRequest{
			Status: makePointerStr("done"),
		}, ReturnValue: structure.Todo{
			ID:       1,
			Task:     "Updated Task 1",
			Priority: "high",
			Status:   "done",
			UserName: "rainbow777",
		},
	}, {
		Name:   "case6",
		Code:   http.StatusOK,
		TaskID: 1,
		UpdateRequest: structure.UpdateTaskRequest{
			Task: makePointerStr("Reupdated Task 1"),
		}, ReturnValue: structure.Todo{
			ID:       1,
			Task:     "Reupdated Task 1",
			Priority: "high",
			Status:   "done",
			UserName: "rainbow777",
		},
	},
	{
		Name:   "case7",
		Code:   http.StatusOK,
		TaskID: 2,
		UpdateRequest: structure.UpdateTaskRequest{
			Task: makePointerStr("Updated Task 2"),
		}, ReturnValue: structure.Todo{
			ID:       2,
			Task:     "Updated Task 2",
			Priority: "high",
			Status:   "not_done",
			UserName: "rainbow777",
		},
	},
	{
		Name:   "case8",
		Code:   http.StatusBadRequest,
		TaskID: 2,
		UpdateRequest: structure.UpdateTaskRequest{
			Priority: makePointerStr("midium"),
		},
		ExpectedError: myerrors.MyAppError{
			ErrCode: myerrors.BadRequest,
			Message: "Please enter a valid Column name",
		},
	},
	{
		Name: "case9",
		Code: http.StatusBadRequest,
		UpdateRequest: structure.UpdateTaskRequest{
			Priority: makePointerStr("medium"),
		},
		ExpectedError: myerrors.MyAppError{
			ErrCode: myerrors.BadPath,
			Message: "TaskID is requred in path parameter",
		},
	},

	{
		Name:   "case10",
		Code:   http.StatusOK,
		TaskID: 2,
		UpdateRequest: structure.UpdateTaskRequest{
			Priority: makePointerStr("medium"),
			Status:   makePointerStr("done"),
		}, ReturnValue: structure.Todo{
			ID:       2,
			Task:     "Updated Task 2",
			Priority: "medium",
			Status:   "done",
			UserName: "rainbow777",
		},
	},
}
