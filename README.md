# Proyecto de E-commerce en Go

Este es un avance de un sistema de e-commerce en Go con MySQL. Actualmente, el proyecto permite crear la base de datos y la tabla de productos, así como insertar un producto en la base de datos.

## Estructura del Proyecto

- `config/db.go`: Configuración y conexión de la base de datos.
- `models/producto.go`: Definición del modelo de producto y su inserción en la base de datos.
- `main.go`: Archivo principal para inicializar la base de datos, crear tablas, y probar la inserción de un producto.

## Requisitos

- Go 1.23.2
- MySQL.

## Configuración

1. Asegúrate de instalar el driver de MySQL con:
   ```bash
   go get -u github.com/go-sql-driver/mysql
