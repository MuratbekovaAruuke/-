package main

import "fmt"

func main() {
	ponies := map[int]string{
		1: "Twilight Sparkle",
		2: "Rainbow Dash",
		3: "Applejack",
		4: "Fluttershy",
		5: "Pinkie Pie",
		6: "Rarity",
	}
	fmt.Println("Все пони:")
	for _, name := range ponies {
		fmt.Println("-", name)
	}
}
