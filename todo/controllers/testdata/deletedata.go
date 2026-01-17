package testdata

import (
	"net/http"
)

type DeleteTestData struct {
	Name        string
	Code        int
	URL         string
	ExpectedRes string
}

var DeleteTestCases = []DeleteTestData{

	{
		Name:        "case1",
		Code:        http.StatusOK,
		URL:         "/todo/delete/1",
		ExpectedRes: "The specified task has been deleted.",
	},
	{
		Name:        "case2",
		Code:        http.StatusOK,
		URL:         "/todo/delete/2",
		ExpectedRes: "The specified task has been deleted.",
	},
	{
		Name:        "case3",
		Code:        http.StatusBadRequest,
		URL:         "/todo/delete/あ",
		ExpectedRes: "TaskID is requred in path parameter",
	},
}
