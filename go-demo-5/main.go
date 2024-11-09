package main

import (
	"demo/weather/geo"
	"flag"
	"fmt"
)

func main() {
	fmt.Println("новый проект")
	name := flag.String("name", "Anton", "Имя пользователя")
	city := flag.String("city", "", "Город")
	age := flag.Int("age", 18, "Возраст пользователя")
	isAdmin := flag.Bool("isAdmin", false, "Является админом")
	flag.Parse()

	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*isAdmin)
	fmt.Println(*city)

	geoData, err := geo.GetMyLocation(*city)

	if err != nil {
		fmt.Println("Ошибка")

	}
	fmt.Println(geoData)

	// r := strings.NewReader("Привет! Я поток данных")

	// block := make([]byte, 4)

	// for {
	// 	_, err := r.Read(block)
	// 	fmt.Printf("%q\n", block)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }
}
