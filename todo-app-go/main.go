package main

import (
	"encoding/json"
	"net/http"
	"todoApp/internal/db"
	"todoApp/internal/models"
)

func main() {
	http.HandleFunc("/todos", todosHandler)
	http.HandleFunc("/register", registerHandler)

	http.ListenAndServe(":8080", nil)
}

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
	_, err := db.Insert(req)
	// DB登録処理が失敗なら、エラーメッセージを格納したResponseデータに変更
	if err != nil {
		response = models.Response{
			Result: "Data register error.",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func todosHandler(w http.ResponseWriter, r *http.Request) {
	todos := db.SelectAll()

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, "Filed to encode users", http.StatusInternalServerError)
		return
	}
}
