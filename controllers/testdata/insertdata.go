package testdata

import (
	"net/http"

	"github.com/rainbow777/todolist/myerrors"
	"github.com/rainbow777/todolist/structure"
)

type InsertedTaskTestData struct {
	Name          string
	Code          int
	InsertTask    structure.InsertData
	ExpectedTask  structure.Todo
	ExpectedError myerrors.MyAppError
}

var InsertTaskTestCases = []InsertedTaskTestData{
	{
		Name: "case1",
		Code: http.StatusOK,

		InsertTask: structure.InsertData{
			Task: "洗濯", Priority: "medium", Status: "not_done", AuthUserName: "rainbow777",
		},

		ExpectedTask: structure.Todo{
			ID: 1, Task: "洗濯", Priority: "medium", Status: "not_done", UserName: "rainbow777",
		},
	},

	{
		Name: "case2",
		Code: http.StatusForbidden,

		InsertTask: structure.InsertData{
			Task: "面接", Priority: "high", Status: "not_done", AuthUserName: "rainbow77",
		},

		ExpectedError: myerrors.MyAppError{
			ErrCode: myerrors.NotMatchUserName, Message: "You don't have the authority to do that",
		},
	},

	{
		Name: "case3",
		Code: http.StatusOK,

		InsertTask: structure.InsertData{
			Task: "ミーティング", Priority: "high", Status: "not_done", AuthUserName: "rainbow777",
		},

		ExpectedTask: structure.Todo{
			ID: 2, Task: "ミーティング", Priority: "high", Status: "not_done", UserName: "rainbow777",
		},
	},
}
