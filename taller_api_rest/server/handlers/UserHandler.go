package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"microservicios/taller_api/connection"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	result, err := connection.DB.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", newUser.Username, newUser.Password, newUser.Email)
	if err != nil {
		http.Error(w, "Error al registrar el usuario en la base de datos", http.StatusInternalServerError)
		log.Println("Error al registrar el usuario en la base de datos:", err)
		return
	}

	id, _ := result.LastInsertId()
	newUser.ID = int(id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)

}

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID de usuario no válido", http.StatusBadRequest)
		return
	}

	var user User
	err = connection.DB.QueryRow("SELECT id, username, password FROM usuarios WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Usuario no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Error al obtener el usuario de la base de datos", http.StatusInternalServerError)
		log.Println("Error al obtener el usuario de la base de datos:", err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := connection.DB.Query("SELECT id, username FROM usuarios")
	if err != nil {
		http.Error(w, "Error al obtener usuarios de la base de datos", http.StatusInternalServerError)
		log.Println("Error al obtener usuarios de la base de datos:", err)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username)
		if err != nil {
			http.Error(w, "Error al escanear usuarios de la base de datos", http.StatusInternalServerError)
			log.Println("Error al escanear usuarios de la base de datos:", err)
			return
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID de usuario no válido", http.StatusBadRequest)
		return
	}

	_, err = connection.DB.Exec("DELETE FROM usuarios WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Error al eliminar el usuario de la base de datos", http.StatusInternalServerError)
		log.Println("Error al eliminar el usuario de la base de datos:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Usuario con ID %d eliminado correctamente\n", id)
}
