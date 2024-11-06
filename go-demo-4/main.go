package main

import (
	"demo/password/account"
	"fmt"

	"github.com/fatih/color"
) 



func main() {
	fmt.Println("____Менеджер паролей")
	vault := account.NewVault()

Menu: 
	for {
		variant := getMenu()
		switch variant{
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}
	
	
	
}



func deleteAccount(vault *account.Vault){
	url := promptData("Введите URL")
	vault.DeleteAccountByURL(url)
}

func findAccount(vault *account.Vault){
	url := promptData("Введите URL")
	accounts := vault.FindAccountsByURL(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}


func createAccount(vault *account.Vault){
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите url")
	myAccount, err := account.NewAccount(login, password, url)
	// files.ReadFile()
	if (err != nil) {
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


func getMenu() int{
	var variant int
	fmt.Println("Выберите враиант:")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")

	fmt.Scanln(&variant)

	return variant
}


func double(number *int) {
	*number = *number * 2
}


func promptData(str string) (string) {
	var prompt string;
	fmt.Println(str);
	fmt.Scanln(&prompt);
	return prompt
}


// func reverse(numbers *[]int) {
// 	for index, value := range *numbers {
// 		numbers[len(*numbers) - 1 - index] = value
// 	}
// }