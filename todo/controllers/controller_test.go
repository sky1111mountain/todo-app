package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/rainbow777/todolist/api/middlewares"
	"github.com/rainbow777/todolist/controllers"
	"github.com/rainbow777/todolist/controllers/testdata"
)

func TestInsertTaskHandler(t *testing.T) {
	for _, test := range testdata.InsertTaskTestCases {
		t.Run(test.Name, func(t *testing.T) {

			con := GetMockInInsertCon(&test)

			jsonBody, err := json.Marshal(test.InsertTask)
			if err != nil {
				t.Fatalf("Failed to encode a testdata to JSON %v", err)
			}

			// ハンドラに渡すリクエストを自作
			req := httptest.NewRequest(http.MethodPost, "/todo/insert", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			req = middlewares.SetUserName(req, "rainbow777")

			// テスト用レスポンスを生成
			w := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/todo/insert", con.InsertTaskHandler).Methods(http.MethodPost)
			router.ServeHTTP(w, req)

			CheckResponseInsertedData(w, t, &test)
		})
	}
}

func TestGetTaskHandler(t *testing.T) {
	for _, test := range testdata.GetTaskTestCases {
		t.Run(test.Name, func(t *testing.T) {

			con := GetMockInTaskCon(&test)

			// ハンドラに渡すリクエストを自作
			url := fmt.Sprintf("http://localhost:8080/todo/gettask/%s", test.ID)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			req.Header.Set("Content-Type", "application/json")

			req = middlewares.SetUserName(req, "rainbow777")

			w := httptest.NewRecorder()

			router := mux.NewRouter()
			// ハンドラのテストをしたいので正規表現を削除
			router.HandleFunc("/todo/gettask/{id}", con.GetTaskHandler).Methods(http.MethodGet)
			router.ServeHTTP(w, req)

			CheckResponseTask(w, t, &test)
		})
	}
}

func TestGetListHandler(t *testing.T) {
	for _, test := range testdata.GetListTestCases {
		t.Run(test.Name, func(t *testing.T) {

			con := GetMockInListCon(&test)

			url := fmt.Sprintf("http://localhost:8080/todo/getlist%s", test.QueryParam)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			req.Header.Set("Content-Type", "application/json")

			req = middlewares.SetUserName(req, "rainbow777")

			w := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/todo/getlist", con.GetListHandler).Methods(http.MethodGet)
			router.ServeHTTP(w, req)

			CheckResponseList(w, t, &test)
		})
	}
}

func TestUpdateTaskHandler(t *testing.T) {
	for _, test := range testdata.UpdateTaskTestCases {
		t.Run(test.Name, func(t *testing.T) {

			con := GetMockInUpdateCon(&test)

			jsonBody, err := json.Marshal(test.UpdateRequest)
			if err != nil {
				t.Fatalf("Failed to encode a testdata to JSON %v", err)
			}

			// ハンドラに渡すリクエストを自作
			url := fmt.Sprintf("/todo/update/%d", test.TaskID)
			req := httptest.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			req = middlewares.SetUserName(req, "rainbow777")

			w := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/todo/update/{id}", con.UpdateTaskHandler).Methods(http.MethodPatch)
			router.ServeHTTP(w, req)

			CheckResponseUpdateTask(w, t, &test)
		})
	}
}

func TestDeleteTaskHandler(t *testing.T) {
	for _, test := range testdata.DeleteTestCases {
		t.Run(test.Name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodDelete, test.URL, nil)
			req.Header.Set("Content-Type", "application/json")
			req = middlewares.SetUserName(req, "rainbow777")

			w := httptest.NewRecorder()

			mockService := testdata.NewServiceMock()
			con := controllers.NewMyAppController(mockService)

			router := mux.NewRouter()
			router.HandleFunc("/todo/delete/{id}", con.DeleteTaskHandler).Methods(http.MethodDelete)
			router.ServeHTTP(w, req)

			CheckResponseDeleteHandler(w, t, &test)
		})
	}
}
