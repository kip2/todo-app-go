package db

import (
	"fmt"
	"log"
	"todoApp/env"
	errorpkg "todoApp/error"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
データをDBにINSERTする
*/
func Insert(name string) {
	envVar := "DATABASE"
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
