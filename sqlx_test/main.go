package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func main() {
	// dbとのコネクションを作成
	db := createDBConnection()

	// dbコネクションを閉じるためのデコンストラクタ
	defer db.Close()

	// クエリを実行して結果を取得
	var users []User
	err := db.Select(&users, "SELECT id, name FROM users WHERE id=?", 1)
	if err != nil {
		log.Fatalln(err)
	}

	// クエリした結果を表示して確認
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s\n", user.ID, user.Name)
	}
}

/*
DBコネクションを作成する関数
*/
func createDBConnection() *sqlx.DB {
	dsn := loadEnv("DATABASE")

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
