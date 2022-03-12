package main

import (
	"fmt"
	"log"
	"net/http"
)

//handler
func hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("El metodo http usado es " + r.Method)
	fmt.Fprintln(rw, "Hola mi ciela")
}

func paginaNoFunciona(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}
func Error(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "F MI CIELA", http.StatusBadRequest)
}

func saludar(rw http.ResponseWriter, r *http.Request) {
	//Devuelve url del endpoint
	fmt.Println(r.URL)
	//Devuelve en un string los argumentos
	fmt.Println(r.URL.RawQuery)
	mapaParametros := r.URL.Query()
	fmt.Println(mapaParametros)
	fmt.Fprintf(rw, "Nombre valor: %s\n", mapaParametros.Get("nombre"))
}
func main() {
	//Mux ruta asociada a handler
	mux := http.NewServeMux()

	mux.HandleFunc("/page404", paginaNoFunciona)
	mux.HandleFunc("/errorRequest", Error)
	mux.HandleFunc("/saludar", saludar)
	mux.HandleFunc("/", hola)
	//Router con mux por defecto
	/*http.HandleFunc("/page404", paginaNoFunciona)
	http.HandleFunc("/errorRequest", Error)
	http.HandleFunc("/saludar", saludar)
	http.HandleFunc("/", hola)*/

	//Crear Servidor
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	fmt.Println("Server corriendo en el puerto 8080")
	fmt.Println("Run server: http://localhost:8080")
	//Server por defecto
	//log.Fatal(http.ListenAndServe("localhost:8080", mux))
	log.Fatal(server.ListenAndServe())

}
