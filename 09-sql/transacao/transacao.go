package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:199718@/cursogo")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// db.Exec("DELETE FROM usuarios")
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO usuarios(id, nome) VALUES(?, ?)")
	stmt.Exec(1, "Rômulo")
	stmt.Exec(2, "Rodrigo")
	_, err = stmt.Exec(1, "Rogério") // ? Gera um erro de chave duplicada.

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
