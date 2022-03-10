package main

import "fmt"

func main() {
	var hola string = "Hola"
	mundo := "Mundo"
	fmt.Print(hola) //Imprime en una sola linea, obvio
	fmt.Print(mundo)

	nombre := "Seba"
	edad := 34
	fmt.Printf(" Hola , %s esto es igual que C por ende la edad debe ser %d \n", nombre, edad)
	fmt.Printf(" Hola , %v esto es igual que C pero v es pa hacerla flaite por ende la edad debe ser %v \n",
		nombre, edad)
	mensaje := fmt.Sprintf("Hola %s", nombre)
	fmt.Println(mensaje)

	fmt.Printf("nombre: %T", nombre)
	tipoDatoNombre := fmt.Sprintf("%T", nombre)
	fmt.Printf(tipoDatoNombre)

	fmt.Println("Ingrese nombre mi ciela")
	fmt.Scanln(&nombre) //el scanf mi ciela
	fmt.Println("Hola mi ciela", nombre)
}
