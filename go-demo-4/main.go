package main

import (
	"demo/password/account"
	"demo/password/encrypter"
	"demo/password/files"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByLogin,
	"3": findAccountByUrl,
	"4": deleteAccount,
}

func main() {
	fmt.Println("____Менеджер паролей")

	// err := godotenv.Load()
	// if err != nil {
	// 	output.PrintError("Не удалось найти env файл")
	// }

	for _, e := range os.Environ() {
		fmt.Println(strings.SplitN(e, "=", 2))
		fmt.Println(e)
	}
	vault := account.NewVault(files.NewJsonDb("data.vault"), *encrypter.NewEncrypter())

Menu:
	for {
		// variant := getMenu()
		variant := promptData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт по логину",
			"3. Найти аккаунт по url",
			"4. Удалить аккаунт",
			"5. Выход",
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

// func findAccount(vault *account.VaultWithDb) {
// 	url := promptData([]string{"Введите URL"})
// 	accounts := vault.FindAccounts(url, checkUrl)
// 	if len(accounts) == 0 {
// 		color.Red("Аккаунтов не найдено")
// 	}
// 	for _, account := range accounts {
// 		account.Output()
// 	}
// }

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите login"})
	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outputResult(&accounts)
}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите url"})
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outputResult(&accounts)
}

func outputResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range *accounts {
		account.Output()
	}
}

func checkUrl(acc account.Account, str string) bool {
	return strings.Contains(acc.Url, str)
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
