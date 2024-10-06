package db

import (
	"fmt"
	"log"
	"todoApp/internal/env"
	errorpkg "todoApp/internal/error"
	"todoApp/internal/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const envVar = "DATABASE"

/*
すべてのTodoをDBから取得する
*/
func SelectAll() []models.Todo {
	db := CreateDBConnection(envVar)
	defer db.Close()

	var todo []models.Todo
	err := db.Select(&todo, "SELECT * FROM todos")
	errorpkg.CheckError(err)

	return todo
}

/*
データをDBにINSERTする
*/
func Insert(data models.RegisterRequest) (int64, error) {
	db := CreateDBConnection(envVar)
	defer db.Close()

	result, err := db.Exec("INSERT INTO todos (Content, Until) VALUES (?, ?)", data.Content, data.Until)
	errorpkg.CheckError(err)

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, err
}

/*
指定したIDのTodoをDBから取得する
*/
func SelectById(id int) models.Todo {
	db := CreateDBConnection(envVar)
	defer db.Close()

	var todo models.Todo
	err := db.Get(&todo, "SELECT * FROM todos WHERE id=?", id)
	errorpkg.CheckError(err)

	return todo
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

/*
データをDBにINSERTする(test用)
*/
func InsertUserTestData(name string) {
	db := CreateDBConnection(envVar)
	defer db.Close()

	result, err := db.Exec("INSERT INTO users (name) VALUES (?)", name)
	errorpkg.CheckError(err)

	lastInsertID, err := result.LastInsertId()
	errorpkg.CheckError(err)

	fmt.Printf("Inserted user with ID: %d\n", lastInsertID)
}
