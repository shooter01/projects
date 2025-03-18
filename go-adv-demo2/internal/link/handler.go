package link

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	request "go/adv-demo2/pkg/req"
	response "go/adv-demo2/pkg/res"
)

//func verify(w http.ResponseWriter, req *http.Request) {
//	fmt.Println("Hello World")
//}

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

type LinkHandler struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *mux.Router, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("/link", handler.Create()).Methods("POST")
	router.HandleFunc("/link", handler.Patch()).Methods("PATCH")
	router.HandleFunc("/link/{id}", handler.Delete()).Methods("DELETE")
	router.HandleFunc("/{alias}", handler.Get()).Methods("GET")
}

func (h *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[LinkCreateRequest](w, req)
		if err != nil {
			return
		}
		link := NewLink(body.Url)
		createdLink, err := h.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		response.Json(w, createdLink, http.StatusCreated)

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
