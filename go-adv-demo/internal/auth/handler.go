package auth

import (
	"fmt"
	"net/http"

	"go/adv-demo/configs"
	responses "go/adv-demo/pkg"
)

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

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(handler.Config.Auth.Secret)
		fmt.Println("Login")
		res := LoginResponse{
			Token: "123",
		}

		responses.Json(w, res, 200)

	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Login")
	}
}
