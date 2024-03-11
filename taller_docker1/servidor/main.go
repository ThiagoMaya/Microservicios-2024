package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Definir un manejador (handler) para las solicitudes HTTP

	http.HandleFunc("/saludo", handlerSaludo)

	http.HandleFunc("/login", LoginHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Recurso no encontrado", http.StatusNotFound)
	})

	fmt.Println("Servidor corriendo en el puerto 80")

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		return
	}

}

func handlerSaludo(w http.ResponseWriter, r *http.Request) {

	nombre := r.URL.Query().Get("nombre")

	if nombre == "" {

		http.Error(w, "Solicitud no válida: El nombre es obligatorio", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	mensaje := fmt.Sprintf("Hola, %s", nombre)

	fmt.Fprintln(w, mensaje)
}
