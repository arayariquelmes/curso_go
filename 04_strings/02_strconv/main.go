package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num string
	fmt.Print("Ingrese numero:")
	fmt.Scanln(&num)

	entero, error := strconv.Atoi(num)
	flotante, errorF := strconv.ParseFloat(num, 64)
	fmt.Printf("Numero int %d numero float %.2f  error int %T error float %T",
		entero, flotante, error, errorF)
}
