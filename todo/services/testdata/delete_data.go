package testdata

import (
	"errors"

	"github.com/rainbow777/todolist/myerrors"
)

type DeleteTaskTestData struct {
	TestName      string
	TaskID        int
	AuthUserName  string
	RepoReturnErr error
	ExpectedErr   myerrors.MyAppError
}

var DleteTaskTestCases = []DeleteTaskTestData{
	{
		TestName: "case1_Succusses",

		TaskID:       1,
		AuthUserName: "rainbow777",
	},
	{
		TestName:     "case2_Err_DeleteTask",
		TaskID:       2,
		AuthUserName: "rainbow777",
		RepoReturnErr: &myerrors.MyAppError{
			ErrCode: myerrors.TaskDeleteFailed,
			Message: "Failed to delete task",
			Err:     errors.New("SIMULATED: Here is Orizinal Err from db.Exec"),
		},
		ExpectedErr: myerrors.MyAppError{
			ErrCode: myerrors.DeleteDataFailed,
			Message: "Error occurred in DeleteTaskDB",
			Err: &myerrors.MyAppError{
				ErrCode: myerrors.TaskDeleteFailed,
				Message: "Failed to delete task",
				Err:     errors.New("SIMULATED: Here is Orizinal Err from db.Exec"),
			},
		},
	},
	{
		TestName:     "case3_Err_RowAffected",
		TaskID:       3,
		AuthUserName: "rainbow777",
		RepoReturnErr: &myerrors.MyAppError{
			ErrCode: myerrors.NoRowsAffected,
			Message: "Could not confirm database changes were complete",
			Err:     errors.New("SIMULATED: Here is Orizinal Err from result.RowsAffected"),
		},
		ExpectedErr: myerrors.MyAppError{
			ErrCode: myerrors.DeleteDataFailed,
			Message: "Error occurred in DeleteTaskDB",
			Err: &myerrors.MyAppError{
				ErrCode: myerrors.NoRowsAffected,
				Message: "Could not confirm database changes were complete",
				Err:     errors.New("SIMULATED: Here is Orizinal Err from result.RowsAffected"),
			},
		},
	},
	{
		TestName:     "case4_DBNotChange",
		TaskID:       3,
		AuthUserName: "rainbow777",
		RepoReturnErr: &myerrors.MyAppError{
			ErrCode: myerrors.UnChangeRows,
			Message: "Database has not changed",
			Err:     errors.New("No affectedRows"),
		},
		ExpectedErr: myerrors.MyAppError{
			ErrCode: myerrors.DeleteDataFailed,
			Message: "Error occurred in DeleteTaskDB",
			Err: &myerrors.MyAppError{
				ErrCode: myerrors.UnChangeRows,
				Message: "Database has not changed",
				Err:     errors.New("No affectedRows"),
			},
		},
	},
}
