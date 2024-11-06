package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Datos de conexión a MySQL
const (
	mysqlUser     = "tu_usuario"
	mysqlPassword = "tu_contraseña"
	mysqlHost     = "localhost"
	mysqlPort     = "3306"
)

// Crear la base de datos si no existe
func CrearBaseDeDatos() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/", mysqlUser, mysqlPassword, mysqlHost, mysqlPort)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("Error al conectar a MySQL:", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS ecommerce")
	if err != nil {
		log.Println("Error al crear la base de datos:", err)
	} else {
		log.Println("Base de datos 'ecommerce' creada o ya existente.")
	}
}

// Conectar a la base de datos ecommerce
func ConectarBD() (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/ecommerce", mysqlUser, mysqlPassword, mysqlHost, mysqlPort)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Crear la tabla de productos si no existe
func CrearTablas(db *sql.DB) {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS productos (
        id INT AUTO_INCREMENT PRIMARY KEY,
        nombre VARCHAR(100) NOT NULL,
        descripcion TEXT,
        precio DECIMAL(10, 2) NOT NULL,
        stock INT NOT NULL
    )`)
	if err != nil {
		log.Fatal("Error al crear la tabla productos:", err)
	}

	log.Println("Tabla 'productos' creada exitosamente o ya existente.")
}
