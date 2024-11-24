package middleware

import (
	"fmt"
	"net/http"
)

func Loggin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Loggin")
		next.ServeHTTP(w, r)
		fmt.Println("After")

	})
}
