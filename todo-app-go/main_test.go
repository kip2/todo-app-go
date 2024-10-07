package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
	"todoApp/internal/db"
	"todoApp/internal/env"
	errorpkg "todoApp/internal/error"
	"todoApp/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestDeleteHandler(t *testing.T) {
	// deleteテスト用のデータを作成
	untilTime := "2024-12-31"
	untilDate, err := time.Parse("2006-01-02", untilTime)
	errorpkg.CheckError(err)

	testData := models.RegisterRequest{
		Content: "todo test content",
		Until:   untilDate,
	}

	// testデータのインサート
	err = db.InsertById(1, testData)
	if err != nil {
		t.Fatalf("Failed to insert query: %v", err)
	}

	// delete用のリクエストデータを作成
	reqBody := models.DeleteRequest{
		ID: 1,
	}

	jsonData, err := json.Marshal((reqBody))
	if err != nil {
		t.Fatalf("Failed to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", "/delete", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed to delete request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(deleteHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var resBody models.Response
	if err := json.NewDecoder(rr.Body).Decode(&resBody); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	expectedMessage := "SUCCESS"
	if resBody.Result != expectedMessage {
		t.Errorf("Handler returned unexpected body: got %v want %v", resBody.Result, expectedMessage)
	}
}

func TestRegisterHandler(t *testing.T) {
	// リクエスト用のJSONデータの作成
	untilTime := "2024-12-31"
	untilDate, err := time.Parse("2006-01-02", untilTime)
	errorpkg.CheckError(err)

	reqBody := models.RegisterRequest{
		Content: "todo test content",
		Until:   untilDate,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Failed to marshal request: %v", err)
	}

	// JSONリクエストの作成
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// レスポンス記録のためのレコーダーを用意
	rr := httptest.NewRecorder()

	// ハンドラーの呼び出し
	handler := http.HandlerFunc(registerHandler)
	handler.ServeHTTP(rr, req)

	// ステータスコードが200かの確認
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// レスポンスの内容を確認
	var resBody models.Response
	if err := json.NewDecoder(rr.Body).Decode(&resBody); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	expectedMessage := "SUCCESS"
	if resBody.Result != expectedMessage {
		t.Errorf("Handler returned unexpected body: got %v want %v", resBody.Result, expectedMessage)
	}
}

/*
テスト用DBへの接続と値の取得テスト
*/
func TestDBQuery(t *testing.T) {
	envVar := "DATABASE"
	db := db.CreateDBConnection(envVar)
	defer db.Close()

	var users []models.User
	err := db.Select(&users, "SELECT id, name FROM users WHERE id=?", 1)

	// クエリ実行時のエラーをテスト
	assert.NoError(t, err, "クエリ実行時にエラーが発生しました")

	// 期待するUserデータ
	var expectedUser = models.User{
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
	dsn := env.LoadEnv("DATABASE")
	assert.Equal(t, "test-dsn", dsn, "環境変数の値が正しく読み込まれていません")
}

func TestCreateDBConnection(t *testing.T) {
	db := db.CreateDBConnection("DATABASE")
	defer db.Close()

	assert.NotNil(t, db, "DBコネクションが作成されていません")
}
