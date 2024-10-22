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
指定したIDのDoneを更新する。
*/
func Update(data models.UpdateRequest) error {
	db := CreateDBConnection(envVar)
	defer db.Close()

	// クエリ実行
	result, err := db.Exec("UPDATE todos SET Done = IF(Done = 1, 0, 1) WHERE id = ?", data.ID)
	if err != nil {
		return fmt.Errorf("failed to execute update: %v", err)
	}

	// 実際に更新した行数を取得する
	rowsAffected, err := result.RowsAffected()
	// 更新した行数取得に失敗した場合のエラーを返す
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows: %v", err)
	}

	// 更新した行数が0ならエラーを返す
	if rowsAffected == 0 {
		return fmt.Errorf("no rows updated, ID %d not found", data.ID)
	}

	// 正常終了のため、nilを返す
	return nil
}

/*
指定されたIDのデータを削除する
*/
func Delete(data models.DeleteRequest) error {
	db := CreateDBConnection(envVar)
	defer db.Close()

	// クエリ実行
	result, err := db.Exec("DELETE FROM todos WHERE ID = ?", data.ID)
	if err != nil {
		return fmt.Errorf("failed to execute delete: %v", err)
	}

	// 実際に削除した行数を取得する
	rowsAffected, err := result.RowsAffected()
	// 削除行数取得に失敗した場合のエラーを返す
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows: %v", err)
	}

	// 削除した行数が0ならエラーを返す
	if rowsAffected == 0 {
		return fmt.Errorf("no rows deleted, ID %d not found", data.ID)
	}

	// 正常終了のため、nilを返す
	return nil
}

/*
指定したIDのデータをDBにINSERTする
*/
func InsertById(id int, data models.RegisterRequest) error {
	db := CreateDBConnection(envVar)
	defer db.Close()

	_, err := db.Exec("INSERT INTO todos (ID, Content) VALUES (?, ?)", id, data.Content)
	if err != nil {
		return fmt.Errorf("failed to execute insert: %v", err)
	}

	return nil
}

/*
データをDBにINSERTする
*/
func Insert(data models.RegisterRequest) (int64, error) {
	db := CreateDBConnection(envVar)
	defer db.Close()

	result, err := db.Exec("INSERT INTO todos (Content) VALUES (?)", data.Content)
	if err != nil {
		return 0, fmt.Errorf("failed to execute insert: %v", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve last insert ID: %v", err)
	}

	return lastInsertID, nil
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
