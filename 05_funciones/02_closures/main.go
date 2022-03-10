package main

import (
	"fmt"
	"strings"
)

//Closure, retorna una funcion, retorna una funcion de tipo string
func repeat(n int) func(cadena string) string {
	return func(cadenita string) string {
		return strings.Repeat(cadenita, n)
	}
}

func main() {

	resultado := repeat(5)("Hola")
	fmt.Println(resultado)
}
