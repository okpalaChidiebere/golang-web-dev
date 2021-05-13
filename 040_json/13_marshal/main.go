package main

import (
	"encoding/json"
	"fmt"
)

/*
This code shows an example of what happens when you marshall an empty struct
*/

type model struct {
	State    bool
	Pictures []string
}

func main() {
	m := model{}

	fmt.Println(m) //we get the zero values printed out like {false []}

	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(string(bs)) //we get {"State": false, "Pictures: null"} as JSON
}
