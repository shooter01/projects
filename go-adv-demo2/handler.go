package main

import (
	"fmt"
	"net/http"
)

//func hello(w http.ResponseWriter, req *http.Request) {
//	fmt.Println("Hello World")
//}

type HelloHandler struct{}

func NewHelloHandler(router *http.ServeMux) {
	handler := &HelloHandler{}
	router.HandleFunc("/hello", handler.Hello())
}

func (h HelloHandler) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Hello World")
	}
}
