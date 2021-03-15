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

	stmt, _ := db.Prepare("UPDATE usuarios SET nome = ? WHERE id = ?")
	stmt2, _ := db.Prepare("DELETE FROM usuarios WHERE id = ?")
	stmt.Exec("Ana Belly", 3)
	stmt2.Exec(3)
}
