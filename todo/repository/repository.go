package repository

import (
	"context"
	"database/sql"

	"github.com/rainbow777/todolist/myerrors"
	"github.com/rainbow777/todolist/structure"
)

type RepositoryInterFace interface {
	InsertTaskDB(insertData structure.InsertData) (structure.Todo, error)
	GetTaskDB(taskID int, username string) (structure.Todo, error)
	GetTodolistDB(getListRequest structure.GetListRequest) ([]structure.Todo, error)
	UpdateTaskDB(updateData *structure.UpdateData) (structure.Todo, error)
	DeleteTaskDB(taskID int, username string) error
}

type MyAppRepository struct {
	DB *sql.DB
}

func NewMyAppRepository(db *sql.DB) *MyAppRepository {
	return &MyAppRepository{DB: db}
}

// データベースに新規タスクを追加する関数
func (r *MyAppRepository) InsertTaskDB(insertData structure.InsertData) (structure.Todo, error) {
	// クエリ文の生成
	insertQuery := `insert into todolist (task, priority, status, username, created_at) values
		(?,?,?,?,now());`

	// データベースに新規タスクを追加する処理
	result, err := r.DB.Exec(insertQuery, insertData.Task, insertData.Priority,
		insertData.Status, insertData.AuthUserName)

	if err != nil {
		wrappedErr := myerrors.InsertFailed.Wrap(err, "Error occurred in db.Exec in InsertTaskDB")
		return structure.Todo{}, wrappedErr
	}

	// データベースに新規タスクが追加されたかを確認
	err = CheckAffectedRows(result)
	if err != nil {
		return structure.Todo{}, err
	}

	return GetInsertedTask(r, result, insertData.AuthUserName)
}

// 指定したIDのタスクをデータベースから取得する関数
func (r *MyAppRepository) GetTaskDB(taskID int, username string) (structure.Todo, error) {
	// クエリ文の生成
	selectQuery := `SELECT id, task, priority, status, username, created_at FROM todolist WHERE id = ? AND username = ?;`
	// クエリ文を実行して、タスクの情報を取得
	row := r.DB.QueryRow(selectQuery, taskID, username)

	var task structure.Todo
	// データベースから取得したタスク情報をGo構造体の各フィールドに格納する
	err := row.Scan(&task.ID, &task.Task, &task.Priority, &task.Status, &task.UserName, &task.CreatedAt)

	if err != nil {
		wrappedErr := myerrors.ScanFailed.Wrap(err, "Error occurred in row.Scan in GetTaskDB")
		return structure.Todo{}, wrappedErr
	}
	return task, nil
}

// Todoリストを取得してGo構造体に格納する関数
// 優先度の指定が可能
func (r *MyAppRepository) GetTodolistDB(getListRequest structure.GetListRequest) ([]structure.Todo, error) {
	// 自作関数getRowsでTodoリストを取得
	rows, err := GetRows(r.DB, getListRequest)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todolist []structure.Todo
	// 自作関数scanRows関数でTodoリストをGoの構造体に格納
	todolist, err = ScanRows(rows)
	if err != nil {
		return nil, err
	}
	return todolist, nil
}

// 既存のタスクの情報を変更する関数
func (r *MyAppRepository) UpdateTaskDB(updateData *structure.UpdateData) (todo structure.Todo, err error) {

	tx, err := r.DB.BeginTx(context.Background(), nil)
	if err != nil {
		return structure.Todo{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query, arguments, err := MakeQueryAndArgs(updateData)
	if err != nil {
		return structure.Todo{}, err
	}

	err = ExecuteUpdate(tx, query, arguments)
	if err != nil {
		return structure.Todo{}, err
	}

	todo, err = GetUpdatedTask(tx, updateData.TaskID)
	if err != nil {
		return structure.Todo{}, err
	}

	err = tx.Commit()
	if err != nil {
		return structure.Todo{}, err
	}

	return todo, nil
}

// 指定されたIDのタスクを削除する関数
func (r *MyAppRepository) DeleteTaskDB(taskID int, username string) error {
	//　クエリ文の生成
	deleteQuery := `delete from todolist where id=? AND username = ?`

	// クエリ文の実行
	result, err := r.DB.Exec(deleteQuery, taskID, username)
	if err != nil {
		wrappedErr := myerrors.TaskDeleteFailed.Wrap(err, "Failed to delete task")
		return wrappedErr
	}

	// データベース内が変更されたかを確認
	err = CheckAffectedRows(result)
	if err != nil {
		return err
	}

	return nil
}
