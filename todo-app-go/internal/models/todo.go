package models

import "time"

type DeleteRequest struct {
	ID int `db:"ID"`
}

type RegisterRequest struct {
	Content string    `db:"Content"`
	Until   time.Time `db:"Until"`
}

type Response struct {
	Result string `json:"result"`
}

type Todo struct {
	ID        int        `db:"ID"`
	Content   string     `db:"Content"`
	Done      bool       `db:"Done"`
	Until     *time.Time `db:"Until"`
	CreatedAt time.Time  `db:"CreatedAt"`
	UpdatedAt time.Time  `db:"UpdatedAt"`
}

// 動作test用の構造体
type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
