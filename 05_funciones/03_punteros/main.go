package main

import "fmt"

func modificarValor(cadena *string) {
	fmt.Printf("%p\n", &cadena)
	*cadena = "Hola funcioncita papu que modifica referencia"
}

func main() {
	cadena := "Hola mundo"
	fmt.Printf("%p\n", &cadena)
	fmt.Println("Antes de funcion:", cadena)
	modificarValor(&cadena)
	fmt.Println("Despues de la funcion:", cadena)
}
