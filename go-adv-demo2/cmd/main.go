package main

import (
	"fmt"
	"go/adv-demo2/configs"
	"go/adv-demo2/db"

	"go/adv-demo2/internal/auth"
	"go/adv-demo2/internal/link"
	"go/adv-demo2/internal/verify"
	"net/http"

	"github.com/gorilla/mux"
)

//func verify(w http.ResponseWriter, req *http.Request) {
//	fmt.Println("Hello World")
//}

func main() {
	conf := configs.LoadConfig()
	dbName := db.NewDb(conf)
	router := mux.NewRouter()
	verify.NewHelloHandler(router)

	// Repositories
	linkRepository := link.NewLinkRepository(dbName)

	//handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("SERVER IS LISTENING")
	server.ListenAndServe()
}
