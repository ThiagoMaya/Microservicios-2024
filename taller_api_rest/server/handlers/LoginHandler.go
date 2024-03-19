package handlers

import (
	"database/sql"
	"encoding/json"
	"microservicios/taller_api/connection"
	"microservicios/taller_api/model"
	"microservicios/taller_api/security"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Parsear la solicitud y obtener las credenciales del usuario
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	query := "SELECT id FROM users WHERE username = ? AND password = ?"
	var userID int
	err = connection.DB.QueryRow(query, user.Username, user.Password).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Usuario no encontrado o credenciales incorrectas
			http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
			return
		}
		// Otro error al ejecutar la consulta
		http.Error(w, "Error al verificar las credenciales del usuario", http.StatusInternalServerError)
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
