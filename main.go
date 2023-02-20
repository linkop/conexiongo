package main

import (
	"conexiongo/db"
	"conexiongo/models"
	"fmt"
)

func main() {
	db.Connect()

	//crear una tabla desde go
	//db.CreateTable(models.LibroSchema)
	//db.Ping()

	//ingresar datos en la tabla desde go
	user := models.CreateUser("final", "final", "final@gmail.com")
	fmt.Println(user)

	//mostrar todas los campos de la tabla
	//users := models.ListaUsers()
	//fmt.Println(users)

	//mostrar filtrado solo un campo de la tabla
	//user := models.GetUser(4)
	//fmt.Println(user)

	//ahora vamos a actualizar el campo que fue filtrado
	//user.Username = "Linkp"
	//user.Save()
	//fmt.Println(models.ListaUsers())

	//eliminar registro
	//user.Delete()
	//fmt.Println(models.ListaUsers())
	//db.TruncateTable("libros")
	db.Close()
}
