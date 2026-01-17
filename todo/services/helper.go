package services

import (
	"errors"

	"github.com/rainbow777/todolist/myerrors"
	"github.com/rainbow777/todolist/structure"
)

// バリデーション　リクエストの中身を確認
func ValidationRequest(insertData structure.InsertData) error {
	if insertData.Task == "" {
		return myerrors.BadRequest.Wrap(myerrors.ErrRequest, "Please enter the task content")
	}

	if insertData.Priority == "" {
		return myerrors.BadRequest.Wrap(myerrors.ErrRequest, "Please enter the priority")
	}

	if insertData.Priority != "high" && insertData.Priority != "medium" && insertData.Priority != "low" {
		return myerrors.BadRequest.Wrap(errors.New("invalid priority"),
			"Please specify a valid priority")
	}

	if insertData.Status == "" {
		return myerrors.BadRequest.Wrap(myerrors.ErrRequest, "Please enter the status")
	}

	if insertData.Status != "done" && insertData.Status != "not_done" {
		return myerrors.BadRequest.Wrap(errors.New("invalid priority"),
			"Status must be all, done, not_done")
	}

	return nil
}
