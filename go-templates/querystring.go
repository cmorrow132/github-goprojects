package main

import (
	//"fmt"
	"net/http"
	"io"
)

func pageHandler(w http.ResponseWriter,r *http.Request) {
	key:="q"
	val:=r.URL.Query().Get(key)
	io.WriteString(w,"You searched for " + val)
}

func main() {
	http.HandleFunc("/",pageHandler)
	http.ListenAndServe(":8080",nil)
}
