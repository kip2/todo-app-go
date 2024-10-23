package main

import (
	"encoding/json"
	"net/http"
	"todoApp/internal/db"
	"todoApp/internal/models"

	_ "todoApp/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Go Todo API
// @version 1.0
// @description TodoアプリのバックエンドAPIです。
// @host localhost:8080
// @BasePath /api
func main() {
	// swaggerドキュメントの設定
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	// リスト(todo)の一覧を取得するハンドラのバインド
	http.HandleFunc("/api/todos", todosHandler)
	// リクエストしたデータを登録するハンドラのバインド
	http.HandleFunc("/api/register", registerHandler)
	// リクエストしたデータを削除するハンドラのバインド
	http.HandleFunc("/api/delete", deleteHandler)
	// リクエストしたデータを更新するハンドラのバインド
	http.HandleFunc("/api/update", updateHandler)

	// 画面を返すエンドポイント
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)
}

// updateHandler godoc
// @Summary IDに紐づくTodoのDoneを更新する
// @Description IDに紐づいているTodoのステータスを更新する。呼び出す度に、Doneのステータスをトグルする。
// @Tags todos
// @Accept json
// @Produce json
// @Param updateRequest body models.UpdateRequest true "Update Todo"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /api/update [put]
/*
リクエストで指定したIDのデータの状態を更新するハンドラ
*/
func updateHandler(w http.ResponseWriter, r *http.Request) {
	var req models.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := models.Response{
		Result: "SUCCESS",
	}

	// データの更新を行う
	err := db.Update(req)

	if err != nil {
		response = models.Response{
			Result: "Data update error.",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handler godoc
// @Summary IDに紐づくTodoを削除する
// @Description IDに紐づくTodoをDBから削除する
// @Tags todos
// @Accept json
// @Produce json
// @Param deleteRequest body models.DeleteRequest true "Delete Todo"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /api/delete [delete]
/*
リクエストで指定したIDのデータを削除するハンドラ
*/
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	var req models.DeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := models.Response{
		Result: "SUCCESS",
	}

	// DBの削除処理を行う
	err := db.Delete(req)

	if err != nil {
		response = models.Response{
			Result: "Data delete error.",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// registerhandler godoc
// @Summary TodoをDBに登録する
// @Description リクエストに含まれるTodoのデータをDBに登録する
// @Tags todos
// @Accept json
// @Produce json
// @Param registerRequest body models.RegisterRequest true "Register Todo"
// @Success 200 {object} models.Todo
// @Failure 400 {object} models.Todo
// @Failure 500 {object} models.Todo
// @Router /api/register [post]
/*
リクエストに含まれるデータをDBに登録するハンドラ
*/
func registerHandler(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// DBへの登録処理を行う
	insertId, err := db.Insert(req)
	// DB登録処理が失敗なら、エラーメッセージを格納したResponseデータに変更
	if err != nil {
		errResponse := models.Response{
			Result: "Data register error.",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	// 登録成功時のレスポンス
	response := db.SelectById(int(insertId))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// todosHandler godoc
// @Summary Todoのリストを取得する
// @Description DBに登録されているすべてのTodoをリストで取得する
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {object} models.Todo
// @Failure 400 {object} models.Todo
// @Failure 500 {object} models.Todo
// @Router /api/todos [get]
/*
DBからTodoの全リストを取得して、レスポンスするハンドラ
*/
func todosHandler(w http.ResponseWriter, r *http.Request) {
	todos := db.SelectAll()

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, "Filed to encode users", http.StatusInternalServerError)
		return
	}
}
