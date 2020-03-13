package main

import (
	"log"
	"net/http"
)

func main() {
	// "Signin" and "Welcome" are the handlers that we will implement
	http.HandleFunc("/signin", signInHandler)
	http.HandleFunc("/validate", validateToken)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}
