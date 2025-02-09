package main

import "fmt"

func main() {
	transactions := [5]string{"1", "2"}
	banks := [3]string{"test", "test2"}
	fmt.Println(len(transactions[1:5]))
	fmt.Println(banks)
}
