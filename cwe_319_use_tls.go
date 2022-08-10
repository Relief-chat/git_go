package main

import (
    "net/http"
    "fmt"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.write([]byte("Hello, world!"))
}

func main() {
    http.HandleFunc("/index", Handler)
    // bad
    http.ListenAndServe(":80", nil)
    // ok
    http.ListenAndServeTLS(":80", nil)
}
