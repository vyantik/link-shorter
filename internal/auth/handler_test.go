package auth_test

import (
	"app/test/configs"
	"app/test/internal/auth"
	"app/test/internal/user"
	"app/test/pkg/db"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func bootstrap() (*auth.AuthHandler, sqlmock.Sqlmock, error) {
	connection, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	database, err := gorm.Open(postgres.New(postgres.Config{Conn: connection}))
	if err != nil {
		return nil, nil, err
	}
	userRepository := user.NewUserRepository(&db.Db{
		DB: database,
	})
	handler := auth.AuthHandler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		AuthService: auth.NewAuthService(userRepository),
	}
	return &handler, mock, nil
}

func TestLoginSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("a@a.ru", "$2a$10$HnXvm42medGoKidGm8016.aHR2Gy07jvA9SBceTanT/5OD84uoanS")
	mock.ExpectQuery("SELECT").
		WillReturnRows(rows)
	if err != nil {
		t.Fatal(err)
	}
	defer mock.ExpectClose()
	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a@a.ru",
		Password: "12345678",
	})
	reader := bytes.NewBuffer(data)
	writer := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/auth/login", reader)
	handler.Login()(writer, req)
	if writer.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, writer.Code)
	}
}

func TestRegisterSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	rows := sqlmock.NewRows([]string{"email", "password", "username"})
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.ExpectClose()
	data, _ := json.Marshal(&auth.RegisterRequest{
		Email:    "a@a.ru",
		Username: "vyantik",
		Password: "12345678",
	})
	reader := bytes.NewBuffer(data)
	writer := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/auth/register", reader)
	handler.Register()(writer, req)
	if writer.Code != http.StatusCreated {
		t.Fatalf("Expected status code %d, got %d", http.StatusCreated, writer.Code)
	}
}
