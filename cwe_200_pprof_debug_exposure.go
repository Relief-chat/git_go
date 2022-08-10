package main

import (
	"fmt"
	"log"
	"net/http"

	_ "net/http/pprof"
)

func ok() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
    // ok 
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func ok2() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
    // ok 
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func ok3() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	mux := http.NewServeMux()
    // ok 
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
    // bad
	log.Fatal(http.ListenAndServe(":8080", nil))
}
