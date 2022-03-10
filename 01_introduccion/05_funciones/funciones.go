package main

import "fmt"

func saludar(nombre string) {
	fmt.Println("Hola Rey", nombre)
}

func getSaludo(nombre string) string {
	return "Hola Rey" + nombre
}
func sumar(a, b int) int {
	return a + b
}

func main() {

	saludar("Seba")
	fmt.Println(getSaludo("Seba mi ciela"))
	fmt.Println(sumar(12, 12))
}
