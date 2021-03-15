package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

	return result
}

func main() {
	db, err := sql.Open("mysql", "root:199718@/")

	if err != nil {
		panic(err)
	}

	// ! Só será executado no fim da função main
	defer db.Close()

	exec(db, "CREATE DATABASE IF NOT EXISTS cursogo") // ! Cria a base de dados se ela não existir.
	exec(db, "USE cursogo")                           // ! Seleciona a base de dados (os proxímos comandos serão executados nela.)
	exec(db, "DROP TABLE IF EXISTS usuarios")         // ! Exclui a tabela de usuários se ela existir.
	exec(db, `
	CREATE TABLE usuarios (
		id integer auto_increment,
		nome VARCHAR(80),
		PRIMARY KEY (id)	
	)`) // ! Cria a tabela usuários na base de dados cursogo.
}
