package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"microservicios/taller_api/connection"
	"microservicios/taller_api/model"
	"microservicios/taller_api/security"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func PostUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if !security.VerifyToken(w, r) {

		return
	}

	var newUser model.User
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

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if !security.VerifyToken(w, r) {

		return
	}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID de usuario no válido", http.StatusBadRequest)
		return
	}

	var user model.User

	err = connection.DB.QueryRow("SELECT id, username, password FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Password)
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
	// Verificar el método de la solicitud
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Verificar la validez del token de autenticación
	if !security.VerifyToken(w, r) {
		return
	}

	// Parámetros de paginación
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	// Página predeterminada si no se proporciona
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	// Tamaño predeterminado de la página si no se proporciona
	size := 10
	if sizeStr != "" {
		size, _ = strconv.Atoi(sizeStr)
	}

	// Calcular el offset
	offset := (page - 1) * size

	// Consulta SQL con limit y offset para la paginación
	query := "SELECT id, username FROM users LIMIT ? OFFSET ?"
	rows, err := connection.DB.Query(query, size, offset)
	if err != nil {
		http.Error(w, "Error al obtener usuarios de la base de datos", http.StatusInternalServerError)
		log.Println("Error al obtener usuarios de la base de datos:", err)
		return
	}
	defer rows.Close()

	// Obtener los usuarios de la consulta
	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Username)
		if err != nil {
			http.Error(w, "Error al escanear usuarios de la base de datos", http.StatusInternalServerError)
			log.Println("Error al escanear usuarios de la base de datos:", err)
			return
		}
		users = append(users, user)
	}

	// Respuesta con los usuarios y la información de paginación
	response := map[string]interface{}{
		"page":  page,
		"size":  size,
		"total": len(users), // Aquí deberías obtener el total de usuarios de la base de datos
		"users": users,
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if !security.VerifyToken(w, r) {

		return
	}
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID de usuario no válido", http.StatusBadRequest)
		return
	}

	_, err = connection.DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Error al eliminar el usuario de la base de datos", http.StatusInternalServerError)
		log.Println("Error al eliminar el usuario de la base de datos:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Usuario con ID %d eliminado correctamente\n", id)
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	query := "SELECT * FROM users WHERE email = ?"
	_, err = connection.DB.Exec(query, user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Usuario no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Error al obtener el usuario de la base de datos", http.StatusInternalServerError)
		log.Println("Error al obtener el usuario de la base de datos:", err)
		return
	}

	tokenString := security.CreateToken(user.Username)

	response := map[string]string{"token": tokenString}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
