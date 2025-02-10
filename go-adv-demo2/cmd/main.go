package main

import (
	"fmt"
	"go/adv-demo2/configs"
	"go/adv-demo2/internal/hello"
	"net/http"
)

//func hello(w http.ResponseWriter, req *http.Request) {
//	fmt.Println("Hello World")
//}

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	hello.NewHelloHandler(router)
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("SERVER IS LISTENING")
	server.ListenAndServe()
}
