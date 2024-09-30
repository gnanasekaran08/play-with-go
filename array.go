package main

import "fmt"

func main() {
	laptops := []string{"Dell", "HP", "Lenovo", "Apple"}

	phones := [][]string{{"Samsung", "Apple"}, {"OnePlus", "Xiaomi"}}

	for laptop_index := 0; laptop_index < len(laptops); laptop_index++ {
		println("Laptop: ", laptops[laptop_index])
	}

	for phone_index := 0; phone_index < len(phones); phone_index++ {
		fmt.Println("Phone Index: ", phones[phone_index])
		for phone_brand_index := 0; phone_brand_index < len(phones[phone_index]); phone_brand_index++ {
			println("Phone: ", phones[phone_index][phone_brand_index])
		}

	}
}
