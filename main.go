package main

import (
	"fmt"
	"math"
)

func main() {
	maxIterations := 3
	transactions := []float64{}
	for {
		transaction := getTransactionsPrompt();
		transactions = append(transactions, transaction)
		
		maxIterations--

		if (maxIterations < 0) {
			break
		}

		// answer := getUserAnswer();
		// if (answer == "Yes") {
		// 	userHeight, userKg := getUserInput();
		// 	var IMT = calculateIMT(userHeight, userKg)
		// 	outputResult(IMT)
		// } else {
		// 	break
		// }
	}

	fmt.Println(transactions)
	fmt.Println("Сумма: ")
	var sum float64 = 0
	for _, v := range transactions {
		sum+=v
	}
	fmt.Println(sum)

}

func outputResult(IMT float64){
	result := fmt.Sprintf("Ваш индекс %.2f", IMT)
	fmt.Println(result)
	switch  {
	case IMT < 16:
		fmt.Println("Недовес")
	default: 
		fmt.Println("Все круто")
	}
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


func getTransactionsPrompt() (float64) {
	var number float64;
	fmt.Print("Введите число: ");
	fmt.Scan(&number);
	return number
}


func get() (float64) {
	var number float64;
	fmt.Print("Введите число: ");
	fmt.Scan(&number);
	return number
}

func getUserAnswer() (string) {
	var answer string;
	fmt.Println("Рассчитать индекс тела?");
	fmt.Scan(&answer);
	return answer
}