package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go/adv-demo2/configs"
	"go/adv-demo2/internal/auth"
	"go/adv-demo2/internal/verify"
	"net/http"
)

//func verify(w http.ResponseWriter, req *http.Request) {
//	fmt.Println("Hello World")
//}

func main() {
	conf := configs.LoadConfig()
	router := mux.NewRouter()
	verify.NewHelloHandler(router)
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("SERVER IS LISTENING")
	server.ListenAndServe()
}
