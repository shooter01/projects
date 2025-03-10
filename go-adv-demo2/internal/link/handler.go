package link

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"go/adv-demo2/configs"
)

//func verify(w http.ResponseWriter, req *http.Request) {
//	fmt.Println("Hello World")
//}

type LinkHandlerDeps struct {
	*configs.Config
}

type LinkHandler struct {
	*configs.Config
}

func NewLinkHandler(router *mux.Router, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		Config: deps.Config,
	}
	router.HandleFunc("/link", handler.Create()).Methods("POST")
	router.HandleFunc("/link", handler.Patch()).Methods("PATCH")
	router.HandleFunc("/link/{id}", handler.Delete()).Methods("DELETE")
	router.HandleFunc("/{alias}", handler.Get()).Methods("GET")
}

func (h *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}

func (h LinkHandler) Patch() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

	}
}

func (h LinkHandler) Delete() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		id := mux.Vars(req)["id"]
		fmt.Println(id)
	}
}

func (h LinkHandler) Get() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

	}
}
