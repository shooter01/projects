package main

import (
	"fmt"
)

func main() {
	var rubCount = 100.0;
	const USD = 1.4;
	const EUR = 1.5;
	fmt.Print(rubCount * USD)
}

func coverter(count int, fiatFrom string, fiatTo string){
	
}

func getUserInput() (float64, float64, float64) {
	var count float64;
	var fiatFrom float64;
	var fiatTo float64;
	fmt.Println("Введите данные");
	fmt.Print("Введите количество денег: ");
	fmt.Scan(&count);
	fmt.Print("Валюта откуда: ");
	fmt.Scan(&fiatFrom)
	fmt.Print("Валюта куда: ");
	fmt.Scan(&fiatTo);
	return count, fiatFrom, fiatTo
}