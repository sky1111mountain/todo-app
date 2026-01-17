package controllers

import (
	"net/http"
	"strconv"

	"github.com/rainbow777/todolist/myerrors"
	"github.com/rainbow777/todolist/structure"
)

// クエリパラメータの内容を確認する関数
func CheckQueryParam(req *http.Request) (*structure.GetListRequest, error) {
	getListRequest := structure.GetListRequest{
		Status: "not_done",
		Offset: 0,
		Limit:  20,
	}

	reqQuery := req.URL.Query()

	if len(reqQuery) == 0 {
		return &getListRequest, nil
	}

	// クエリ名の適正値リストの定義
	allowedQueries := map[string]bool{"priority": true, "status": true, "page": true}

	// クエリ名の検証
	for query := range reqQuery {
		if !allowedQueries[query] {
			return nil, myerrors.BadQuery.Wrap(myerrors.ErrQuery, "Query name must be priority or status or page")
		}
	}

	if reqQuery["priority"] != nil {
		getListRequest.Priorities = reqQuery["priority"]
		for _, priority := range getListRequest.Priorities {
			if priority != "high" && priority != "medium" && priority != "low" {
				return nil, myerrors.BadQuery.Wrap(myerrors.ErrQuery,
					"priority requier high or medium or low")
			}
		}
	}

	if statusSlice := reqQuery["status"]; len(statusSlice) > 1 {
		return nil, myerrors.BadQuery.Wrap(myerrors.ErrQuery, "Status must be only one parameter")
	}

	allowedStatus := map[string]bool{"done": true, "all": true, "All": true, "not_done": true}

	for _, status := range reqQuery["status"] {
		if !allowedStatus[status] {
			return nil, myerrors.BadQuery.Wrap(myerrors.ErrQuery, "Status must be all, done, not_done")
		}
		getListRequest.Status = status
	}

	page := 1
	if reqQuery["page"] != nil {
		var err error
		pageSLise := reqQuery["page"]
		if len(pageSLise) == 1 {
			page, err = strconv.Atoi(pageSLise[0])
			if err != nil {
				return nil, myerrors.BadQuery.Wrap(myerrors.ErrQuery,
					"page must be a valid number")
			}
		} else if len(pageSLise) > 1 {
			return nil, myerrors.BadQuery.Wrap(myerrors.ErrQuery,
				"Only one page can be specified at a time")
		}

		if page < 1 {
			return nil, myerrors.BadQuery.Wrap(myerrors.ErrQuery,
				"Page number must be 1 or more")
		}

		if page > 1000 {
			return nil, myerrors.BadQuery.Wrap(myerrors.ErrQuery,
				"You can specify a maximum of 1000 pages")
		}

	}

	getListRequest.Offset = (page - 1) * getListRequest.Limit

	return &getListRequest, nil
}
