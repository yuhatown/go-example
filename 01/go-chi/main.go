package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}

func main() {
	mux := chi.NewRouter()
	mux.Get("/hello", helloHandler)
	http.ListenAndServe(":8080", mux)

	db, err := sql.Open("mysql", "test:test1234@localhost/test_go")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	

	var answer string
	db.QueryRow("INSERT INTO test (value) VALUES (1+1)")
	db.QueryRow("SELECT value FROM test").Scan(&answer)
	fmt.Println(answer)
}
 