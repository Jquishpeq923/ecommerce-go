package main

import (
	"ecommerce/config"
	"ecommerce/models"
	"fmt"
	"log"
)

func main() {
	// Crear la base de datos si no existe
	config.CrearBaseDeDatos()

	// Conectar a la base de datos ecommerce
	db, err := config.ConectarBD()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos ecommerce:", err)
	}
	defer db.Close()

	// Crear la tabla de productos si no existe
	config.CrearTablas(db)

	// Crear un nuevo producto como ejemplo
	producto := models.Producto{
		Nombre:      "Producto de ejemplo",
		Descripcion: "Este es un producto de ejemplo",
		Precio:      9.99,
		Stock:       10,
	}

	err = producto.CrearProducto(db)
	if err != nil {
		log.Fatal("Error al crear producto:", err)
	}

	fmt.Printf("Producto creado exitosamente con ID: %d\n", producto.ID)
}
