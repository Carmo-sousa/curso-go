package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Abrindo conexão com o banco de dados.
	db, err := sql.Open("mysql", "root:199718@/cursogo")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.Exec("DELETE FROM usuarios")
	stmt, _ := db.Prepare("INSERT INTO usuarios(nome) VALUES(?)")
	stmt.Exec("Maria")
	stmt.Exec("João")

	res, _ := stmt.Exec("Pedro")

	id, _ := res.LastInsertId()
	rows, _ := res.RowsAffected()

	fmt.Println(id, rows)
}
