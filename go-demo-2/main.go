package main

import (
	//pkg "demo/app-2/pkg"
	"fmt"
)

type bookmarksMap = map[string]string

type account struct {
	login    string
	password string
	url      string
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(reverse(&arr))

	login := promptData("Введите логин")
	password := promptData("Введите password")
	url := promptData("Введите url")

	account1 := account{
		login:    login,
		password: password,
		url:      url,
	}

	printPrompt(&account1)

	//transactions := []float64{1, 2, 3, 4}
	//
	//calc := pkg.Calculator{
	//	Data: "test1",
	//}
	//data := calc.NewCalc()
	//fmt.Println(data)
	//maxIter := 3
	//for {
	//	number := getTrabsaction()
	//	transactions = append(transactions, number)
	//	maxIter--
	//	if maxIter < 0 {
	//		break
	//	}
	//}

	//for {
	//	number := enterBookmark()
	//	transactions = append(transactions, number)
	//	maxIter--
	//	if maxIter < 0 {
	//		break
	//	}
	//}
	//
	//createMap()
	//fmt.Println(transactions)

	//fmt.Println(calc(transactions))
	//fmt.Println(transactions)
	//transactions[3] = "2"
	//transactionsNew := transactions
	//transactions = append(transactions, "100")
	//banks := [3]string{"test", "test2"}
	//fmt.Println(len(transactions[1:5]))
	//fmt.Println(banks)
	//fmt.Println(transactions)
	//fmt.Println(transactionsNew)
}

func printPrompt(data *account) {
	fmt.Println(data)
}

func promptData(prompt string) string {
	var number string
	fmt.Print(prompt + ": ")
	fmt.Scan(&number)
	return number
}

func calc(tr []float64) float64 {
	var summa float64
	for _, key := range tr {
		summa = summa + float64(key)
	}
	tr[0] = 2222
	return summa
}

func createMap() bookmarksMap {
	myMap := map[string]string{"a": "1", "b": "2", "c": "3"}
	for key, value := range myMap {
		fmt.Println(key, value)
	}
	return myMap
}

func reverse(arr *[]int) *[]int {
	n := len(*arr)
	result := make([]int, n)
	for i, v := range *arr {
		result[n-1-i] = v
	}
	return &result
}

func showBookmarks(myMap bookmarksMap) {
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}

func enterBookmark(myMap *bookmarksMap) *bookmarksMap {
	var key, value string
	fmt.Print("Enter bookmark key: ")
	fmt.Scan(&key)
	fmt.Print("Enter bookmark value: ")
	fmt.Scan(&value)
	(*myMap)[key] = value
	return myMap
}

func deleteBookmark(myMap *bookmarksMap, key string) {
	delete(*myMap, key)
}
