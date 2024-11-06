package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
) 



func main() {
	fmt.Println("____Менеджер паролей")
Menu: 
	for {
		variant := getMenu()
		switch variant{
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
	
	
	
}

func deleteAccount(){

}

func findAccount(){
	
}


func createAccount(){
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите url")
	myAccount, err := account.NewAccount(login, password, url)
	// files.ReadFile()
	if (err != nil) {
		fmt.Println("'Неверный формат логин или урл'")
		return
	}
	file, err := myAccount.ToByte()
	if err != nil {
		fmt.Println("Не удалось преобразовать данные")
	}
	files.WriteFile(string(file), "data.json")
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