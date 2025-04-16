package auth

import (
	"fmt"
	"net/http"

	"go/adv-demo/configs"
	"go/adv-demo/pkg/req"
	responses "go/adv-demo/pkg/res"
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
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)
		fmt.Println(err.Error())

		if err != nil {
			return
		}
		fmt.Println(body)
		res := LoginResponse{
			Token: "123",
		}
		responses.Json(w, res, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)
		fmt.Println(err.Error())

		if err != nil {
			return
		}
		fmt.Println(body)

		res := RegisterResponse{
			Token: "333",
		}
		responses.Json(w, res, 200)
	}
}
