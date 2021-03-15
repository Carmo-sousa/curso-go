package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:199718@/cursogo")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, _ := db.Query("SELECT id, nome FROM usuarios WHERE id < ?", 3)
	defer rows.Close()

	for rows.Next() {
		var usr usuario
		rows.Scan(&usr.id, &usr.nome)
		fmt.Println(usr)
	}
}
