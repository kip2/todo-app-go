package models

import "time"

type UpdateRequest struct {
	ID int `db:"ID"`
}

type DeleteRequest struct {
	ID int `db:"ID"`
}

type Response struct {
	Result string `json:"result"`
}

type RegisterRequest struct {
	Content string `db:"Content"`
}

type Todo struct {
	ID        int       `db:"ID"`
	Content   string    `db:"Content"`
	Done      bool      `db:"Done"`
	CreatedAt time.Time `db:"CreatedAt"`
	UpdatedAt time.Time `db:"UpdatedAt"`
}

// 動作test用の構造体
type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
