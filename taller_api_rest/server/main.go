package main

import (
	"fmt"
	"log"
	"microservicios/taller_api/connection"
	"microservicios/taller_api/handlers"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Importa el controlador MySQL
	"github.com/gorilla/mux"
)

func main() {

	// Verifica si se puede conectar a la base de datos
	connection.ConnectDatabase()

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/recoverPassword", handlers.ForgotPassword).Methods("GET") //prefix para ruta users

	apiRouter := router.PathPrefix("/users").Subrouter()

	//apiRouter.HandleFunc("/", HomeHandler).Methods("GET")
	apiRouter.HandleFunc("/", handlers.PostUserHandler).Methods("POST")
	apiRouter.HandleFunc("/", handlers.GetUsersHandler).Methods("GET")
	apiRouter.HandleFunc("/{id}", handlers.GetUserByIDHandler).Methods("GET")
	apiRouter.HandleFunc("/{id}", handlers.DeleteUserHandler).Methods("DELETE")

	server := &http.Server{
		Addr:    ":8080", // Puerto en el que el servidor escuchará
		Handler: router,  // Usamos el enrutador mux
	}

	fmt.Println("Servidor escuchando en el puerto 8080...")
	// Iniciamos el servidor
	log.Fatal(server.ListenAndServe())

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Bienvenido a mi API con Go y mux!")
}
