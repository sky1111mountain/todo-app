package testdata

import (
	"github.com/rainbow777/todolist/myerrors"
	"github.com/rainbow777/todolist/structure"
)

// InsertTaskDB用
var InsertTaskTestCases = []struct {
	TestName     string
	InsertTask   structure.InsertData
	ExpectedTask structure.Todo
}{
	{
		TestName: "case1",
		InsertTask: structure.InsertData{
			Task:         "Insert Test Data 1",
			Priority:     "high",
			Status:       "not_done",
			AuthUserName: "rainbow777",
		},
		ExpectedTask: structure.Todo{
			ID:       1,
			Task:     "Insert Test Data 1",
			Priority: "high",
			Status:   "not_done",
			UserName: "rainbow777",
		},
	},

	{
		TestName: "case2",
		InsertTask: structure.InsertData{
			Task:         "Insert Test Data 2",
			Priority:     "medium",
			Status:       "done",
			AuthUserName: "rainbow777",
		},
		ExpectedTask: structure.Todo{
			ID:       2,
			Task:     "Insert Test Data 2",
			Priority: "medium",
			Status:   "done",
			UserName: "rainbow777",
		},
	},
}

var TestTaskForInsertErr = structure.InsertData{
	Task:         "test",
	Priority:     "high",
	Status:       "done",
	AuthUserName: "rainbow777",
}

var GetTaskTestCases = []struct {
	TestName     string
	TaskID       int
	AuthUserName string
	ExpectedTask structure.Todo
}{
	{
		TestName:     "Case1",
		TaskID:       1,
		AuthUserName: "rainbow777",
		ExpectedTask: structure.Todo{
			ID:       1,
			Task:     "Test Data 1",
			Priority: "high",
			Status:   "not_done",
			UserName: "rainbow777",
		},
	},
	{
		TestName:     "Case2",
		TaskID:       2,
		AuthUserName: "rainbow777",
		ExpectedTask: structure.Todo{
			ID:       2,
			Task:     "Test Data 2",
			Priority: "medium",
			Status:   "done",
			UserName: "rainbow777",
		},
	},
	{
		TestName:     "Case3",
		TaskID:       3,
		AuthUserName: "rainbow777",
		ExpectedTask: structure.Todo{
			ID:       3,
			Task:     "Test Data 3",
			Priority: "low",
			Status:   "not_done",
			UserName: "rainbow777",
		},
	},
}

var GetListTestCases = []struct {
	TestName     string
	Request      structure.GetListRequest
	ExpectedList []structure.Todo
}{
	{
		TestName: "Case1_high",
		Request: structure.GetListRequest{
			AuthUserName: "rainbow777",
			Priorities:   []string{"high"},
			Offset:       1,
			Limit:        20,
		},
		ExpectedList: []structure.Todo{
			{
				ID:       1,
				Task:     "Test Data 1",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			},
			{
				ID:       2,
				Task:     "Test Data 2",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			},
			{
				ID:       3,
				Task:     "Test Data 3",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			},
		},
	},
	{
		TestName: "Case2_medium",
		Request: structure.GetListRequest{
			AuthUserName: "rainbow777",
			Priorities:   []string{"medium"},
			Offset:       1,
			Limit:        20,
		},
		ExpectedList: []structure.Todo{
			{
				ID:       4,
				Task:     "Test Data 4",
				Priority: "medium",
				Status:   "done",
				UserName: "rainbow777"},
			{
				ID:       5,
				Task:     "Test Data 5",
				Priority: "medium",
				Status:   "done",
				UserName: "rainbow777",
			},
		},
	},
	{
		TestName: "Case3_low",
		Request: structure.GetListRequest{
			AuthUserName: "rainbow777",
			Priorities:   []string{"low"},
			Offset:       1,
			Limit:        20,
		},
	},
	{
		TestName: "Case4_NotQuery",
		Request: structure.GetListRequest{
			AuthUserName: "rainbow777",
			Offset:       1,
			Limit:        20,
		},
		ExpectedList: []structure.Todo{
			{
				ID:       1,
				Task:     "Test Data 1",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			},
			{
				ID:       2,
				Task:     "Test Data 2",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			},
			{
				ID:       3,
				Task:     "Test Data 3",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			},
		},
	},
}

func makePointer(s string) *string { return &s }

// update用
var UpdateTaskTestCases = []struct {
	TestName     string
	UpdateData   structure.UpdateData
	ExpectedTask structure.Todo
	ExpectedErr  myerrors.MyAppError
}{
	{
		TestName: "Case1_Success",
		UpdateData: structure.UpdateData{
			TaskID:       1,
			AuthUserName: "rainbow777",
			UpdateRequest: structure.UpdateTaskRequest{
				Priority: makePointer("medium"),
			},
		},
		ExpectedTask: structure.Todo{
			ID:       1,
			Task:     "Test Data 1",
			Priority: "medium",
			Status:   "not_done",
			UserName: "rainbow777",
		},
	},

	{
		TestName: "Case2_Success",
		UpdateData: structure.UpdateData{
			TaskID:       1,
			AuthUserName: "rainbow777",
			UpdateRequest: structure.UpdateTaskRequest{
				Status: makePointer("done"),
			},
		},
		ExpectedTask: structure.Todo{
			ID:       1,
			Task:     "Test Data 1",
			Priority: "medium",
			Status:   "done",
			UserName: "rainbow777",
		},
	},

	{
		TestName: "Case3_Success",
		UpdateData: structure.UpdateData{
			TaskID:       1,
			AuthUserName: "rainbow777",
			UpdateRequest: structure.UpdateTaskRequest{
				Task: makePointer("Updated Data 1"),
			},
		},
		ExpectedTask: structure.Todo{
			ID:       1,
			Task:     "Updated Data 1",
			Priority: "medium",
			Status:   "done",
			UserName: "rainbow777",
		},
	},
	{
		TestName: "Case4_Success",
		UpdateData: structure.UpdateData{
			TaskID:       2,
			AuthUserName: "rainbow777",
			UpdateRequest: structure.UpdateTaskRequest{
				Priority: makePointer("low"),
			},
		},
		ExpectedTask: structure.Todo{
			ID:       2,
			Task:     "Test Data 2",
			Priority: "low",
			Status:   "not_done",
			UserName: "rainbow777",
		},
	},
	{
		TestName: "Case5_Success",
		UpdateData: structure.UpdateData{
			TaskID:       2,
			AuthUserName: "rainbow777",
			UpdateRequest: structure.UpdateTaskRequest{
				Status: makePointer("done"),
			},
		},
		ExpectedTask: structure.Todo{
			ID:       2,
			Task:     "Test Data 2",
			Priority: "low",
			Status:   "done",
			UserName: "rainbow777",
		},
	},
	{
		TestName: "Case6_Success",
		UpdateData: structure.UpdateData{
			TaskID:       2,
			AuthUserName: "rainbow777",
			UpdateRequest: structure.UpdateTaskRequest{
				Task: makePointer("Updated Data 2"),
			},
		},
		ExpectedTask: structure.Todo{
			ID:       2,
			Task:     "Updated Data 2",
			Priority: "low",
			Status:   "done",
			UserName: "rainbow777",
		},
	},
}

var UpdateData = structure.UpdateData{
	TaskID:       1,
	AuthUserName: "rainbow777",
	UpdateRequest: structure.UpdateTaskRequest{
		Task: makePointer("newdata"),
	},
}

// delete用
var DeleteTaskTestData = []struct {
	TestName string
	DataInDB structure.Todo
}{
	{
		TestName: "case1",
		DataInDB: structure.Todo{
			ID:       1,
			Task:     "Test Data 1",
			Priority: "high",
			Status:   "not_done",
			UserName: "rainbow777",
		},
	},
	{
		TestName: "case2",
		DataInDB: structure.Todo{
			ID:       2,
			Task:     "Test Data 2",
			Priority: "medium",
			Status:   "done",
			UserName: "rainbow777",
		},
	},
	{
		TestName: "case3",
		DataInDB: structure.Todo{
			ID:       3,
			Task:     "Test Data 3",
			Priority: "low",
			Status:   "not_done",
			UserName: "rainbow777",
		},
	},
}
