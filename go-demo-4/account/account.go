package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"strings"
)

type Account struct {
	Login    string `json:"login"`
	Password string	`json:"password"`
	Url      string	`json:"url"`
}


func (acc *Account) outputPassword() {
	// acc.password = generatePassword(10)
}


func (acc *Account) Output() {
	fmt.Println(acc.Login)
	fmt.Println(acc.Password)
	fmt.Println(acc.Url)
}

func NewAccount(login, password, urlString string) (*Account, error) {

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	newAcc := &Account{
		Url:      urlString,
		Login:    login,
		Password: password,
	}

	if password == "" {
		newAcc.Password = generatePassword(10)
	}

	return newAcc, nil

}

func generatePassword(number int) string {
	stringArray := make([]string, number)
	for i := range stringArray {
		stringArray[i] = string(rand.IntN(1000)) 
	}
	// for i := 0; i < len(stringArray); i++ {
	// 	stringArray[i] = string(rand.IntN(1000)) 
	// }
	fmt.Println(stringArray)
	fmt.Println(rand.IntN(10))
	return strings.Join(stringArray, "")
}
