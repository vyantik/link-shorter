package jwt_test

import (
	"app/test/pkg/jwt"
	"testing"
)

func TestJWTSign(t *testing.T) {
	jwtInstance := jwt.NewJWT("gsxL6L1mnFJXi92GOC8HokhCvBHmY86hWibgGPtVHOY=")
	email := "a@a.ru"

	token, err := jwtInstance.Generate(jwt.JWTData{
		Email: email,
	})

	if err != nil {
		t.Fatalf("[JWT] - [TestJWTSign] - [ERROR] %s", err)
	}

	isValid, data := jwtInstance.Parse(token)

	if !isValid {
		t.Fatal("Token is invalid")
	}

	if data.Email != email {
		t.Fatalf("Email %s not equal %s", data.Email, email)
	}
}
