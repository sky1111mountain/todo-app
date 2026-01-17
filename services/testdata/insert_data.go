package testdata

import (
	"github.com/rainbow777/todolist/myerrors"
	"github.com/rainbow777/todolist/structure"
)

type InsertTaskServiceData struct {
	TestName     string
	InsertTask   structure.InsertData
	ExpectedTask structure.Todo
	ExpectedErr  myerrors.MyAppError
}

var InsertTaskTestCases = []InsertTaskServiceData{

	{
		TestName: "case1",
		InsertTask: structure.InsertData{
			Task:         "test1",
			Priority:     "high",
			Status:       "not_done",
			AuthUserName: "rainbow777",
		},
		ExpectedTask: structure.Todo{
			ID:       1,
			Task:     "test1",
			Priority: "high",
			Status:   "not_done",
			UserName: "rainbow777",
		},
	},
	{
		TestName: "case2",
		InsertTask: structure.InsertData{
			Task:         "test2",
			Priority:     "medium",
			Status:       "not_done",
			AuthUserName: "rainbow777",
		},
		ExpectedTask: structure.Todo{
			ID:       2,
			Task:     "test2",
			Priority: "medium",
			Status:   "not_done",
			UserName: "rainbow777",
		},
	},
	{
		TestName: "case3",
		InsertTask: structure.InsertData{
			Task:         "入金の確認",
			Priority:     "very high",
			Status:       "not_done",
			AuthUserName: "rainbow777",
		},
		ExpectedErr: myerrors.MyAppError{
			ErrCode: myerrors.BadRequest,
			Message: "Please specify a valid priority",
		},
	},

	{
		TestName: "case4",
		InsertTask: structure.InsertData{
			Task:         "",
			Priority:     "high",
			Status:       "not_done",
			AuthUserName: "rainbow777",
		},
		ExpectedErr: myerrors.MyAppError{
			ErrCode: myerrors.BadRequest,
			Message: "Please enter the task content",
		},
	},
}
