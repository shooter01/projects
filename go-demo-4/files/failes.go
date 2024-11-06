package files

import (
	"fmt"
	"os"
)


func WriteFile(content string, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
	
}

func ReadFile(name string) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}