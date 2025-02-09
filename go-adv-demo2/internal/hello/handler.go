package hello

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//func hello(w http.ResponseWriter, req *http.Request) {
//	fmt.Println("Hello World")
//}

type HelloHandler struct{}

func NewHelloHandler(router *http.ServeMux) {
	handler := &HelloHandler{}
	router.HandleFunc("/hello", handler.Hello())
	router.HandleFunc("/random", handler.Random())
}

func (h HelloHandler) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Hello World")
	}
}

func (h HelloHandler) Random() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
		num := rand.Intn(6) + 1          // Генерация числа от 1 до 6

		w.Write([]byte(strconv.Itoa(num))) // Отправка числа как строки
	}
}
