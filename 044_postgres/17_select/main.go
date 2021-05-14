package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Book struct {
	//NOTE: the order and the names of theses fields must match what you have in the database. This is IMPORTANT
	isbn   string
	title  string
	author string
	price  float32
}

func main() {
	un := os.Getenv("POSTGRES_USERNAME") //FYI: this user must have the "SUPERUSER" role in postgress to do this operations https://www.postgresql.org/docs/12/sql-alteruser.html
	pwd := os.Getenv("POSTGRES_PASSWORD")
	h := os.Getenv("POSTGRES_HOST")
	pdb := os.Getenv("POSTGRES_DATABASE")

	db, err := sql.Open("postgres", "postgres://"+un+":"+pwd+"@"+h+"/"+pdb+"?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close() //we defer close that pointer to db

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")

	//We created books table and populated it with a few books already using postbird
	rows, err := db.Query("SELECT * FROM books;") //NOTE: the semi-colon must be at the end of the SQL stement
	if err != nil {
		panic(err)
	}
	defer rows.Close() //we defer close that pointer to rows

	bks := make([]Book, 0) //the 0 means we length of the slice. it is not a fixed length
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
		if err != nil {
			panic(err)
		}
		bks = append(bks, bk) //we add the book to our slice of book
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	//we range over the slice
	for _, bk := range bks {
		// fmt.Println(bk.isbn, bk.title, bk.author, bk.price)
		fmt.Printf("%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}
