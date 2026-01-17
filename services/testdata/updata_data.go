package testdata

import (
	"errors"

	"github.com/rainbow777/todolist/myerrors"
	"github.com/rainbow777/todolist/structure"
)

type UpdateTaskTestData struct {
	TestName       string
	UpdateData     structure.UpdateData
	RepoReturnTask structure.Todo // 呼び出し先であるレポジトリ層から返ってくるTask
	ExpectedTask   structure.Todo
	RepoReturnErr  error // 呼び出し先であるレポジトリ層から返ってくるErr
	ExpectedErr    myerrors.MyAppError
}

func makePointer(s string) *string { return &s }

var UpdateTaskTestCases = []UpdateTaskTestData{
	{
		TestName: "case1_Status_Update",
		UpdateData: structure.UpdateData{
			TaskID:       1,
			AuthUserName: "rainbow777",
			UpdateRequest: structure.UpdateTaskRequest{
				Status: makePointer("done"),
			},
		},
		RepoReturnTask: structure.Todo{
			ID:       1,
			Task:     "タイヤ交換",
			Priority: "high",
			Status:   "done",
			UserName: "rainbow777",
		},
		ExpectedTask: structure.Todo{
			ID:       1,
			Task:     "タイヤ交換",
			Priority: "high",
			Status:   "done",
			UserName: "rainbow777",
		},
	},
	{
		TestName: "case2_Task_Update",
		UpdateData: structure.UpdateData{
			TaskID:       1,
			AuthUserName: "rainbow777",
			UpdateRequest: structure.UpdateTaskRequest{
				Task: makePointer("タイヤ交換(ステップワゴン)"),
			},
		},
		RepoReturnTask: structure.Todo{
			ID:       1,
			Task:     "タイヤ交換(ステップワゴン)",
			Priority: "high",
			Status:   "done",
			UserName: "rainbow777",
		},
		ExpectedTask: structure.Todo{
			ID:       1,
			Task:     "タイヤ交換(ステップワゴン)",
			Priority: "high",
			Status:   "done",
			UserName: "rainbow777",
		},
	},
	{
		TestName: "case3_Priority_Update",
		UpdateData: structure.UpdateData{
			TaskID:       1,
			AuthUserName: "rainbow777",
			UpdateRequest: structure.UpdateTaskRequest{
				Priority: makePointer("medium"),
			},
		},
		RepoReturnTask: structure.Todo{
			ID:       1,
			Task:     "タイヤ交換(ステップワゴン)",
			Priority: "medium",
			Status:   "done",
			UserName: "rainbow777",
		},
		ExpectedTask: structure.Todo{
			ID:       1,
			Task:     "タイヤ交換(ステップワゴン)",
			Priority: "medium",
			Status:   "done",
			UserName: "rainbow777",
		},
	},

	{
		TestName: "case5_Err_Unupdate",

		UpdateData: structure.UpdateData{
			TaskID:       1,
			AuthUserName: "rainbow777",
			UpdateRequest: structure.UpdateTaskRequest{
				Task: makePointer("タイヤ交換(ノート)"),
			},
		},
		RepoReturnTask: structure.Todo{
			ID:       1,
			Task:     "タイヤ交換(ステップワゴン)",
			Priority: "medium",
			Status:   "done",
			UserName: "UserA",
		},
		ExpectedTask: structure.Todo{
			ID:       1,
			Task:     "タイヤ交換(ステップワゴン)",
			Priority: "medium",
			Status:   "done",
			UserName: "UserA",
		},
		ExpectedErr: myerrors.MyAppError{
			ErrCode: myerrors.UnUpdatedTask,
			Message: "Could not check for data updates.",
			Err:     myerrors.ErrUnUpdate,
		},
	},

	{
		TestName: "case6_Err_ColumnLess",
		UpdateData: structure.UpdateData{
			TaskID:       1,
			AuthUserName: "rainbow777",
		},
		ExpectedErr: myerrors.MyAppError{
			ErrCode: myerrors.NoUpdateColumn,
			Message: "At least one colomn required",
			Err:     errors.New("No update colomn"),
		},
	},

	{
		TestName: "case7_Err_RepoExec",
		UpdateData: structure.UpdateData{
			TaskID:       1,
			AuthUserName: "rainbow777",
			UpdateRequest: structure.UpdateTaskRequest{
				Task: makePointer("タイヤ交換(ノート)"),
			},
		},
		RepoReturnErr: &myerrors.MyAppError{
			ErrCode: myerrors.TaskUpdateFailed,
			Message: "Error occurred in tx.Exec",
			Err:     errors.New("SIMULATED: Failed to execute Exec "),
		},
		ExpectedErr: myerrors.MyAppError{
			ErrCode: myerrors.UpdateDataFailed,
			Message: "Error occurred in UpdateTaskService",
			Err: &myerrors.MyAppError{
				ErrCode: myerrors.TaskUpdateFailed,
				Message: "Error occurred in tx.Exec",
				Err:     errors.New("SIMULATED: Failed to execute Exec "),
			},
		},
	},

	{
		TestName: "case8_Err_RepoGet",
		UpdateData: structure.UpdateData{
			TaskID:       1,
			AuthUserName: "rainbow777",
			UpdateRequest: structure.UpdateTaskRequest{
				Task: makePointer("タイヤ交換(ノート)"),
			},
		},
		RepoReturnErr: &myerrors.MyAppError{
			ErrCode: myerrors.ScanFailed,
			Message: "Error occurred in row.Scan",
			Err:     errors.New("SIMULATED:  This is original err from row.Scan"),
		},
		ExpectedErr: myerrors.MyAppError{
			ErrCode: myerrors.UpdateDataFailed,
			Message: "Error occurred in UpdateTaskService",
			Err: &myerrors.MyAppError{
				ErrCode: myerrors.ScanFailed,
				Message: "Error occurred in row.Scan",
				Err:     errors.New("SIMULATED:  This is original err from row.Scan"),
			},
		},
	},

	{
		TestName: "case9_ALL_Update",
		UpdateData: structure.UpdateData{
			TaskID:       1,
			AuthUserName: "rainbow777",
			UpdateRequest: structure.UpdateTaskRequest{
				Task:     makePointer("タイヤ空気圧点検"),
				Priority: makePointer("low"),
				Status:   makePointer("not_done"),
			},
		},
		RepoReturnTask: structure.Todo{
			ID:       1,
			Task:     "タイヤ空気圧点検",
			Priority: "low",
			Status:   "not_done",
			UserName: "rainbow777",
		},
		ExpectedTask: structure.Todo{
			ID:       1,
			Task:     "タイヤ空気圧点検",
			Priority: "low",
			Status:   "not_done",
			UserName: "rainbow777",
		},
	},

	{
		TestName: "case10_Error_InvalidPriority",
		UpdateData: structure.UpdateData{
			TaskID:       1,
			AuthUserName: "rainbow777",
			UpdateRequest: structure.UpdateTaskRequest{
				Priority: makePointer("midium"),
			},
		},
		ExpectedErr: myerrors.MyAppError{
			ErrCode: myerrors.BadRequest,
			Message: "Priority require high or medium or low not midium",
			Err:     errors.New("invalid priority"),
		},
	},
}
