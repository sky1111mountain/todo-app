package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/rainbow777/todolist/myerrors"
	"github.com/rainbow777/todolist/structure"
)

// データベースが変更されたかを確認する関数
func CheckAffectedRows(result sql.Result) error {
	affectedRows, err := result.RowsAffected()

	// RowsAffected()関数のエラー処理
	if err != nil {
		wrappedErr := myerrors.NoRowsAffected.Wrap(err, "Failed to check database changes were complete")
		return wrappedErr
	}

	// データベースの変更を確認できなかった場合のエラー処理
	if affectedRows < 1 {
		wrappedErr := myerrors.UnChangeRows.Wrap(errors.New("No affectedRows"), "Database has not changed")
		return wrappedErr
	}

	return nil
}

func GetInsertedTask(r *MyAppRepository, result sql.Result, authUserName string) (structure.Todo, error) {
	// 追加した新規タスクのIDを取得
	id, err := result.LastInsertId()
	if err != nil {
		return structure.Todo{}, myerrors.InsertFailed.Wrap(err, "Error occurred in LastInsertId")
	}
	addedTaskID := int(id)

	// データベースから追加したタスクを取得して、サービス層に返す
	insertedTask, err := r.GetTaskDB(addedTaskID, authUserName)
	if err != nil {
		wrappedErr := myerrors.GetTaskFailed.Wrap(err, "Failed to get last inserted task")
		return structure.Todo{}, wrappedErr
	}
	return insertedTask, nil
}

// Todoリストをデータベースから取得する関数
func GetRows(db *sql.DB, getListRequest structure.GetListRequest) (*sql.Rows, error) {

	baseQuery := "SELECT id, task, priority, status, username, created_at FROM todolist WHERE username = ?"
	args := []any{getListRequest.AuthUserName}

	if getListRequest.Status != "all" {
		baseQuery += " AND status = ?"
		args = append(args, getListRequest.Status)
	}

	if len(getListRequest.Priorities) > 0 {
		questionMarks := make([]string, len(getListRequest.Priorities))
		for i := range questionMarks {
			questionMarks[i] = "?"
		}
		placeholders := strings.Join(questionMarks, ",")

		baseQuery += fmt.Sprintf(" AND priority IN (%s)", placeholders)

		for _, p := range getListRequest.Priorities {
			args = append(args, p)
		}
	}

	baseQuery += " ORDER BY id DESC LIMIT ? OFFSET ?"
	args = append(args, getListRequest.Limit, getListRequest.Offset)

	rows, err := db.Query(baseQuery, args...)
	if err != nil {
		return nil, myerrors.GetListfailed.Wrap(err, "Error occured in db.Query in GetRows func")
	}

	return rows, nil
}

// TodoリストをGo構造体に格納する関数
func ScanRows(rows *sql.Rows) ([]structure.Todo, error) {
	// 複数のタスクデータを格納するため、構造体のスライスを用意
	var todolist []structure.Todo
	// rows内のデータがなくなるまで実行
	for rows.Next() {
		var task structure.Todo
		// タスクの情報を、構造体の各フィールドに格納
		err := rows.Scan(&task.ID, &task.Task, &task.Priority,
			&task.Status, &task.UserName, &task.CreatedAt)
		// エラー処理
		if err != nil {
			return nil, myerrors.ScanFailed.Wrap(err, "Error occurred in rows.Scan in ScanRows Func")
		}

		// タスク情報を格納した構造体をスライスに格納
		todolist = append(todolist, task)
	}

	if err := rows.Err(); err != nil {
		return nil, myerrors.ScanFailed.Wrap(err, "rows iteration error")
	}

	return todolist, nil
}

func MakeQueryAndArgs(updateData *structure.UpdateData) (string, []any, error) {

	var columns []string
	var values []any

	if updateData.UpdateRequest.Task != nil {
		columns = append(columns, "task = ?")
		values = append(values, *updateData.UpdateRequest.Task)
	}

	if updateData.UpdateRequest.Priority != nil {
		columns = append(columns, "priority = ?")
		values = append(values, *updateData.UpdateRequest.Priority)
	}

	if updateData.UpdateRequest.Status != nil {
		columns = append(columns, "status = ?")
		values = append(values, *updateData.UpdateRequest.Status)
	}

	if len(columns) == 0 {
		return "", nil, myerrors.NoUpdateColumn.Wrap(errors.New("No Column to update"),
			"Update column is requred")
	}

	query := fmt.Sprintf("UPDATE todolist SET %s WHERE id = ? AND username = ?",
		strings.Join(columns, ", "))

	arguments := append(values, updateData.TaskID, updateData.AuthUserName)

	return query, arguments, nil

}

func ExecuteUpdate(tx *sql.Tx, query string, arguments []any) error {
	result, err := tx.Exec(query, arguments...)
	if err != nil {
		wrappedErr := myerrors.TaskUpdateFailed.Wrap(err, "Error occurred in tx.Exec")
		return wrappedErr
	}

	err = CheckAffectedRows(result)
	if err != nil {
		return err
	}

	return nil
}

func GetUpdatedTask(tx *sql.Tx, taskID int) (structure.Todo, error) {
	getQuery := "SELECT id, task, priority, status, username, created_at FROM todolist WHERE id = ?"
	row := tx.QueryRow(getQuery, taskID)

	var updatedTask structure.Todo
	err := row.Scan(&updatedTask.ID, &updatedTask.Task, &updatedTask.Priority,
		&updatedTask.Status, &updatedTask.UserName, &updatedTask.CreatedAt)

	if err != nil {
		return structure.Todo{}, myerrors.ScanFailed.Wrap(err, "Error occurred in row.Scan")
	}

	return updatedTask, nil
}
