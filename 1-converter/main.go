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
}

func converter(count float64, fiatFrom string, fiatTo string) float64 {
	const USDToEURO = 0.94
	const USDToRUB = 95
	const EUROToRUB = USDToRUB / USDToEURO
	switch {
	case fiatFrom == "USD" && fiatTo == "RUB":
		return count * USDToRUB
	case fiatFrom == "RUB" && fiatTo == "USD":
		return count / USDToRUB
	case fiatFrom == "EUR" && fiatTo == "USD":
		return count / USDToEURO
	case fiatFrom == "USD" && fiatTo == "EUR":
		return count * USDToEURO
	case fiatFrom == "EUR" && fiatTo == "RUB":
		return count * EUROToRUB
	case fiatFrom == "RUB" && fiatTo == "EUR":
		return count * EUROToRUB
	default:
		return 0
	}
}

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
