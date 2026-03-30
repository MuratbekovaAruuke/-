package main

import "fmt"

func main() {
	fmt.Println("Привет, мир! Я люблю тебя ❤")

	age := 17
	fmt.Printf("Через 5 лет мне будет %d лет\n", age+5)

	names := []string{"Анеля", "Нурболот", "Аэлита"}
	fmt.Println("Список друзей:")
	for _, name := range names {
		fmt.Println("-", name)
	}
}
