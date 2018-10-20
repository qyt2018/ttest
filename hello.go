package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, golang!", r.Host)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9001", nil)
}
