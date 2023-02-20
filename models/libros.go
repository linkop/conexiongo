package models

type libros struct {
	codigo    int
	titulo    string
	autor     string
	editorial string
}

const LibroSchema string = `create table libros(
	codigo serial,
	titulo varchar(20),
	autor varchar(30),
	editorial varchar(15),
	primary key (codigo))`
