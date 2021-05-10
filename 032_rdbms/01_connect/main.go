package main

import (
	"database/sql"
	"fmt"

	/*
		The reason we have the blanc identifier "_" is we are importing the driver
		But we are not really using anything from this package. This package is
		imported strictly for set up. The code from this driver package will be analyzied
		only when our mysql connection is being made for the drivers to get connected

		For our app, the driver is mysql. You can see that passed as the first parameter in sql.Open() method
	*/
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	db, err = sql.Open("mysql", "awsuser:mypassword@tcp(mydbinstance.cakwl95bxza0.ca-central-1.rds.amazonaws.com:3306)/test02?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "Successfully completed.")
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
