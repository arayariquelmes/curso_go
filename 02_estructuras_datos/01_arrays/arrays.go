package main

import "fmt"

func main() {

	var numeros [5]int
	fmt.Println(numeros)
	numeros[0] = 12
	fmt.Println(numeros[0])

	//Arreglo definido con valores
	nombres := [5]string{"Seba", "Araya"}
	fmt.Println(nombres)
	colores := [...]string{ //infiere cantidad de elementos
		"Rojo", "Verde", "Azul",
	}

	fmt.Println(colores, len(colores))

	//Indices definidos
	monedas := [...]string{
		0: "Dolar",
		2: "Soles",
		3: "Euros",
		7: "asd",
	}
	fmt.Println(monedas, len(monedas))
	subMoneda := monedas[0:3] //slicing como en python
	fmt.Println(subMoneda, len(subMoneda))
}
