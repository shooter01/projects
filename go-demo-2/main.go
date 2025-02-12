package main

import "fmt"

func main() {
	transactions := []float64{}
	maxIter := 3
	for {
		number := getTrabsaction()
		transactions = append(transactions, number)
		maxIter--
		if maxIter < 0 {
			break
		}
	}
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
