package myerrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, err error) {
	var myErr *MyAppError
	// 第三引数で受け取ったerror型のerrを自作エラー型に変換する
	if !errors.As(err, &myErr) {
		myErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed　事前に想定していないエラーが発生しました",
			Err:     err,
		}
	}

	// 開発者向けに詳細なエラーログを出力
	log.Printf("Final error response: MyErrCode=%s Message=%s OriginalErr=%v",
		myErr.ErrCode, myErr.Message, myErr.Err)

	// ユーザーに返すステータスコードの変数を用意
	var statusCode int

	// エラーコード別に該当ステータスコードを格納
	switch myErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case BadRequest, BadPath, BadQuery:
		statusCode = http.StatusBadRequest
	case RequiredAuthorizationHeader, CannotMakeValidator, Unauthorizated:
		statusCode = http.StatusUnauthorized
	case NotMatchUserName:
		statusCode = http.StatusForbidden
	default:
		statusCode = http.StatusInternalServerError
	}

	// ステータスコードをレスポンスのヘッダに格納
	w.WriteHeader(statusCode)

	// ユーザーにエラー内容を送信
	err = json.NewEncoder(w).Encode(myErr)
	if err != nil {
		log.Printf("Fatal: Failed to encode JSON response: %v", err)
	}
}
