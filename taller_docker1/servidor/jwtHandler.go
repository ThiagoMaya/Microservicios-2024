package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Credencial struct {
	Usuario string `json:"usuario"`
	Clave   string `json:"clave"`
}

// Manejador para la ruta /login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Verificar si la solicitud es de tipo POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Parsear el cuerpo de la solicitud para obtener las credenciales
	var creds Credencial
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Error al leer las credenciales", http.StatusBadRequest)
		return
	}

	// Verificar si se proporcionaron usuario y clave
	if creds.Usuario == "" || creds.Clave == "" {
		http.Error(w, "Faltan usuario y claves", http.StatusBadRequest)
		return
	}

	// Generar el token JWT
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = creds.Usuario
	claims["exp"] = time.Now().Add(time.Hour).Unix() // Token válido por una hora
	claims["iss"] = "ingesis.uniquindio.edu.co"

	// Firmar el token con una clave secreta y obtener el string del token
	tokenString, err := token.SignedString([]byte("clave_bien_secreta"))
	if err != nil {
		http.Error(w, "Error al firmar el token", http.StatusInternalServerError)
		return
	}

	// Responder con el token JWT
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, tokenString)
}
