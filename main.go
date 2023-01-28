package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("INFO: %+v\n", r.Header)
	fmt.Fprintf(w, "hello\n")
}

func ping(w http.ResponseWriter, r *http.Request) {
	log.Printf("INFO: %+v\n", r.Header)
	fmt.Fprintf(w, "pong\n")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/ping", ping)
	log.Printf("INFO: server start on port: 8000\n")
	http.ListenAndServe(":8000", nil)
}
