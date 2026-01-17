package services

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/rainbow777/todolist/myerrors"
	"github.com/rainbow777/todolist/repository"
	"github.com/rainbow777/todolist/structure"
)

type ServiceInterFace interface {
	InsertTaskService(insertData structure.InsertData) (structure.Todo, error)
	GetTaskService(taskID int, username string) (structure.Todo, error)
	GetListService(getListRequest structure.GetListRequest) ([]structure.Todo, error)
	UpdateTaskService(updateData *structure.UpdateData) (structure.Todo, error)
	DeleteTaskService(taskID int, username string) error
}

type MyAppService struct {
	Repository repository.RepositoryInterFace
}

// サービス層構造体のコンストラクタ
func NewMyAppService(r repository.RepositoryInterFace) *MyAppService {
	return &MyAppService{Repository: r}
}

// レポジトリ層に新規タスク追加を依頼する関数
func (s MyAppService) InsertTaskService(insertData structure.InsertData) (structure.Todo, error) {

	// リクエストの内容を確認
	err := ValidationRequest(insertData)
	if err != nil {
		return structure.Todo{}, err
	}

	// データベース処理の呼び出し
	addedTask, err := s.Repository.InsertTaskDB(insertData)
	if err != nil {
		err = myerrors.InsertDataFailed.Wrap(err, "Service Error: Failed to insert task to DB")
		return addedTask, err
	}

	// データベースに追加したデータをハンドラ層に返す
	return addedTask, nil
}

// 指定したIDのタスク取得をレポジトリ層に依頼する関数
func (s MyAppService) GetTaskService(taskID int, username string) (structure.Todo, error) {

	// データベース処理の呼び出し
	task, err := s.Repository.GetTaskDB(taskID, username)
	if err != nil {
		// データがなかった場合のエラー処理
		if errors.Is(err, sql.ErrNoRows) {
			Message := fmt.Sprintf("Service Error:  TaskID %d is no data", taskID)
			return structure.Todo{}, myerrors.NAData.Wrap(err, Message)
		}
		// レポジトリ層からエラーが返ってきた場合のエラー処理
		return structure.Todo{}, myerrors.GetDataFailed.Wrap(err, "Failed to get task")
	}
	return task, nil
}

// TodoListの取得をリポジトリ層に依頼する関数
func (s MyAppService) GetListService(getListRequest structure.GetListRequest) ([]structure.Todo, error) {

	// リポジトリ層にTodoListの取得を依頼
	gotTodoList, err := s.Repository.GetTodolistDB(getListRequest)
	if err != nil {
		return []structure.Todo{}, myerrors.GetDataFailed.Wrap(err, "Error occured in GetTodoListDB")
	}

	return gotTodoList, nil
}

// タスク情報の更新をレポジトリ層に依頼する関数
func (s MyAppService) UpdateTaskService(updateData *structure.UpdateData) (structure.Todo, error) {
	// 更新する項目を確認し、更新内容を取得
	if updateData.UpdateRequest.Task == nil &&
		updateData.UpdateRequest.Priority == nil && updateData.UpdateRequest.Status == nil {
		return structure.Todo{}, myerrors.NoUpdateColumn.Wrap(errors.New("No update colomn"),
			"At least one colomn required")
	}

	if updateData.UpdateRequest.Priority != nil {
		if *updateData.UpdateRequest.Priority != "high" &&
			*updateData.UpdateRequest.Priority != "medium" && *updateData.UpdateRequest.Priority != "low" {
			errMessage := fmt.Sprintf("Priority require high or medium or low not %v", *updateData.UpdateRequest.Priority)
			return structure.Todo{}, myerrors.BadRequest.Wrap(errors.New("invalid priority"), errMessage)
		}
	}

	// レポジトリ層にタスクの更新を依頼
	updatedTask, err := s.Repository.UpdateTaskDB(updateData)
	if err != nil {
		return structure.Todo{}, myerrors.UpdateDataFailed.Wrap(err, "Error occurred in UpdateTaskService")
	}

	return updatedTask, nil
}

// 指定されたIDのタスク削除をレポジトリ層に依頼する関数
func (s MyAppService) DeleteTaskService(taskID int, username string) error {

	// リポジトリ層にタスクの削除を依頼
	err := s.Repository.DeleteTaskDB(taskID, username)
	if err != nil {
		return myerrors.DeleteDataFailed.Wrap(err, "Error occurred in DeleteTaskDB")
	}

	return nil
}
