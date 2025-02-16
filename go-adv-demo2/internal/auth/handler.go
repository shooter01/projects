package auth

import (
	"encoding/json"
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

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(h.Config.Auth.Secret)

		res := LoginResponse{
			Token: "123",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}

func (h AuthHandler) Register() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
		num := rand.Intn(6) + 1          // Генерация числа от 1 до 6

		w.Write([]byte(strconv.Itoa(num))) // Отправка числа как строки
	}
}
