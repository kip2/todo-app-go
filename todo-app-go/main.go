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

	// 受け取ったデータをDBに登録する処理

	// responseデータとして登録したことを返す処理
	// todo: DB登録がうまく言ったかどうかを判定する
	response := models.Response{
		Result: "Hello, " + req.Content,
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
