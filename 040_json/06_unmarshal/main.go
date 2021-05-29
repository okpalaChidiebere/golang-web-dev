package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type thumbnail struct {
	URL           string
	Height, Width int
}

type img struct {
	Width, Height int
	Title         string
	Thumbnail     thumbnail
	Animated      bool
	IDs           []int
}

func main() {
	var data img
	rcvd := `{"Width":800,"Height":600,"Title":"View from 15th Floor","Thumbnail":{"Url":"http://www.example.com/image/481989943","Height":125,"Width":100},"Animated":false,"IDs":[116,943,234,38793]}` //this string is an example of a JSOn we receive from the client that cam as a result of JSON.stringify() in Javascript

	err := json.Unmarshal([]byte(rcvd), &data) //you have to pass in an address of a the value or data structure stored (a pointer) that we will unmarshal to. The value is the "data" variable. This same thing applies to json.Decode as well
	if err != nil {
		log.Fatalln("error unmarshalling", err)
	}

	fmt.Println(data)

	/*
		From here down, you see we are now accessing the data from the client with Go!
	*/

	//we range over the IDs
	for i, v := range data.IDs {
		fmt.Println(i, v)
	}

	fmt.Println(data.Thumbnail.URL)
}
