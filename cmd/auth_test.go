package main

import (
	"app/test/internal/auth"
	"app/test/internal/user"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("[DB] - [NewDb] - [ERROR] %s", err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("[DB] - [NewDb] - [ERROR] %s", err)
	}

	return db
}

func initData(db *gorm.DB) {
	db.Create(&user.User{
		Email:    "a@a.ru",
		Password: "$2a$10$HnXvm42medGoKidGm8016.aHR2Gy07jvA9SBceTanT/5OD84uoanS",
		Username: "vyantik",
	})
}

func removeData(db *gorm.DB) {
	db.Unscoped().Where("email = ?", "a@a.ru").Delete(&user.User{})
}

func TestLoginSuccess(t *testing.T) {
	//prepare
	//===============================================
	db := initDb()
	initData(db)
	defer removeData(db)
	//===============================================

	//test
	//===============================================
	app, _ := App()
	ts := httptest.NewServer(app)
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a@a.ru",
		Password: "12345678",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	var response auth.LoginResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if response.Token == "" {
		t.Fatal("expected token, got empty")
	}
	//===============================================
}

func TestLoginInvalidPassword(t *testing.T) {
	//prepare
	//===============================================
	db := initDb()
	initData(db)
	defer removeData(db)
	//===============================================

	//test
	//===============================================
	app, _ := App()
	ts := httptest.NewServer(app)
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a@a.ru",
		Password: "123456789",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected status code %d, got %d", http.StatusUnauthorized, res.StatusCode)
	}
	//===============================================
}
