package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
