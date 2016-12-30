package main

import (
	"fmt"
	"net/http"
)

func mainWebServ(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You requested ", r.Method)
}

func main() {
	http.HandleFunc("/", mainWebServ)
	http.ListenAndServe(":8000", nil)
}
