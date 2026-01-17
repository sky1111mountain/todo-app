package testdata

import (
	"database/sql"
	"errors"

	"github.com/rainbow777/todolist/myerrors"
	"github.com/rainbow777/todolist/structure"
)

type GetTaskTestData struct {
	TestName      string
	RequestID     int
	UserName      string
	ExpectedTask  structure.Todo
	RepoReturnErr error
	ExpectedErr   myerrors.MyAppError
}

var GetTaskTestCases = []GetTaskTestData{
	{
		TestName:  "case1",
		RequestID: 1,
		UserName:  "rainbow777",
		ExpectedTask: structure.Todo{
			ID:       1,
			Task:     "食器洗い",
			Priority: "medium",
			Status:   "not_done",
			UserName: "rainbow777",
		},
	},
	{
		TestName:  "case2",
		RequestID: 2,
		UserName:  "rainbow777",
		ExpectedTask: structure.Todo{
			ID:       2,
			Task:     "タイヤ交換",
			Priority: "high",
			Status:   "not_done",
			UserName: "rainbow777",
		},
	},
	{
		TestName:     "case3",
		RequestID:    3,
		UserName:     "rainbow777",
		ExpectedTask: structure.Todo{},
		RepoReturnErr: &myerrors.MyAppError{
			ErrCode: myerrors.ScanFailed,
			Message: "Error occurred in row.Scan in GetTaskDB",
			Err:     sql.ErrNoRows,
		},
		ExpectedErr: myerrors.MyAppError{
			ErrCode: myerrors.NAData,
			Message: "Service Error:  TaskID 3 is no data",
			Err: &myerrors.MyAppError{
				ErrCode: myerrors.ScanFailed,
				Message: "Error occurred in row.Scan in GetTaskDB",
				Err:     sql.ErrNoRows,
			},
		},
	},

	{
		TestName:     "case4",
		RequestID:    2,
		UserName:     "rainbow777",
		ExpectedTask: structure.Todo{},
		RepoReturnErr: &myerrors.MyAppError{
			ErrCode: myerrors.ScanFailed,
			Message: "Error occurred in row.Scan in GetTaskDB",
			Err:     errors.New("SIMULATED: Orizinal Error from row.Scan"),
		},
		ExpectedErr: myerrors.MyAppError{
			ErrCode: myerrors.GetDataFailed,
			Message: "Failed to get task",
			Err: &myerrors.MyAppError{
				ErrCode: myerrors.ScanFailed,
				Message: "Error occurred in row.Scan in GetTaskDB",
				Err:     errors.New("SIMULATED: Orizinal Error from row.Scan"),
			},
		},
	},
}
