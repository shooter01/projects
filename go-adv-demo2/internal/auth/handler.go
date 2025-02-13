package auth

import (
	"fmt"
	"go/adv-demo2/configs"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//func hello(w http.ResponseWriter, req *http.Request) {
//	fmt.Println("Hello World")
//}

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct{}

func NewAuthHandler(router *http.ServeMux, deps *AuthHandlerDeps) {
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
