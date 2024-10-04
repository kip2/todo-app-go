package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

/*
.envファイルから特定のキーに紐づく値を取得する関数
*/
func LoadEnv(key string) string {
	// .envに定義した環境変数をロード
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// 環境変数から値を取得
	env := os.Getenv(key)

	return env
}
