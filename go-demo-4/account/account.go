package account

import (
	"encoding/json"
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

func (acc *Account) ToByte() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil

}


func (acc *Account) outputPassword() {
	// acc.password = generatePassword(10)
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
