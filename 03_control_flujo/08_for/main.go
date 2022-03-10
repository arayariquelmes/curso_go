package main

import "fmt"

func main() {
	//Bucle infinito
	// for{}

	//Bucle tipo while
	numeros := 12455
	var c int
	for numeros > 0 {
		numeros /= 10
		c++
	}
	fmt.Println("Cantidad digitos", c)

	//Bucle for
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	//ForEach

	nombres := make([]string, 0, 1)
	nombres = append(nombres, "Seba")
	nombres = append(nombres, "Araya")
	for i := 0; i < len(nombres); i++ {
		fmt.Println(nombres[i])
	}

	for indice, elemento := range nombres {
		fmt.Println(indice, elemento)
	}

	for _, elemento := range nombres {
		fmt.Println(elemento)
	}
}
