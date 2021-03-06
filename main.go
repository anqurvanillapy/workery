package main

import (
	"log"
	"net/http"
	"runtime"
)

// Roughly estimate the capacity of the worker pool.
var nworker int = runtime.NumCPU() * 2

func worker(id int, jobs <-chan int, results chan<- int) {
	// TODO
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!\n"))
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":2929", nil))
}
