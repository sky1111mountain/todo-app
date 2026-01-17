package myerrors

import "fmt"

type MyAppError struct {
	ErrCode
	Message string
	Err     error `json:"-"`
}

func (myErr *MyAppError) Error() string {
	if myErr == nil {
		return "nil MyAppError"
	}

	output := fmt.Sprintf("[%s] %s", myErr.ErrCode, myErr.Message)

	if myErr.Err != nil {
		output = fmt.Sprintf("%s: %s", output, myErr.Err.Error())
	}

	return output
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}

func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}
