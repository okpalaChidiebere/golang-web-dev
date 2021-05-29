package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB //this variable is of type pointer to a sql database. This varible is now a package level scope

func init() {
	var err error
	un := os.Getenv("POSTGRES_USERNAME")
	pwd := os.Getenv("POSTGRES_PASSWORD")
	h := os.Getenv("POSTGRES_HOST")
	pdb := os.Getenv("POSTGRES_DATABASE")
	db, err = sql.Open("postgres", "postgres://"+un+":"+pwd+"@"+h+"/"+pdb+"?sslmode=disable") //we use the = instad of :=
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

//main function launches our listen and server as normal
func main() {
	http.HandleFunc("/books", booksIndex)
	http.ListenAndServe(":8080", nil)
}

//we will do our query inside this function
func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" { //if the method is not get, we return an error
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed) //REMINDER: having th number 405 or http.StatusMethodNotAllowed are thesame!
		return
	}

	rows, err := db.Query("SELECT * FROM books") //we can access the package level variable here
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}
