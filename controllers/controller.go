package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rainbow777/todolist/api/middlewares"
	"github.com/rainbow777/todolist/myerrors"
	"github.com/rainbow777/todolist/services"
	"github.com/rainbow777/todolist/structure"
)

type MyAppController struct {
	Service services.ServiceInterFace
}

func NewMyAppController(s services.ServiceInterFace) *MyAppController {
	return &MyAppController{Service: s}
}

// 新規タスクをTodoListに追加するメソッド
func (c MyAppController) InsertTaskHandler(w http.ResponseWriter, r *http.Request) {

	var insertData structure.InsertData
	err := json.NewDecoder(r.Body).Decode(&insertData)
	if err != nil {
		err = myerrors.ReqBodyDecodeFailed.Wrap(err, "Failed to decode r.Body")
		myerrors.ErrorHandler(w, err)
		return
	}
	insertData.AuthUserName = middlewares.GetUserName(r.Context())

	// 新規タスクの追加をサービス層に依頼
	addedtask, err := c.Service.InsertTaskService(insertData)
	if err != nil {
		myerrors.ErrorHandler(w, err)
		return
	}

	// レスポンス処理
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(addedtask)
	if err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		return
	}
}

// 指定したIDのタスク情報をTodoListから取得するメソッド
func (c MyAppController) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータからIDを取得
	taskID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		err = myerrors.BadPath.Wrap(err, "invalid path parameter")
		myerrors.ErrorHandler(w, err)
		return
	}

	authUserName := middlewares.GetUserName(r.Context())

	// 対象IDのタスク情報の取得をサービス層に依頼
	task, err := c.Service.GetTaskService(taskID, authUserName)
	if err != nil {
		myerrors.ErrorHandler(w, err)
		return
	}
	// レスポンス処理
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		return
	}
}

// TodoListを取得するメソッド
func (c MyAppController) GetListHandler(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータにpriorityが指定されているかを確認
	getListRequest, err := CheckQueryParam(r)
	if err != nil {
		myerrors.ErrorHandler(w, err)
		return
	}

	getListRequest.AuthUserName = middlewares.GetUserName(r.Context())

	// テータベースからTodoListを取得する処理をサービス層に依頼
	// priorityが空の場合、list内の全タスクが戻ってくる
	todolist, err := c.Service.GetListService(*getListRequest)
	if err != nil {
		myerrors.ErrorHandler(w, err)
		return
	}

	// レスポンス処理
	err = json.NewEncoder(w).Encode(todolist)
	if err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		return
	}
}

// TodoList内のタスク情報を更新するメソッド
func (c MyAppController) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var updateData structure.UpdateData
	var err error
	updateData.TaskID, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		err = myerrors.BadPath.Wrap(err, "TaskID is requred in path parameter")
		myerrors.ErrorHandler(w, err)
		return
	}
	updateData.AuthUserName = middlewares.GetUserName(r.Context())

	// var updateValues structure.Todo
	err = json.NewDecoder(r.Body).Decode(&updateData.UpdateRequest)
	if err != nil {
		err = myerrors.ReqBodyDecodeFailed.Wrap(err, "Your request failed to be loaded.")
		myerrors.ErrorHandler(w, err)
		return
	}

	// データベースの更新処理をサービス層に依頼
	updatedTask, err := c.Service.UpdateTaskService(&updateData)
	if err != nil {
		myerrors.ErrorHandler(w, err)
		return
	}

	// リスポンス処理
	err = json.NewEncoder(w).Encode(updatedTask)
	if err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		return
	}
}

// TodoList内のタスクを削除するメソッド
func (c MyAppController) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	taskID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		err = myerrors.BadPath.Wrap(err, "TaskID is requred in path parameter")
		myerrors.ErrorHandler(w, err)
		return
	}

	username := middlewares.GetUserName(r.Context())

	// データベース内のタスク削除をサービス層に依頼
	err = c.Service.DeleteTaskService(taskID, username)
	if err != nil {
		myerrors.ErrorHandler(w, err)
		return
	}

	// リスポンス処理
	err = json.NewEncoder(w).Encode("The specified task has been deleted.")
	if err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		return
	}
}
