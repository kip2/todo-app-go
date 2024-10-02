package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	name := "David"
	insert(name)
}

/*
データをDBにINSERTする
*/
func insert(name string) {
	envVar := "DATABASE"
	db := createDBConnection(envVar)
	defer db.Close()

	result, err := db.Exec("INSERT INTO users (name) VALUES (?)", name)
	checkError(err)

	lastInsertID, err := result.LastInsertId()
	checkError(err)

	fmt.Printf("Inserted user with ID: %d\n", lastInsertID)
}

/*
エラーハンドリング用のコード(マクロ代わり)
*/
func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

/*
MySQLのDBコネクションを作成する関数
*/
func createDBConnection(envVar string) *sqlx.DB {
	dsn := loadEnv(envVar)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

/*
.envファイルから特定のキーに紐づく値を取得する関数
*/
func loadEnv(key string) string {
	// .envに定義した環境変数をロード
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// 環境変数から値を取得
	env := os.Getenv(key)

	return env
}
