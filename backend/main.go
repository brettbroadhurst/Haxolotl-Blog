// backend/main.go - Main
// Written by Brett Broadhurst <brettbroadhurst@gmail.com>

package main

import (
	"fmt"
	"log"
	"net/http"
)

// Index route
func IndexRoute(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello, %s\n", "Brett")
}

func main() {
	http.HandleFunc("/", IndexRoute)
	log.Println("Server running on port 8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
