package db

import (
	"fmt"
	"log"
	"time"
	"todoApp/env"
	errorpkg "todoApp/error"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const envVar = "DATABASE"

type Todo struct {
	ID        int        `db:"ID"`
	Content   string     `db:"Content"`
	Done      bool       `db:"Done"`
	Until     *time.Time `db:"Until"`
	CreatedAt time.Time  `db:"CreatedAt"`
	UpdatedAt time.Time  `db:"UpdatedAt"`
	DeletedAt *time.Time `db:"DeletedAt"`
}

/*
すべてのTodoをDBから取得する
*/
func SelectAll() []Todo {
	db := CreateDBConnection(envVar)
	defer db.Close()

	var todo []Todo
	err := db.Select(&todo, "SELECT * FROM todos")
	errorpkg.CheckError(err)

	return todo
}

/*
指定したIDのTodoをDBから取得する
*/
func SelectById(id int) Todo {
	db := CreateDBConnection(envVar)
	defer db.Close()

	var todo Todo
	err := db.Get(&todo, "SELECT * FROM todos WHERE id=?", id)
	errorpkg.CheckError(err)

	return todo
}

/*
データをDBにINSERTする(test用)
*/
func Insert(name string) {
	db := CreateDBConnection(envVar)
	defer db.Close()

	result, err := db.Exec("INSERT INTO users (name) VALUES (?)", name)
	errorpkg.CheckError(err)

	lastInsertID, err := result.LastInsertId()
	errorpkg.CheckError(err)

	fmt.Printf("Inserted user with ID: %d\n", lastInsertID)
}

/*
MySQLのDBコネクションを作成する関数
*/
func CreateDBConnection(envVar string) *sqlx.DB {
	dsn := env.LoadEnv(envVar)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
