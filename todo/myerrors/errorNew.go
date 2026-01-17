package myerrors

import "errors"

var ErrNoData = errors.New("get 0 record from db.Query")
var ErrRequest = errors.New("invalid request")
var ErrQuery = errors.New("invalid query parameter")
var AffectedNoRows = errors.New("no rows affected")
var ErrColumn = errors.New("invalid column")
var ErrAuthUser = errors.New("Usernames do not match")
var ErrUnUpdate = errors.New("Task is unupdated")
