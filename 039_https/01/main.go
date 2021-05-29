package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil) //NOTE: we server HTTPS at port 443 for prod
}

func foo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

// Go to https://localhost:10443/ or https://127.0.0.1:10443/
// list of TCP ports:
// https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers

// Generate unsigned certificate
// go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=somedomainname.com
// for example
// go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost

// WINDOWS
// windows may have issues with go env GOROOT
// go run %(go env GOROOT)%/src/crypto/tls/generate_cert.go --host=localhost

// instead of go env GOROOT
// you can just use the path to the GO SDK
// wherever it is on your computer

/*NOTE the host you want to get your key cert and private key should be from a legit certificate authority approved by web like https://letsencrypt.org/ or
you can find our names of authorized companies that offer TLS encryption that your browser approves of

To get the list of companies, go to this link chrome://settings/security
then Manage certificates, it will open a keychain Access app. In the app, go to "system Roots" and you will see the list of companies that server TLS certificates or HTTPS
The companies charge you a handful though :(

	Extra: https://caniuse.com/http2
*/
