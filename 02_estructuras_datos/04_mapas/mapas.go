package main

import "fmt"

func main() {
	dias := make(map[string]string)
	fmt.Println(dias)
	dias["L"] = "Lunes"
	dias["M"] = "Martes"
	dias["MIE"] = "Miercoles"
	fmt.Println(dias)
	//Eliminar elemento de mapa
	delete(dias, "L")
	fmt.Println(dias)
	fmt.Println(dias, len(dias))

	//Map mas complejo
	estudiantes := make(map[string][]int)
	estudiantes["Seba"] = []int{1, 2, 3}
	estudiantes["Araya"] = []int{1, 5, 6}
	fmt.Println(estudiantes)
}
