package connection

import (
	"database/sql"

	"os"

	_ "github.com/go-sql-driver/mysql" // Importa el controlador MySQL
)

var DB *sql.DB

func ConnectDatabase() {
	// Abrir una conexión a la base de datos SQLite
	host := os.Getenv("DATABASE")

	var DSN = "root:12345@tcp(" + host + ":3306)/database"

	// Abre una conexión a la base de datos

	db, err := sql.Open("mysql", DSN)
	if err != nil {
		panic(err.Error())
	}

	// Verificar que la conexión a la base de datos sea exitosa

	// Asignar la conexión a la variable global DB
	DB = db

}
