package main

import (
	"fmt"
	"time"
	"todoApp/db"
	errorpkg "todoApp/error"
)

type Todo struct {
	ID        int        `db:"ID"`
	Content   string     `db:"Content"`
	Done      bool       `db:"Done"`
	Until     *time.Time `db:"Until"`
	CreatedAt time.Time  `db:"CreatedAt"`
	UpdatedAt time.Time  `db:"UpdatedAt"`
	DeletedAt *time.Time `db:"DeletedAt"`
}

func main() {

	envVar := "DATABASE"
	db := db.CreateDBConnection(envVar)
	defer db.Close()

	var todo Todo
	err := db.Get(&todo, "SELECT * FROM todos WHERE id=?", 1)
	errorpkg.CheckError(err)

	fmt.Printf("Todo: %+v\n", todo)
}
