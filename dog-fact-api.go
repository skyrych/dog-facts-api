package main

import (
	"fmt"
	"net/http"
)

func facts(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "This is the page of random facts from life of Needy Trussardi")
}

func main() {
	http.HandleFunc("/facts", facts)
	http.ListenAndServe(":8080", nil)
}
