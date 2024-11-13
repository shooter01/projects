package main

import (
	"fmt"
	"strings"
)

type currencyMap = map[string]map[string]float64

var conversionRates = currencyMap{
	"USD": {"EUR": 0.94, "RUB": 95},
	"EUR": {"USD": 1 / 0.94, "RUB": 95 / 0.94},
	"RUB": {"USD": 1.0 / 95, "EUR": 0.94 / 95},
}

var arrOfFiats = []string{"USD", "RUB", "EUR"}

func main() {

	fromFiat := getUserInput("Валюта откуда (исходная валюта)", "")
	toFiat := getUserInput("Валюта куда (целевая валюта)", fromFiat)
	fiatValue := getFiatValue()
	fmt.Printf("Количество валюты, которую вы получите: %f %s\n", converter(fiatValue, fromFiat, toFiat, conversionRates), toFiat)
}

func converter(count float64, fiatFrom string, fiatTo string, conversionRates currencyMap) float64 {
	if rate, exists := conversionRates[fiatFrom][fiatTo]; exists {
		return count * rate
	}
	return 0
}

func getFiatsString(excludingFiat string) string {
	filtered := remove(arrOfFiats, excludingFiat)
	return "(" + strings.Join(filtered, ", ") + ")"
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func getUserInput(prompt string, exclude string) string {
	var currency string
FiatCycle:
	for {
		fmt.Printf("%s %s: ", prompt, getFiatsString(exclude))
		fmt.Scan(&currency)
		currency = strings.ToUpper(currency)
		if currency == exclude {
			fmt.Println("Данная валюта указана в качестве исходной, попробуйте снова")
		} else if isValidCurrency(currency) {
			break FiatCycle
		} else {
			fmt.Println("Такая валюта не обнаружена, попробуйте снова")
		}
	}
	return currency
}

func isValidCurrency(currency string) bool {
	for _, c := range arrOfFiats {
		if c == currency {
			return true
		}
	}
	return false
}

func getFiatValue() float64 {
	var fiatValue float64
	fmt.Print("Количество валюты (исходной): ")
	fmt.Scan(&fiatValue)

	return fiatValue
}
