// A simple echo server.
package main

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!\n"))
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":2929", nil))
}
