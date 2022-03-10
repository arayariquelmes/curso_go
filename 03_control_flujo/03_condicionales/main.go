package main

import "fmt"

func main() {

	var consumo, descuento float64
	var datosDescuento string

	fmt.Println("Ingrese consumo:")
	fmt.Scanln(&consumo)

	if consumo >= 0 {
		//Descuento 10%
		if consumo <= 100 {
			datosDescuento = "10%"
			descuento = consumo * 10 / 100
		} else if consumo > 100 && consumo <= 200 {
			datosDescuento = "20%"
			descuento = consumo * 20 / 100

		} else {
			datosDescuento = "30%"
			descuento = consumo * 30 / 100
		}
		fmt.Printf("Descuento: %f , Consumo: %f , Datos: %s", descuento, consumo, datosDescuento)
	} else {
		fmt.Println("Error al ingresar consumo")
	}

}
