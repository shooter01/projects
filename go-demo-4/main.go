package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccount,
	"3": deleteAccount,
}

func main() {
	fmt.Println("____Менеджер паролей")
	vault := account.NewVault(files.NewJsonDb("data.json"))

Menu:
	for {
		// variant := getMenu()
		variant := promptData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите вариант",
		})

		menuFunction := menu[variant]

		if menuFunction == nil {
			break Menu
		}

		menuFunction(vault)

		// switch variant {
		// case "1":
		// 	createAccount(vault)
		// case "2":
		// 	findAccount(vault)
		// case "3":
		// 	deleteAccount(vault)
		// default:
		// 	break Menu
		// }
	}

}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL"})
	vault.DeleteAccountByURL(url)
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL"})
	accounts := vault.FindAccountsByURL(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	url := promptData([]string{"Введите url"})
	myAccount, err := account.NewAccount(login, password, url)
	// files.ReadFile()
	if err != nil {
		fmt.Println("'Неверный формат логин или урл'")
		return
	}
	vault.AddAccount(*myAccount)
	// data, err := vault.ToByte()
	// if err != nil {
	// 	fmt.Println("Не удалось преобразовать данные")
	// }
	// files.WriteFile(string(data), "data.json")
	// color.RGB(255, 128, 0).Println(myAccount)
}

// func getMenu() int {
// 	var variant int
// 	fmt.Println("Выберите враиант:")
// 	fmt.Println("1. Создать аккаунт")
// 	fmt.Println("2. Найти аккаунт")
// 	fmt.Println("3. Удалить аккаунт")
// 	fmt.Println("4. Выход")

// 	fmt.Scanln(&variant)

// 	return variant
// }

func double(number *int) {
	*number = *number * 2
}

func promptData[T any](prompt []T) string {
	for i, line := range prompt {
		if len(prompt)-1 == i {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var str string
	// fmt.Println(str)
	fmt.Scanln(&str)
	return str
}

// func reverse(numbers *[]int) {
// 	for index, value := range *numbers {
// 		numbers[len(*numbers) - 1 - index] = value
// 	}
// }
