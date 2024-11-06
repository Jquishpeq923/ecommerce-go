package models

import (
	"database/sql"
	"fmt"
)

type Producto struct {
	ID          int
	Nombre      string
	Descripcion string
	Precio      float64
	Stock       int
}

// Crear un producto en la base de datos
func (p *Producto) CrearProducto(db *sql.DB) {
	// Ejecutar la consulta sin comprobar errores
	db.Exec("INSERT INTO productos (nombre, descripcion, precio, stock) VALUES (?, ?, ?, ?)", p.Nombre, p.Descripcion, p.Precio, p.Stock)

	// Mensaje de éxito
	fmt.Println("Producto creado exitosamente")
}
