package main

import (
	pkg "demo/app-2/pkg"
	"fmt"
)

type bookmarksMap = map[string]string

func main() {
	transactions := []float64{1, 2, 3, 4}

	calc := pkg.Calculator{
		Data: "test1",
	}
	data := calc.NewCalc()
	fmt.Println(data)
	//maxIter := 3
	//for {
	//	number := getTrabsaction()
	//	transactions = append(transactions, number)
	//	maxIter--
	//	if maxIter < 0 {
	//		break
	//	}
	//}

	for {
		number := enterBookmark()
		transactions = append(transactions, number)
		maxIter--
		if maxIter < 0 {
			break
		}
	}

	createMap()
	fmt.Println(transactions)

	//fmt.Println(calc(transactions))
	//fmt.Println(transactions)
	//transactions[3] = "2"
	//transactionsNew := transactions
	//transactions = append(transactions, "100")
	//banks := [3]string{"test", "test2"}
	//fmt.Println(len(transactions[1:5]))
	//fmt.Println(banks)
	fmt.Println(transactions)
	//fmt.Println(transactionsNew)
}

func getTrabsaction() float64 {
	var number float64
	fmt.Print("Enter number: ")
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
