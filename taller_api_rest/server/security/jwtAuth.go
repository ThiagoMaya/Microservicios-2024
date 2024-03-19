package security

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("clave_bien_secreta")

func CreateToken(username string) string {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = username
	claims["exp"] = time.Now().Add(time.Hour).Unix() // Token válido por una hora
	claims["iss"] = "ingesis.uniquindio.edu.co"

	// Firmar el token con una clave secreta y obtener el string del token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println("No se ha podido firmar el token")

	}

	return tokenString

}

func VerifyToken(w http.ResponseWriter, r *http.Request) bool {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "No se proporcionó el token JWT", http.StatusUnauthorized)
		return false
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
		return false
	}

	// Verificar si el token JWT es válido y obtener las reclamaciones
	if !token.Valid {
		http.Error(w, "Token JWT inválido", http.StatusUnauthorized)
		return false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "No se pudieron obtener las reclamaciones del token JWT", http.StatusInternalServerError)
		return false
	}

	// Verificar el emisor del token
	emisor, ok := claims["iss"].(string)
	if !ok || emisor != "ingesis.uniquindio.edu.co" {
		http.Error(w, "Emisor del token no válido", http.StatusUnauthorized)
		return false
	}

	return true

}
