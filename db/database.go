package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//const url = "postgres:1234@tcp(localhost:5432)/postgres"

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "goweb_db"
)

// guarda la conexion
var db *sql.DB

// realiza la conexion
func Connect() {
	url := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	conection, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion exitosa")
	db = conection

}

// cerrar la conexion
func Close() {
	db.Close()
}

// verificar la conexion
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

// crea una tabla
func CreateTable(schema string) {
	db.Exec(schema)
}

// borrar los datos de una tabla
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE TABLE %s", tableName)
	Exec(sql)
}

// polimorfismo de exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

// polimorfismo query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}
