package testdata

import (
	"net/http"

	"github.com/rainbow777/todolist/myerrors"
	"github.com/rainbow777/todolist/structure"
)

// 以下　GetListHandler用テストデータ
type GetListTestData struct {
	Name          string
	QueryParam    string
	Code          int
	ExpectedList  []structure.Todo    // 成功時のみ使用
	ExpectedError myerrors.MyAppError //失敗時のみ使用
}

var GetListTestCases = []GetListTestData{
	{
		Name:       "case1_High",
		QueryParam: "?priority=high",
		Code:       http.StatusOK,
		ExpectedList: []structure.Todo{
			{
				ID:       1,
				Task:     "Test Task 1",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			}, {
				ID:       2,
				Task:     "Test Task 2",
				Priority: "high",
				Status:   "done",
				UserName: "rainbow777",
			}, {
				ID:       3,
				Task:     "Test Task 3",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			},
		},
	},

	{
		Name:       "case2_Medium",
		QueryParam: "?priority=medium",
		Code:       http.StatusOK,
		ExpectedList: []structure.Todo{
			{
				ID:       4,
				Task:     "Test Task 4",
				Priority: "medium",
				Status:   "not_done",
				UserName: "rainbow777",
			}, {
				ID:       5,
				Task:     "Test Task 5",
				Priority: "medium",
				Status:   "done",
				UserName: "rainbow777",
			}, {
				ID:       6,
				Task:     "Test Task 6",
				Priority: "medium",
				Status:   "not_done",
				UserName: "rainbow777",
			},
		},
	},

	{
		Name:         "case3_low_NoData",
		QueryParam:   "?priority=low",
		Code:         http.StatusOK,
		ExpectedList: []structure.Todo{},
	},

	{
		Name:       "case4_2Priority",
		QueryParam: "?priority=high&priority=medium",
		Code:       http.StatusOK,
		ExpectedList: []structure.Todo{
			{
				ID:       1,
				Task:     "Test Task 1",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			}, {
				ID:       2,
				Task:     "Test Task 2",
				Priority: "high",
				Status:   "done",
				UserName: "rainbow777",
			}, {
				ID:       3,
				Task:     "Test Task 3",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			}, {
				ID:       4,
				Task:     "Test Task 4",
				Priority: "medium",
				Status:   "not_done",
				UserName: "rainbow777",
			}, {
				ID:       5,
				Task:     "Test Task 5",
				Priority: "medium",
				Status:   "done",
				UserName: "rainbow777",
			}, {
				ID:       6,
				Task:     "Test Task 6",
				Priority: "medium",
				Status:   "not_done",
				UserName: "rainbow777",
			},
		},
	},

	{
		Name:       "case5_Status_done",
		QueryParam: "?status=done",
		Code:       http.StatusOK,
		ExpectedList: []structure.Todo{
			{
				ID:       2,
				Task:     "Test Task 2",
				Priority: "high",
				Status:   "done",
				UserName: "rainbow777",
			}, {
				ID:       5,
				Task:     "Test Task 5",
				Priority: "medium",
				Status:   "done",
				UserName: "rainbow777",
			},
		},
	},

	{
		Name:       "case6_Status_not_done",
		QueryParam: "?status=not_done",
		Code:       http.StatusOK,
		ExpectedList: []structure.Todo{
			{
				ID:       1,
				Task:     "Test Task 1",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			}, {
				ID:       3,
				Task:     "Test Task 3",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			}, {
				ID:       4,
				Task:     "Test Task 4",
				Priority: "medium",
				Status:   "not_done",
				UserName: "rainbow777",
			}, {
				ID:       6,
				Task:     "Test Task 6",
				Priority: "medium",
				Status:   "not_done",
				UserName: "rainbow777",
			},
		},
	},

	{
		Name:       "case7_Status_all",
		QueryParam: "?status=all",
		Code:       http.StatusOK,
		ExpectedList: []structure.Todo{
			{
				ID:       1,
				Task:     "Test Task 1",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			}, {
				ID:       2,
				Task:     "Test Task 2",
				Priority: "high",
				Status:   "done",
				UserName: "rainbow777",
			}, {
				ID:       3,
				Task:     "Test Task 3",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			}, {
				ID:       4,
				Task:     "Test Task 4",
				Priority: "medium",
				Status:   "not_done",
				UserName: "rainbow777",
			}, {
				ID:       5,
				Task:     "Test Task 5",
				Priority: "medium",
				Status:   "done",
				UserName: "rainbow777",
			}, {
				ID:       6,
				Task:     "Test Task 6",
				Priority: "medium",
				Status:   "not_done",
				UserName: "rainbow777",
			},
		},
	},

	{
		Name:       "case8_NoQuery",
		QueryParam: "",
		Code:       http.StatusOK,
		ExpectedList: []structure.Todo{
			{
				ID:       1,
				Task:     "Test Task 1",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			}, {
				ID:       3,
				Task:     "Test Task 3",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			}, {
				ID:       4,
				Task:     "Test Task 4",
				Priority: "medium",
				Status:   "not_done",
				UserName: "rainbow777",
			}, {
				ID:       6,
				Task:     "Test Task 6",
				Priority: "medium",
				Status:   "not_done",
				UserName: "rainbow777",
			},
		},
	},

	{
		Name:       "case9_BadParam_hig",
		QueryParam: "?priority=hig",
		Code:       http.StatusBadRequest,
		ExpectedError: myerrors.MyAppError{
			ErrCode: myerrors.BadQuery,
			Message: "priority requier high or medium or low",
		},
	},

	{
		Name:       "case10_BadParam_notdone",
		QueryParam: "?status=notdone",
		Code:       http.StatusBadRequest,
		ExpectedError: myerrors.MyAppError{
			ErrCode: myerrors.BadQuery,
			Message: "Status must be all, done, not_done",
		},
	},

	{
		Name:       "case11_BadParam_2Status",
		QueryParam: "?status=done&status=not_done",
		Code:       http.StatusBadRequest,
		ExpectedError: myerrors.MyAppError{
			ErrCode: myerrors.BadQuery,
			Message: "Status must be only one parameter",
		},
	},

	{
		Name:       "case12_BadParam_page=a",
		QueryParam: "?page=a",
		Code:       http.StatusBadRequest,
		ExpectedError: myerrors.MyAppError{
			ErrCode: myerrors.BadQuery,
			Message: "page must be a valid number",
		},
	},

	{
		Name:       "case13_BadParam_page=0",
		QueryParam: "?page=0",
		Code:       http.StatusBadRequest,
		ExpectedError: myerrors.MyAppError{
			ErrCode: myerrors.BadQuery,
			Message: "Page number must be 1 or more",
		},
	},

	{
		Name:       "case15_BadQuery_2pages",
		QueryParam: "?page=1&page=3",
		Code:       http.StatusBadRequest,
		ExpectedError: myerrors.MyAppError{
			ErrCode: myerrors.BadQuery,
			Message: "Only one page can be specified at a time",
		},
	},

	{
		Name:       "case15_BadQuery_LimitOver",
		QueryParam: "?page=9999999",
		Code:       http.StatusBadRequest,
		ExpectedError: myerrors.MyAppError{
			ErrCode: myerrors.BadQuery,
			Message: "You can specify a maximum of 1000 pages",
		},
	},

	{
		Name:       "case14_Bad_QueryName",
		QueryParam: "?prioriti=high",
		Code:       http.StatusBadRequest,
		ExpectedError: myerrors.MyAppError{
			ErrCode: myerrors.BadQuery,
			Message: "Query name must be priority or status or page",
		},
	},

	{
		Name:       "case15_MixedBadQueryName",
		QueryParam: "?priority=medium&badquery=high",
		Code:       http.StatusBadRequest,
		ExpectedError: myerrors.MyAppError{
			ErrCode: myerrors.BadQuery,
			Message: "Query name must be priority or status or page",
		},
	},

	{
		Name:       "case7_Status_all",
		QueryParam: "?status=All",
		Code:       http.StatusOK,
		ExpectedList: []structure.Todo{
			{
				ID:       1,
				Task:     "Test Task 1",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			}, {
				ID:       2,
				Task:     "Test Task 2",
				Priority: "high",
				Status:   "done",
				UserName: "rainbow777",
			}, {
				ID:       3,
				Task:     "Test Task 3",
				Priority: "high",
				Status:   "not_done",
				UserName: "rainbow777",
			}, {
				ID:       4,
				Task:     "Test Task 4",
				Priority: "medium",
				Status:   "not_done",
				UserName: "rainbow777",
			}, {
				ID:       5,
				Task:     "Test Task 5",
				Priority: "medium",
				Status:   "done",
				UserName: "rainbow777",
			}, {
				ID:       6,
				Task:     "Test Task 6",
				Priority: "medium",
				Status:   "not_done",
				UserName: "rainbow777",
			},
		},
	},
}
