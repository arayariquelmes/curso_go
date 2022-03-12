package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		//fmt.Printf("%s esta funcionando \n", servidor)
		canal <- servidor + "esta funcionando \n"
	} else {
		canal <- servidor + "no esta funcionando \n"
	}
}
func main() {
	inicio := time.Now()

	canal := make(chan string)
	servidores := []string{
		"https://skynux.cl/",
		"https://www.udemy.com",
		"https://www.facebook.com",
	}

	for _, servidor := range servidores {
		//La palabra go se llama Go Runtime y ejecuta operaciones concurrente
		go revisarServidor(servidor, canal)
	}
	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}
	tiempoPaso := time.Since(inicio)
	fmt.Println("Tiempo de ejecucion:", tiempoPaso)

}
