package main

import (
	"fmt"
	"net/http"
)

type hotdog int

/* A Handler interface interface is any type that implements the method with the signature

ServeHTTP(ResponseWriter, *Request). So in our casem hotdog type is a Handler!

hotdog implicitly implements the Handler interface

GoLang is about types!
*/
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Any code you want in this func")
}

func main() {

}
