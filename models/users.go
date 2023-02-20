package models

import (
	"conexiongo/db"
)

type User struct {
	Id       int64
	Username string
	Password string
	Email    string
}

type Users []User

const UserSchema string = `create table users (
	id serial,
	username varchar(30) not null,
	password varchar(100) not null,
	email varchar(50),
	primary key(id))`

// construir usuario
func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}
	return user
}

func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	//antes del cambio
	//user.insert()
	user.Save()
	return user
}

// insertar registro
func (user *User) insert() {
	sql := `insert into users(username,password,email) VALUES ($1,$2,$3)`
	result, _ := db.Exec(sql, user.Username, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}

// Listar todos los registros
func ListaUsers() Users {
	//sql := `select id,username,password,email from users` otra forma de hacerlo
	sql := `select * from users`
	users := Users{}
	rows, _ := db.Query(sql)

	//funciona para recorrer la tabla
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}
	return users
}

// obtener un registro
func GetUser(id int) *User {
	user := NewUser("", "", "")
	sql := `select id,username,password,email from users where id=$1`
	rows, _ := db.Query(sql, id)
	//recorremos la tabla para obtener lo que pide el sql
	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	}
	return user
}

// actualizar registros
func (user *User) update() {
	sql := `update users set username=$1,password=$2,email=$3 where id=$4`
	db.Exec(sql, user.Username, user.Password, user.Email, user.Id)
}

// guardar o editar registro
func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}

// eliminar un registro
func (user *User) Delete() {
	sql := `delete from users where id=$1`
	db.Exec(sql, user.Id)
}

/*sqlQuery := `insert into usuario(nombreusuario,contra,idpersona,
	idtiporol,habilitado) values($1,$2,$3,$4,true)`
_, err := tx.Exec(sqlQuery, ousuario.Nombreusuario, claveCifrada,
	ousuario.Idpersona, ousuario.Idtiporol)*/
