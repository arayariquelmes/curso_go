package main

import (
	"fmt"
	"strings"
)

func reverse(palabra string) string {
	arrayCadena := strings.Split(palabra, "")
	arrayInvertida := make([]string, 0)
	var i int
	for i = len(arrayCadena) - 1; i >= 0; i-- {
		arrayInvertida = append(arrayInvertida, arrayCadena[i])
	}
	fmt.Println(arrayInvertida)
	return strings.Join(arrayInvertida, "")
}

func esPalindromo(palabra string) {
	fmt.Println("Palabra", palabra)
	palabra = strings.ToLower(palabra)
	fmt.Println("Replace")
	palabra = strings.ReplaceAll(palabra, " ", "")
	fmt.Println(palabra)
	var palabraInvertida string = reverse(palabra)
	fmt.Println("Esto fue invertida", palabraInvertida)
}

func main() {
	esPalindromo("hOLA MUNDO")
}
