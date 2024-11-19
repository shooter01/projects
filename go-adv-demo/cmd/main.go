package main

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/internal/auth"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()

	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
	// code := make(chan int)
	// arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	// arr1 := arr[0 : len(arr)/3]
	// arr2 := arr[len(arr)/3 : len(arr)-len(arr)/3]
	// arr3 := arr[len(arr)-len(arr)/3:]

	// // fmt.Println("Привет!")
	// // t := time.Now()
	// var wg sync.WaitGroup
	// wg.Add(3)

	// go func() {
	// 	sumArray(arr1, code)
	// 	wg.Done()
	// }()
	// go func() {
	// 	sumArray(arr2, code)
	// 	wg.Done()
	// }()
	// go func() {
	// 	sumArray(arr3, code)
	// 	wg.Done()
	// }()

	// go func() {
	// 	wg.Wait()
	// 	close(code)
	// }()
	// sum := 0

	// for res := range code {
	// 	sum = sum + res
	// }
	// fmt.Println(sum)

	// // println(time.Since(t).Seconds())
	// // time.Sleep(time.Second)
}

func makeRequest() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("status code %d \n", resp.StatusCode)
	// defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(string(body))
}

func sumArray(numbers []int, code chan int) {
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	code <- result
	// return result
}
