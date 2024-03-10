package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// Credenciales es la estructura para las credenciales de usuario
type Credentials struct {
	Usuario string `json:"usuario"`
	Clave   string `json:"clave"`
}

// Read implements io.Reader.
func (Credentials) Read(p []byte) (n int, err error) {
	panic("unimplemented")
}

// Respuesta es la estructura para la respuesta del servidor
type Respuesta struct {
	Mensaje string `json:"mensaje"`
}

func main() {
	// servidorURL := os.Getenv("SERVIDORGO")
	cliente := &http.Client{}
	URL := "http://localhost:80/saludo?nombre=santi"

	// if servidorURL == "" {
	// 	fmt.Println("Variable de entorno SERVIDORGO no configurada")
	// 	return
	// }

	rand.Seed(time.Now().UnixNano())
	usuario := "santiago"
	clave := "contrasenia"

	// Enviar solicitud a la ruta /login
	credenciales := Credentials{Usuario: usuario, Clave: clave}
	fmt.Println(credenciales)

	var cuerpo []byte
	jsonDatos, _ := json.Marshal(credenciales)
	cuerpo = jsonDatos

	// Se crea una request para enviarla al servidor
	request, err := http.NewRequest("GET", URL, bytes.NewBuffer(cuerpo))
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Add("Content-Type", "application/json")

	response, err := cliente.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("---Imprimiendo respuesta del servidor---")
	fmt.Println(string(body))
}
