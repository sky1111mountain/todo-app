package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rainbow777/todolist/controllers"
	"github.com/rainbow777/todolist/controllers/testdata"
	"github.com/rainbow777/todolist/myerrors"
	"github.com/rainbow777/todolist/structure"
	"github.com/stretchr/testify/assert"
)

func GetMockInInsertCon(test *testdata.InsertedTaskTestData) *controllers.MyAppController {

	mockService := testdata.NewServiceMock()
	mockService.InsertTaskServiceFunc = func(insertData structure.InsertData) (structure.Todo, error) {
		if test.ExpectedError.ErrCode != "" {
			return structure.Todo{}, &test.ExpectedError
		}
		return test.ExpectedTask, nil
	}

	return controllers.NewMyAppController(mockService)
}

func CheckResponseInsertedData(w *httptest.ResponseRecorder, t *testing.T, test *testdata.InsertedTaskTestData) {

	assert.Equal(t, test.Code, w.Code, "Status Code mismatch")
	responseBody := w.Body.Bytes()
	if w.Code == http.StatusOK {

		var returnedTask structure.Todo
		err := json.Unmarshal(responseBody, &returnedTask)
		if err != nil {
			t.Fatalf("Failed to unmarshal responseBody %v", err)
		}

		assert.Equal(t, test.ExpectedTask.ID, returnedTask.ID, "Task ID mismatch")
		assert.Equal(t, test.ExpectedTask.Task, returnedTask.Task, "Task content mismatch")
		assert.Equal(t, test.ExpectedTask.Priority, returnedTask.Priority, "Priority mismatch")
		assert.Equal(t, test.ExpectedTask.Status, returnedTask.Status, "Status mismatch")
		assert.Equal(t, test.ExpectedTask.UserName, returnedTask.UserName, "UserName mismatch")

	} else {

		var returnedError myerrors.MyAppError
		err := json.Unmarshal(responseBody, &returnedError)
		if err != nil {
			t.Fatalf("Failed to unmarshal responseBody %v", err)
		}

		assert.Equal(t, test.ExpectedError.ErrCode, returnedError.ErrCode, "ErrCode mismatch")
		assert.Equal(t, test.ExpectedError.Message, returnedError.Message, "Message mismatch")
	}

}

func GetMockInTaskCon(test *testdata.GetTaskTestData) *controllers.MyAppController {

	mockService := testdata.NewServiceMock()
	mockService.GetTaskServiceFunc = func(id int, username string) (structure.Todo, error) {
		if test.ExpectedError.ErrCode != "" {
			return structure.Todo{}, &test.ExpectedError
		}
		return test.ExpectedTask, nil
	}

	return controllers.NewMyAppController(mockService)
}

func CheckResponseTask(w *httptest.ResponseRecorder, t *testing.T, test *testdata.GetTaskTestData) {
	assert.Equal(t, test.Code, w.Code, "Status Code mismatch")

	responseBody := w.Body.Bytes()

	if w.Code == http.StatusOK {

		var returnedTask structure.Todo
		err := json.Unmarshal(responseBody, &returnedTask)
		if err != nil {
			t.Fatalf("Failed to unmarshal responseBody %v", err)
		}

		assert.Equal(t, test.ExpectedTask.ID, returnedTask.ID, "Task ID mismatch")
		assert.Equal(t, test.ExpectedTask.Task, returnedTask.Task, "Task content mismatch")
		assert.Equal(t, test.ExpectedTask.Priority, returnedTask.Priority, "Priority mismatch")
		assert.Equal(t, test.ExpectedTask.Status, returnedTask.Status, "Status mismatch")
		assert.Equal(t, test.ExpectedTask.UserName, returnedTask.UserName, "UserName mismatch")

	} else {

		var returnedError myerrors.MyAppError
		err := json.Unmarshal(responseBody, &returnedError)
		if err != nil {
			t.Fatalf("Failed to unmarshal responseBody %v", err)
		}

		assert.Equal(t, test.ExpectedError.ErrCode, returnedError.ErrCode, "ErrCode mismatch")
		assert.Equal(t, test.ExpectedError.Message, returnedError.Message, "Message mismatch")
	}

}

func GetMockInListCon(test *testdata.GetListTestData) *controllers.MyAppController {
	mockService := testdata.NewServiceMock()
	mockService.GetListServiceFunc = func(getListRequest structure.GetListRequest) ([]structure.Todo, error) {
		if test.ExpectedError.ErrCode != "" {
			return []structure.Todo{}, &test.ExpectedError

		}
		return test.ExpectedList, nil
	}

	return controllers.NewMyAppController(mockService)
}

func CheckResponseList(w *httptest.ResponseRecorder, t *testing.T, test *testdata.GetListTestData) {
	assert.Equal(t, test.Code, w.Code, "Status Code mismatch")

	responseBody := w.Body.Bytes()

	if w.Code == http.StatusOK {

		var returnedList []structure.Todo
		err := json.Unmarshal(responseBody, &returnedList)
		if err != nil {
			t.Fatalf("Failed to unmarshal responseBody %v", err)
		}
		for i, returnedTask := range returnedList {
			assert.Equal(t, test.ExpectedList[i].ID, returnedTask.ID, "Task ID mismatch")
			assert.Equal(t, test.ExpectedList[i].Task, returnedTask.Task, "Task content mismatch")
			assert.Equal(t, test.ExpectedList[i].Priority, returnedTask.Priority, "Priority mismatch")
			assert.Equal(t, test.ExpectedList[i].Status, returnedTask.Status, "Status mismatch")
			assert.Equal(t, test.ExpectedList[i].UserName, returnedTask.UserName, "UserName mismatch")
		}

	} else {

		var returnedError myerrors.MyAppError
		err := json.Unmarshal(responseBody, &returnedError)
		if err != nil {
			t.Fatalf("Failed to unmarshal responseBody %v", err)
		}

		assert.Equal(t, test.ExpectedError.ErrCode, returnedError.ErrCode, "ErrCode mismatch")
		assert.Equal(t, test.ExpectedError.Message, returnedError.Message, "Message mismatch")
	}

}

func CheckResponseUpdateTask(w *httptest.ResponseRecorder, t *testing.T, test *testdata.UpdateTaskData) {
	assert.Equal(t, test.Code, w.Code, "Status Code mismatch")

	responseBody := w.Body.Bytes()

	if w.Code == http.StatusOK {
		var actualBody structure.Todo
		err := json.Unmarshal(responseBody, &actualBody)
		if err != nil {
			t.Fatalf("Failed to Unmarshal actual body%v", err)
		}
		assert.Equal(t, test.ReturnValue.ID, actualBody.ID, "TaskID mismatch")
		assert.Equal(t, test.ReturnValue.Task, actualBody.Task, "Task content mismatch")
		assert.Equal(t, test.ReturnValue.Priority, actualBody.Priority, "Priority mismatch")
		assert.Equal(t, test.ReturnValue.Status, actualBody.Status, "Status mismatch")
		assert.Equal(t, test.ReturnValue.UserName, actualBody.UserName, "UserName mismatch")
	} else {
		var actualError myerrors.MyAppError
		err := json.Unmarshal(responseBody, &actualError)
		if err != nil {
			t.Fatalf("Failed to Unmarshal actual body%v", err)
		}
		assert.Equal(t, test.ExpectedError.ErrCode, actualError.ErrCode, "ErrCode mismatch")
		assert.Equal(t, test.ExpectedError.Message, actualError.Message, "Message mismathc")
	}
}

func GetMockInUpdateCon(test *testdata.UpdateTaskData) *controllers.MyAppController {
	mockService := testdata.NewServiceMock()
	mockService.UpdateTaskServiceFunc = func(updateData *structure.UpdateData) (structure.Todo, error) {
		if test.ExpectedError.ErrCode != "" {
			return structure.Todo{}, &test.ExpectedError
		}
		return test.ReturnValue, nil
	}

	return controllers.NewMyAppController(mockService)
}

func CheckResponseDeleteHandler(w *httptest.ResponseRecorder, t *testing.T, test *testdata.DeleteTestData) {

	assert.Equal(t, test.Code, w.Code, "StatusCode is unmatched")

	if w.Code == http.StatusOK {
		responseBody := w.Body.Bytes()
		var responseMessage string
		err := json.Unmarshal(responseBody, &responseMessage)
		if err != nil {
			t.Fatalf("Failed to Unmarshal responseBody %v", err)
		}
		assert.Equal(t, test.ExpectedRes, responseMessage, "Response Massage is unmatched")
	} else {
		responseBody := w.Body.Bytes()
		var responseErr myerrors.MyAppError
		err := json.Unmarshal(responseBody, &responseErr)
		if err != nil {
			t.Fatalf("Failed to Unmarshal responseBody %v", err)
		}
		assert.Equal(t, test.ExpectedRes, responseErr.Message, "Err Massage is unmatched")
	}

}
