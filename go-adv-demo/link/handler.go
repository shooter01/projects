package link

import (
	"fmt"
	"net/http"

	"go/adv-demo/configs"
	"go/adv-demo/pkg/req"
	responses "go/adv-demo/pkg/res"
)

type LinkHandlerDeps struct {
	*configs.Config
}

type LinkHandler struct {
	*configs.Config
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /link/login", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{alias}", handler.GoTo())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkRequest](w, r)
		fmt.Println(err.Error())

		if err != nil {
			return
		}
		fmt.Println(body)
		res := LinkResponse{
			Token: "123",
		}
		responses.Json(w, res, 200)
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkRequest](w, r)
		fmt.Println(err.Error())

		if err != nil {
			return
		}
		fmt.Println(body)
		res := LinkResponse{
			Token: "123",
		}
		responses.Json(w, res, 200)
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkRequest](w, r)
		fmt.Println(err.Error())

		if err != nil {
			return
		}
		fmt.Println(body)
		res := LinkResponse{
			Token: "123",
		}
		responses.Json(w, res, 200)
	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkRequest](w, r)
		fmt.Println(err.Error())

		if err != nil {
			return
		}
		fmt.Println(body)
		res := LinkResponse{
			Token: "123",
		}
		responses.Json(w, res, 200)
	}
}
