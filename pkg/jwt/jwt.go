package jwt

import "github.com/golang-jwt/jwt/v5"

type JWT struct {
	secretKey string
}

func NewJWT(secretKey string) *JWT {
	return &JWT{
		secretKey: secretKey,
	}
}

func (j *JWT) GenerateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})

	signedToken, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
