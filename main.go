package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"net/http"

	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "josue"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "Has requerido: %s en la página %s\n", title, page)

		psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

		db, err := sql.Open("postgres", psqlconn)
		CheckError(err)

		defer db.Close()

		// insert
		// hardcoded
		insertStmt := `insert into prueba(nombre) values( 'Josue4')`
		_, e := db.Exec(insertStmt)
		CheckError(e)

		// dynamic
		insertDynStmt := `insert into prueba(nombre) values($1)`
		_, e = db.Exec(insertDynStmt, "Jane5")
		CheckError(e)

		fmt.Fprintf(w, "Y has insertado dos nuevos registros sobre la BBDD")
	})

	http.ListenAndServe(":80", r)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
