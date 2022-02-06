package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/jperarm1971/convenios-golang/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "josue"
)

func main() {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()
	sqlStatement := `select row_to_json(row) from (SELECT * FROM public.prueba WHERE id=$1) row`
	row := db.QueryRow(sqlStatement, 1)
	CheckError(err)
	var resultado models.Usuario
	row.Scan(&resultado)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
