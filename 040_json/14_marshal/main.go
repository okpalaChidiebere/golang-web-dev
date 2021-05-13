package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type model struct {
	//Private fields for this struct. Private fields cannot be exported to JSON. We know that are private because they all starts with lowercases
	state    bool
	pictures []string
}

// What went wrong?
// See the next file for the answer.

func main() {
	m := model{
		state: true,
		pictures: []string{
			"one.jpg",
			"two.jpg",
			"three.jpg",
		},
	}

	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}

	os.Stdout.Write(bs) //You get an empty JSON {} printed out because your struct fields are lower case. So they cannot be exported to JSON
}
