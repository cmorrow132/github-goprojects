package main

import (
	"fmt"
	"net/http"
)

func pageHandler(w http.ResponseWriter,r *http.Request) {
	fmt.Println(r.RequestURI)
}

func main() {
	http.HandleFunc("/",pageHandler)
	http.ListenAndServe(":8080",nil)
}
