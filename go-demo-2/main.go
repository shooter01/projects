package main

import "fmt"

func main() {
	transactions := []string{"1", "2"}
	//transactions[3] = "2"
	//transactionsNew := transactions
	transactions = append(transactions, "100")
	//banks := [3]string{"test", "test2"}
	//fmt.Println(len(transactions[1:5]))
	//fmt.Println(banks)
	fmt.Println(transactions)
	//fmt.Println(transactionsNew)
}
