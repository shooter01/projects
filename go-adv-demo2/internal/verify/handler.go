package verify

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//func verify(w http.ResponseWriter, req *http.Request) {
//	fmt.Println("Hello World")
//}

type VerifyHandler struct{}

func NewHelloHandler(router *mux.Router) {
	handler := &VerifyHandler{}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("/verify/{hash}", handler.Hash())
}

func (h *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Hello World")
	}
}

func (h *VerifyHandler) Hash() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req) // Извлекаем параметр {hash}
		hash, exists := vars["hash"]

		if !exists || hash == "" {
			http.Error(w, "Hash parameter not found", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Hash1: %s", hash)
	}
}
