package handler

import (
	"os"

	"github.com/golang-jwt/jwt"
)

func GenerateJwtToken(email string) (string, error) {
	mySigningKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	type MyCustomClaims struct {
		jwt.StandardClaims
	}

	claims := MyCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: 300,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return ss, nil
}

func RefreshToeken(tokenString string) (string, error) {
	return "", nil
}
