package main

import "fmt"

func sumar(nombre string, numeros ...int) (string, int) {
	mensaje := fmt.Sprintf("Wena papu %s", nombre)
	fmt.Println(numeros, len(numeros))
	var total int
	for _, num := range numeros {
		total += num
	}
	return mensaje, total
}

func main() {

	//Funcion anonima que se ejecuta sola
	func() {
		fmt.Println("Hola papu anonima")
	}()
	//Con nombre
	hola := func() {
		fmt.Println("Hola papu con nombre")
	}
	hola()
	sumar("Seba")
	mensaje, result := sumar("Seba", 1, 2, 3, 5, 6, 7)
	fmt.Println(mensaje, result)
}
