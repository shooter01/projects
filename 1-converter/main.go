package main

import (
	"fmt"
	"strings"
)

var arrOfFiats = []string{"USD", "RUB", "EUR"}

func main() {

	fromFiat := getUserInputFrom()
	toFiat := getUserInputTo(fromFiat)
	fiatValue := getFiatValue()
	fmt.Println(converter(fiatValue, fromFiat, toFiat))
	fmt.Println(toFiat)
	fmt.Println(fiatValue)
}

func converter(count int, fiatFrom string, fiatTo string) string {
	const USDToEURO = 0.94
	const USDToRUB = 1.5
	const EUROToRUB = USDToRUB / USDToEURO
	return ""
}

// func getUserInput() (float64, float64, float64) {
// 	var count float64
// 	var fiatFrom float64
// 	var fiatTo float64
// 	fmt.Println("Введите данные")
// 	fmt.Print("Валюта откуда (исходная валюта) (USD, RUB, EUR): ")
// 	fmt.Scan(&fiatFrom)
// 	fmt.Print("Введите количество денег (только цифры): ")
// 	fmt.Scan(&count)
// 	fmt.Print("Валюта куда (целевая валюта) (USD, RUB, EUR): ")
// 	fmt.Scan(&fiatTo)
// 	return count, fiatFrom, fiatTo
// }

func getFiatsString(excludingFiat string) string {
	localArrOfFiats := remove(arrOfFiats, excludingFiat)
	return "(" + strings.Join(localArrOfFiats, ", ") + ")"
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func getUserInputFrom() string {
	var fiatFrom string
FiatFromCycle:
	for {
		fmt.Printf("Валюта откуда (исходная валюта) %s: ", getFiatsString(""))
		fmt.Scan(&fiatFrom)
		if fiatFrom == "USD" || fiatFrom == "EUR" || fiatFrom == "RUB" {
			break FiatFromCycle
		} else {
			fmt.Println("Такая валюта не обнаружена, попробуйте снова")
		}
	}

	return fiatFrom
}

func getUserInputTo(excludingFiat string) string {
	var fiatTo string
FiatFromCycle:
	for {
		fmt.Printf("Валюта куда (целевая валюта) %s:", getFiatsString(excludingFiat))
		fmt.Scan(&fiatTo)
		if excludingFiat == fiatTo {
			fmt.Println("Данная валюта указана в качестве исходной, попробуйте снова")
			continue
		}
		if fiatTo == "USD" || fiatTo == "EUR" || fiatTo == "RUB" {
			break FiatFromCycle
		} else {
			fmt.Print("Такая валюта не обнаружена, попробуйте снова")
		}
	}

	return fiatTo
}

func getFiatValue() float64 {
	var fiatValue float64
	fmt.Print("Количество валюты:")
	fmt.Scan(&fiatValue)

	return fiatValue
}
