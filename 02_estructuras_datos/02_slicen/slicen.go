package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3} //No se define cuantos elementos
	//los slicen son como las listas de python
	fmt.Println(numeros, len(numeros))

	//agregar datos
	numeros = append(numeros, 4)
	fmt.Println(numeros)
	//Sub slicen
	subSlicen := numeros[:2]
	numeros[0] = 100
	fmt.Println(numeros, subSlicen)
	//Puneros
	//Longitud
	//Capacidad

	meses := []string{"Enero", "Febrero", "Marzo"}
	fmt.Printf("LenL %v, Cap: %v, %p \n", len(meses), cap(meses), meses)
	meses = append(meses, "Abril")
	fmt.Printf("LenL %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

}
