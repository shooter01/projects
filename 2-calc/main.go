package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	operation, userNumbers := getUserInput()
	numbers := formatUserNumbers(userNumbers)
	value := getResult(operation, numbers)

	fmt.Println(operation)
	fmt.Println(value)
}

func getUserInput() (string, string) {
	var operation string
	var userNumbers string
	fmt.Println("Введите операцию")
	fmt.Print("Введите операцию (AVG, SUM, MED): ")
	fmt.Scan(&operation)
	fmt.Println(operation)

	operation = strings.Trim(operation, " ")

	if operation != "AVG" && operation != "MED" && operation != "SUM" {
		panic("Неправильный оператор")
	}

	fmt.Println("Введите числа через запятую: ")
	fmt.Scan(&userNumbers)
	return operation, userNumbers
}

func formatUserNumbers(stringToFromat string) []float64 {
	numbers := strings.Split(stringToFromat, ",")
	floatNumbers := []float64{}

	for _, value := range numbers {
		number, _ := strconv.ParseFloat(value, 64)
		floatNumbers = append(floatNumbers, number)
	}
	return floatNumbers
}

func getResult(operationName string, numbers []float64) float64 {
	var total float64 = 0
	for _, value := range numbers {
		total += value
	}
	switch {
	case operationName == "AVG":
		return total / float64(len(numbers))
	case operationName == "SUM":
		return total
	case operationName == "MED":
		return calculateMedian(numbers)
	}
	return total
}

func calculateMedian(data []float64) float64 {
	dataCopy := make([]float64, len(data))
	copy(dataCopy, data)

	sort.Float64s(dataCopy)

	var median float64
	l := len(dataCopy)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		median = (dataCopy[l/2-1] + dataCopy[l/2]) / 2
	} else {
		median = dataCopy[l/2]
	}

	return median
}
