package main

import (
	"app/test/internal/auth"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginSuccess(t *testing.T) {
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
}

func TestLoginInvalidPassword(t *testing.T) {
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
}
