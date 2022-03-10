package main

import (
	"fmt"
	"os"
)

func main() {

	//funcion recover
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Esta custion se fue a la shieet, esto es como un catch")
		}
	}()

	if file, error := os.Open("hola.csv"); error != nil {
		panic("Exploto esta custion") //Equivalente al throws
	} else {
		//Se ejecuta al final de la funcion
		defer func() {
			fmt.Println("Cerrando Archivo")
			file.Close()
		}()
		contenido := make([]byte, 254)
		longitud, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:longitud])
		fmt.Println(contenidoArchivo)

	}
}
