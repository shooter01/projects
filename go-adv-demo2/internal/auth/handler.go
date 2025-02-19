package auth

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go/adv-demo2/configs"
	"go/adv-demo2/pkg"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//func verify(w http.ResponseWriter, req *http.Request) {
//	fmt.Println("Hello World")
//}

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *mux.Router, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("/auth/login", handler.Login()).Methods("POST")
	router.HandleFunc("/auth/register", handler.Register()).Methods("POST")
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(h.Config.Auth.Secret)
		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)
		fmt.Println(payload)
		if err != nil {
			pkg.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		if payload.Email == "" {
			pkg.Json(w, "Email required", http.StatusBadRequest)
			return
		}
		if payload.Password == "" {
			pkg.Json(w, "Password required", http.StatusBadRequest)
			return
		}
		res := LoginResponse{
			Token: "123",
		}
		pkg.Json(w, res, 201)
	}
}

func (h AuthHandler) Register() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
		num := rand.Intn(6) + 1          // Генерация числа от 1 до 6

		w.Write([]byte(strconv.Itoa(num))) // Отправка числа как строки
	}
}
