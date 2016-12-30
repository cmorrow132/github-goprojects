package main

import(
	"fmt"
	//"strings"
	"net/http"
	"log"
)

func handlePageTest1(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprintf(w,"Test1")
}

func main() {
	http.HandleFunc("/test1",handlePageTest1)
	err:=http.ListenAndServe(":8080",nil)
	if err!= nil {
		log.Fatal("ListenAndServe:", err)
	}
}
