package main

import (
	"fmt"
	"math"
)

func main() {
	userHeight, userKg := getUserInput();
	var IMT = calculateIMT(userHeight, userKg)
	outputResult(IMT)
}

func outputResult(IMT float64){
	result := fmt.Sprintf("Ваш индекс %.2f", IMT)
	fmt.Print(result)
}

func calculateIMT(userHeight, userKg float64) float64 {
	return (userKg / math.Pow(userHeight / 100, 2))
}

func getUserInput() (float64, float64) {
	var userHeight float64;
	var userKg float64;
	fmt.Println("Калькулятор массы тела");
	fmt.Print("Введите рост (в сантиметрах): ");
	fmt.Scan(&userHeight);
	fmt.Print("Введите вес: ");
	fmt.Scan(&userKg);
	return userHeight, userKg
}