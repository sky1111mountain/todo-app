package testdata

import (
	"errors"

	"github.com/rainbow777/todolist/myerrors"
	"github.com/rainbow777/todolist/structure"
)

type GetListTestData struct {
	TestName      string
	Request       structure.GetListRequest
	ExpectedTask  []structure.Todo
	RepoReturnErr error // レポジトリ層から返されるエラーケースを格納
	ExpectedErr   myerrors.MyAppError
}

var GetListTestCases = []GetListTestData{
	{
		TestName: "case1_high_SUCCESS",
		Request: structure.GetListRequest{
			AuthUserName: "rainbow777",
			Priorities:   []string{"high"},
		},
		ExpectedTask: []structure.Todo{
			{
				ID:       1,
				Task:     "タイヤ交換",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			},
			{
				ID:       3,
				Task:     "年金納付",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			},
		},
	},
	{
		TestName: "case2_medium_SUCCESS",
		Request: structure.GetListRequest{
			AuthUserName: "rainbow777",
			Priorities:   []string{"medium"},
			Offset:       0,
			Limit:        20,
		},
		ExpectedTask: []structure.Todo{
			{
				ID:       2,
				Task:     "食器洗い",
				Priority: "medium",
				Status:   "not_done",
				UserName: "rainbow777",
			},
		},
	},
	{
		TestName: "case3_low_NADATA",
		Request: structure.GetListRequest{
			AuthUserName: "rainbow777",
			Priorities:   []string{"low"},
			Offset:       0,
			Limit:        20,
		},
		ExpectedTask: []structure.Todo{},
		ExpectedErr: myerrors.MyAppError{
			ErrCode: myerrors.NAData,
			Message: "no data",
			Err:     myerrors.ErrNoData,
		},
	},

	{
		TestName: "case4_FULL_SUCCESS",
		Request: structure.GetListRequest{
			AuthUserName: "rainbow777",
			Status:       "all",
			Offset:       0,
			Limit:        20,
		},
		ExpectedTask: []structure.Todo{
			{
				ID:       1,
				Task:     "タイヤ交換",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			},
			{
				ID:       2,
				Task:     "食器洗い",
				Priority: "medium",
				Status:   "not_done",
				UserName: "rainbow777",
			},
			{
				ID:       3,
				Task:     "年金納付",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			},
		},
	},

	{
		TestName: "case5_ERROR_GETROWS",
		Request: structure.GetListRequest{
			AuthUserName: "rainbow777",
			Priorities:   []string{"low"},
		},
		RepoReturnErr: &myerrors.MyAppError{
			ErrCode: myerrors.GetListfailed,
			Message: "Error occured in db.Query in GetRows func",
			Err:     errors.New("SIMULATED: Database Query failed due to closed connection"),
		},
		ExpectedErr: myerrors.MyAppError{
			ErrCode: myerrors.GetDataFailed,
			Message: "Error occured in GetTodoListDB",
			Err: &myerrors.MyAppError{
				ErrCode: myerrors.GetListfailed,
				Message: "Error occured in db.Query in GetRows func",
				Err:     errors.New("SIMULATED: Database Query failed due to closed connection"),
			},
		},
	},

	{
		TestName: "case6_ERROR_SCANROWS",
		Request: structure.GetListRequest{
			AuthUserName: "rainbow777",
			Priorities:   []string{"low"},
		},
		RepoReturnErr: &myerrors.MyAppError{
			ErrCode: myerrors.GetListfailed,
			Message: "Error occuerred in ScanRows in helper.go",
			Err:     errors.New("SIMULATED: Failed to scan rows"),
		},
		ExpectedErr: myerrors.MyAppError{
			ErrCode: myerrors.GetDataFailed,
			Message: "Error occured in GetTodoListDB",
			Err: &myerrors.MyAppError{
				ErrCode: myerrors.GetListfailed,
				Message: "Error occuerred in ScanRows in helper.go",
				Err:     errors.New("SIMULATED: Failed to scan rows"),
			},
		},
	},
}
