package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gbrlsnchs/jwt"
)

// Cors adiciona os headers para suportar o CORS nos navegadores
func Auth(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	cors(w, r)
	if r.Host == "localhost:8080/api/v1/login" {
		next(w, r)
	}

	// token := w.Header().Get("Authorization")
	authorizationHeader := r.Header.Get("Authorization")

	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) == 2 {
			checked := verifyToken(bearerToken, w, r)
			if checked == false {
				return
			}
		}
	}

	next(w, r)
}

func verifyToken(tokenString []string, w http.ResponseWriter, r *http.Request) bool {

	mySigningKey := os.Getenv("JWT_SECRET_KEY")
	jot, err := jwt.FromString(tokenString[1])

	if err != nil {
		// Handle malformed token...
		fmt.Sprintln("The token is invalid.")
		return false
	}

	if err = jot.Verify(jwt.HS256(mySigningKey)); err != nil {
		// Handle verification error...
		fmt.Sprintln("The token is invalid.")
		return false
	}

	fmt.Sprintln("The token is valid.")
	fmt.Sprintln(tokenString[1])
	auth := fmt.Sprintf("Bearer %s", tokenString[1])

	w.Header().Set("Authorization", auth)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tokenString[1]))
	return true
}

func cors(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		return
	}
}
