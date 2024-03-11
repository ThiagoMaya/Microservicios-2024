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

	// Manejador para la ruta '/'
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/create", handlers.PostUserHandler).Methods("POST")

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

/*
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    // Verificamos si la solicitud es de tipo POST
    if r.Method != http.MethodPost {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    // Decodificamos los datos de la solicitud
    var newUser User
    err := json.NewDecoder(r.Body).Decode(&newUser)
    if err != nil {
        http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
        return
    }

    // Aquí puedes realizar la lógica de registro del usuario, como almacenar en una base de datos, etc.
    // Simplemente para demostración, imprimimos los datos recibidos
    fmt.Printf("Nuevo usuario registrado:\nUsuario: %s\nContraseña: %s\n", newUser.Username, newUser.Password)

    // Respondemos con un mensaje de éxito
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "¡Usuario %s registrado correctamente!\n", newUser.Username)
}
*/
