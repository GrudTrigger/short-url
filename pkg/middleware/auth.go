package middleware

import (
	"fmt"
	"go/backend/configs"
	"go/backend/pkg/jwt"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authedHeader, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		fmt.Println(isValid)
		fmt.Println(data)
		fmt.Println(token)
		next.ServeHTTP(w, r)
	})
}
