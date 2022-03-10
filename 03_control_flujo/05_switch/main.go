package main

import "fmt"

func main() {
	var letra string
	fmt.Println("Ingrese letra")
	fmt.Scanln(&letra)
	fmt.Println("Letra", letra)
	switch {
	case letra == "A" || letra == "a":
		fmt.Println("Es una A")
	case letra == "E" || letra == "e":
		fmt.Println("Es una E")
	default:
		fmt.Println("Mega F")
	}
}
