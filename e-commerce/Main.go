package main

import (
	"ecommerce/config"
	"ecommerce/models"
	"fmt"
)

func main() {
	// Crear la base de datos y la tabla de productos si no existen
	config.CrearBaseDeDatos()

	// Conectar a la base de datos ecommerce
	db := config.ConectarBD()
	if db == nil {
		fmt.Println("Error al conectar a la base de datos ecommerce")
		return
	}

	// Crear un nuevo producto como ejemplo
	producto := models.Producto{
		Nombre:      "libro de dibujo",
		Descripcion: "Libro para dibujar",
		Precio:      9.99,
		Stock:       10,
	}

	producto.CrearProducto(db)

	// Imprimir mensaje de éxito
	fmt.Println("Producto creado exitosamente")
}
