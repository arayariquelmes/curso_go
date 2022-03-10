package main

//Para generar el modulo se usa go mod init nombrePaquete
import (
	"fmt"
	"paquetes/mensajes"
)

func main() {
	fmt.Println(mensajes.Hola())
}
