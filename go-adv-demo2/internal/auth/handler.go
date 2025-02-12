package auth

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//func hello(w http.ResponseWriter, req *http.Request) {
//	fmt.Println("Hello World")
//}

type AuthHandler struct{}

func NewAuthHandler(router *http.ServeMux) {
	handler := &AuthHandler{}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (h AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Login")
	}
}

func (h AuthHandler) Register() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
		num := rand.Intn(6) + 1          // Генерация числа от 1 до 6

		w.Write([]byte(strconv.Itoa(num))) // Отправка числа как строки
	}
}
