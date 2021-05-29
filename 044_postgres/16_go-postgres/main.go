package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" //We have to import the driver as usual. NOTE that the blank identifier is a way to tell the identifier that we are importing the package but we are not using it to avoid syntax or compile error
)

func main() {
	/*
		Good developer practise is to make us of environmetal varibale to avoid comming your secrets to gitHub
		I stored all my variables at my .bash_profile. open the file using
		- VSCode running: 'code ~/.bash_profile'

		Then i used 'source ~/.bash_profile' to make use of the variables within our current terminal
		To confirm the varibles are being used type 'echo $<WHATEVER_ENV_NAME>'

		Each time you update env variables at the bash file, you have to run 'source ~/.bash_profile' again to update the variables in terminal
	*/
	un := os.Getenv("POSTGRES_USERNAME")
	pwd := os.Getenv("POSTGRES_PASSWORD")
	h := os.Getenv("POSTGRES_HOST")
	pdb := os.Getenv("POSTGRES_DATABASE") //we called our database "bookstore"

	//db, err := sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable") //i will use this line if i am connecting to a localHost postgres. The host will be localhost
	db, err := sql.Open("postgres", "postgres://"+un+":"+pwd+"@"+h+"/"+pdb+"?sslmode=disable")
	if err != nil {
		panic(err) //panic will exit the program. we are basically saying if this connection dont work, there is no need to continue; just exit program
	}
	defer db.Close()

	//Pinging your db is not really necessary but if you have a Production code and you want to stop the program if there is no connection, you can do this
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}
