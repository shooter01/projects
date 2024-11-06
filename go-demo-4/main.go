package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
) 



func main() {
	createAccount()
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
		fmt.Println("Не удлаось преобразовать данные")
	}
	files.WriteFile(string(file), "data.json")
	// color.RGB(255, 128, 0).Println(myAccount)
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