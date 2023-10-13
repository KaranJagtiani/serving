package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Docker!!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Starting server on :8081")
	http.ListenAndServe("0.0.0.0:8081", nil)
}
