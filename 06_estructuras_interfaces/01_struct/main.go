package main

import (
	"fmt"
	"strconv"
)

//Struct persona

type Persona struct {
	nombre string
	edad   int
}

func (this *Persona) saludar() string {
	return this.nombre + " papu " + strconv.Itoa(this.edad)
}

func (this *Persona) setNombre(nombre string) {
	this.nombre = nombre
}

//Herencia

type Empleado struct {
	Persona
	sueldo float64
}

func main() {
	p1 := Persona{"Seba", 33}

	p1.setNombre("Fernando")
	fmt.Println(p1)

	p2 := Persona{
		nombre: "Seba",
		edad:   27}
	fmt.Println(p2)
	fmt.Println(p2.saludar())

	empleado1 := Empleado{}
	empleado1.edad = 12
	empleado1.nombre = "Perkin"
	fmt.Println(empleado1.saludar())

}
