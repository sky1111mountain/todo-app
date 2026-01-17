package testdata

import (
	"net/http"

	"github.com/rainbow777/todolist/myerrors"
	"github.com/rainbow777/todolist/structure"
)

// GetTaskHandler用テストデータ
type GetTaskTestData struct {
	Name          string
	Code          int
	ID            string
	ExpectedTask  structure.Todo      // 成功時のみ使用
	ExpectedError myerrors.MyAppError //失敗時のみ使用
}

var GetTaskTestCases = []GetTaskTestData{

	{
		Name: "case1_Success",
		Code: http.StatusOK,
		ID:   "1",
		ExpectedTask: structure.Todo{
			ID:       1,
			Task:     "Test Task 1",
			Priority: "high",
			Status:   "not_done",
			UserName: "rainbow777",
		},
	},
	{
		Name: "case2_Success",
		Code: http.StatusOK,
		ID:   "2",
		ExpectedTask: structure.Todo{
			ID:       2,
			Task:     "Test Task 2",
			Priority: "medium",
			Status:   "not_done",
			UserName: "rainbow777",
		},
	},
	{
		Name: "case3_NaData",
		Code: http.StatusNotFound,
		ID:   "3",
		ExpectedError: myerrors.MyAppError{
			ErrCode: myerrors.NAData,
			Message: "no data",
		},
	},
	{
		Name: "case4_Invalid",
		Code: http.StatusBadRequest,
		ID:   "aaa",
		ExpectedError: myerrors.MyAppError{
			ErrCode: myerrors.BadPath,
			Message: "invalid path parameter",
		},
	},
}
