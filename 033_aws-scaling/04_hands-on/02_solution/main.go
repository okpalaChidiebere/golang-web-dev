package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	//FYI: In your ideal code, username, password and hostEndpoint saved in evironment variables.
	db, err = sql.Open("mysql", "udemyokpaladev:mypassword@tcp(udemyokpaladev.cqwxwfcdjrrt.ca-central-1.rds.amazonaws.com:3306)/test02?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/instance", instance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/amigos", amigos)
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from AWS.")
}

func ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "OK")
}

func instance(w http.ResponseWriter, req *http.Request) {
	s := getInstance()
	io.WriteString(w, s)
}

func amigos(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT aName FROM amigos;`)
	check(err)

	// data to be used in query
	s := getInstance()
	s += "\nRETRIEVED RECORDS:\n"
	var name string

	// query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func getInstance() string {
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	bs := make([]byte, resp.ContentLength)
	resp.Body.Read(bs)
	resp.Body.Close()
	return string(bs)
}
