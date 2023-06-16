package main

import (
	"fmt"
	"net/http"
)

var count int

func counterHandler(w http.ResponseWriter, r *http.Request) {
	count++
	fmt.Fprintf(w, "%d", count)
}

func main() {
	http.HandleFunc("/counter", counterHandler)
	http.ListenAndServe(":8080", nil)
}
