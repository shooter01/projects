package link

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

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
	router.HandleFunc("/link/{id}", handler.Update()).Methods("PUT")
	router.HandleFunc("/link", handler.Patch()).Methods("PATCH")
	router.HandleFunc("/link/{id}", handler.Delete()).Methods("DELETE")
	router.HandleFunc("/{alias}", handler.GoTo()).Methods("GET")
}

func (h *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[LinkCreateRequest](w, req)
		if err != nil {
			return
		}
		link := NewLink(body.Url)
		for {
			existedLink, _ := h.LinkRepository.GetByHash(link.Hash)
			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}

		createdLink, err := h.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		response.Json(w, createdLink, http.StatusCreated)

	}
}

func (h *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[LinkUpdateRequest](w, req)
		if err != nil {
			return
		}
		idString := mux.Vars(req)["id"]
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		link, err := h.LinkRepository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		response.Json(w, link, http.StatusCreated)

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

func (h LinkHandler) GoTo() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		hash := mux.Vars(req)["alias"]
		link, err := h.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, req, link.Url, http.StatusTemporaryRedirect)
	}
}
