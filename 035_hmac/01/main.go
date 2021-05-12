package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

/*
We ran an HMAC on two strings. We noticed that we get two completely different
hashes as result even with the slightest change of the second string.
*/
func main() {
	c := getCode("test@example.com")
	fmt.Println(c)
	c = getCode("test@exampl.com")
	fmt.Println(c)
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey")) //our key will be a secret value we could store as environment variable
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
