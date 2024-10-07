package main

import (
	"encoding/json"
	"net/http"
	"todoApp/internal/db"
	"todoApp/internal/models"
)

func main() {
	// リスト(todo)の一覧を取得するハンドラのバインド
	http.HandleFunc("/todos", todosHandler)
	// リクエストしたデータを登録するハンドラのバインド
	http.HandleFunc("/register", registerHandler)

	http.ListenAndServe(":8080", nil)
}

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

/*
リクエストに含まれるデータをDBに登録するハンドラ
*/
func registerHandler(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// 成功の場合のResponseデータを作成
	response := models.Response{
		Result: "SUCCESS",
	}

	// DBへの登録処理を行う
	err := db.Insert(req)
	// DB登録処理が失敗なら、エラーメッセージを格納したResponseデータに変更
	if err != nil {
		response = models.Response{
			Result: "Data register error.",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

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
