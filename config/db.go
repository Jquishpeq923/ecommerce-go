package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Importa el driver de MySQL
)

// Configuración de conexión a MySQL
const (
	mysqlUser     = "root"
	mysqlPassword = "Nayelimia"
	mysqlHost     = "localhost"
	mysqlPort     = "3306"
)

// Crear la base de datos ecommerce si no existe
func CrearBaseDeDatos() {
	// Cadena de conexión
	connStr := fmt.Sprintf("root:Nayelimia@tcp(localhost:3306)", mysqlUser, mysqlPassword, mysqlHost, mysqlPort)
	// Abrimos la conexión
	db, _ := sql.Open("mysql", connStr)

	// Crear la base de datos ecommerce si no existe
	db.Exec("CREATE DATABASE IF NOT EXISTS ecommerce")
	fmt.Println("Base de datos 'ecommerce' creada con éxito o ya existía.")
	db.Close()
}

// Conectar a la base de datos ecommerce
func ConectarBD() *sql.DB {
	// Cadena de conexión
	connStr := fmt.Sprintf("root:Nayelimia@tcp(localhost:3306)", mysqlUser, mysqlPassword, mysqlHost, mysqlPort)
	db, _ := sql.Open("mysql", connStr)
	return db
}

// Crear las tablas necesarias
func CrearTablas(db *sql.DB) {
	// Crear la tabla productos
	db.Exec(`
    CREATE TABLE IF NOT EXISTS productos (
        id INT AUTO_INCREMENT PRIMARY KEY,
        nombre VARCHAR(100) NOT NULL,
        descripcion TEXT,
        precio DECIMAL(10, 2) NOT NULL,
        stock INT NOT NULL
    )`)

	// Crear la tabla carrito
	db.Exec(`
    CREATE TABLE IF NOT EXISTS carrito (
        id INT AUTO_INCREMENT PRIMARY KEY,
        usuario_id INT NOT NULL,
        producto_id INT NOT NULL,
        cantidad INT NOT NULL,
        FOREIGN KEY (producto_id) REFERENCES productos(id)
    )`)

	// Crear la tabla pedidos
	db.Exec(`
    CREATE TABLE IF NOT EXISTS pedidos (
        id INT AUTO_INCREMENT PRIMARY KEY,
        usuario_id INT NOT NULL,
        total DECIMAL(10, 2) NOT NULL,
        status VARCHAR(50) NOT NULL
    )`)

	fmt.Println("Tablas creadas exitosamente.")
}
