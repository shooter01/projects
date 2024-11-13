package main

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"flag"
	"fmt"
)

func main() {
	fmt.Println("новый проект")
	name := flag.String("name", "Anton", "Имя пользователя")
	city := flag.String("city", "", "Город")
	age := flag.Int("age", 18, "Возраст пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")
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

	weatherData := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
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
