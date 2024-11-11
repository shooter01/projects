package main

import (
	"fmt"
)

func main() {
	const USDToEURO = 0.94
	const USDToRUB = 1.5
	const EUROToRUB = USDToRUB / USDToEURO
	fmt.Print(EUROToRUB)
}

func coverter(count int, fiatFrom string, fiatTo string) {

}

func getUserInput() (float64, float64, float64) {
	var count float64
	var fiatFrom float64
	var fiatTo float64
	fmt.Println("Введите данные")
	fmt.Print("Введите количество денег: ")
	fmt.Scan(&count)
	fmt.Print("Валюта откуда: ")
	fmt.Scan(&fiatFrom)
	fmt.Print("Валюта куда: ")
	fmt.Scan(&fiatTo)
	return count, fiatFrom, fiatTo
}
