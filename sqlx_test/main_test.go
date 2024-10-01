package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
テスト用DBへの接続と値の取得テスト
*/
func TestDBQuery(t *testing.T) {
	envVar := "DATABASE"
	db := createDBConnection(envVar)
	defer db.Close()

	var users []User
	err := db.Select(&users, "SELECT id, name FROM users WHERE id=?", 1)

	// クエリ実行時のエラーをテスト
	assert.NoError(t, err, "クエリ実行時にエラーが発生しました")

	// 期待するUserデータ
	var expectedUser = User{
		ID:   1,
		Name: "Alice",
	}

	// テスト用ユーザーデータの取得をアサート
	assert.Equal(t, 1, len(users), "ユーザーが取得できませんでした")
	assert.Equal(t, expectedUser.ID, users[0].ID, "ユーザーIDが一致しません")
	assert.Equal(t, expectedUser.Name, users[0].Name, "ユーザー名が一致しません")
}

func TestLoadEnv(t *testing.T) {
	os.Setenv("DATABASE", "test-dsn")
	dsn := loadEnv("DATABASE")
	assert.Equal(t, "test-dsn", dsn, "環境変数の値が正しく読み込まれていません")
}

func TestCreateDBConnection(t *testing.T) {
	db := createDBConnection("DATABASE")
	defer db.Close()

	assert.NotNil(t, db, "DBコネクションが作成されていません")
}
