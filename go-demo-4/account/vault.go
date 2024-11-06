package account

import (
	"demo/password/files"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() (*Vault){
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault {
			Accounts: []Account{},
			UpdatedAt: time.Now(),
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
		color.Red(err.Error())
	}
	return &vault
}


func (vault *Vault) FindAccountsByURL(url string) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if (isMatched) {
			accounts = append(accounts, account)
		}
	}
	return accounts
}



func (vault *Vault) AddAccount(acc Account){
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()

	data, err := vault.ToByte()
	if err != nil {
		color.Red("Не удалось преобразовать файл data.json")

		color.Red(err.Error())
	}

	files.WriteFile(string(data), "data.json")
}


func (vault *Vault) ToByte() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil

}
