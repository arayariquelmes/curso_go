package main

import "fmt"

func main() {
	if nombre, edad := "Alex", 28; nombre == "Alex" {
		fmt.Println("Hola", nombre)
	} else {
		fmt.Println(nombre, edad)
	}
	users := make(map[string]string)

	users["Alex"] = "alex@gmail.com"
	users["Seba"] = "seba@gmail.com"
	fmt.Println(users["Juan"])
	correo, ok := users["Alex"] //Ok devuelve si el valor existe
	fmt.Println(correo, ok)

	if correo2, ok2 := users["Juan"]; ok2 {
		fmt.Println(correo2)
	} else {
		fmt.Println("No fue posible")
	}
}
