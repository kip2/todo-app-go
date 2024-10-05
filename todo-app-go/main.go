package main

import (
	"encoding/json"
	"net/http"
	"todoApp/internal/db"
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/todos", todosHandler)
	http.HandleFunc("/register", registerHandler)

	http.ListenAndServe(":8080", nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var req User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := Response{
		Message: "Hello, " + req.Name,
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
