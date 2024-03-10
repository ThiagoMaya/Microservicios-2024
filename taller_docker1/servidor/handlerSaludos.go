package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("clave_bien_secreta") // Cambia esto a tu clave secreta

func SaludoHandler(w http.ResponseWriter, r *http.Request) {

	nombre := r.URL.Query().Get("nombre")

	if nombre == "" {
		http.Error(w, "Solicitud no válida: el nombre es obligatorio", http.StatusBadRequest)
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "No se proporcionó el token JWT", http.StatusUnauthorized)
		return
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Método de firma inesperado: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		http.Error(w, "Token JWT inválido", http.StatusUnauthorized)
		return
	}

	// Verificar si el token JWT es válido y obtener las reclamaciones
	if !token.Valid {
		http.Error(w, "Token JWT inválido", http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "No se pudieron obtener las reclamaciones del token JWT", http.StatusInternalServerError)
		return
	}

	// Verificar el emisor del token
	emisor, ok := claims["iss"].(string)
	if !ok || emisor != "ingesis.uniquindio.edu.co" {
		http.Error(w, "Emisor del token no válido", http.StatusUnauthorized)
		return
	}

	// Verificar si el nombre en el parámetro de la ruta coincide con el identificador en el token
	usuario, ok := claims["sub"].(string)
	if !ok || nombre != usuario {
		http.Error(w, "Nombre en el parámetro de la ruta no coincide con el identificador en el token", http.StatusUnauthorized)
		return
	}

	// Si todas las verificaciones son exitosas, responder con un mensaje de saludo
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Hola %s", nombre)

}
