package models

import (
	"database/sql"
)

type Producto struct {
	ID          int     `json:"id"`
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
	Stock       int     `json:"stock"`
}

// Crear un producto en la base de datos
func (p *Producto) CrearProducto(db *sql.DB) error {
	query := "INSERT INTO productos (nombre, descripcion, precio, stock) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, p.Nombre, p.Descripcion, p.Precio, p.Stock)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = int(id)
	return nil
}
