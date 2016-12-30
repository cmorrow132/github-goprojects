package main

import (
	//"fmt"
	"text/template"
	"net/http"
	"io"
	"log"
)

type PageTags struct {
	Title string	`json:"title"`
	RecordCount int	`json:"recordcount"`
}

type Person struct {
	FirstName	string	`json:"firstname"`
	LastName	string	`json:"lastname"`
	Phone		string	`json:"phone"`
}

func pageHandler(w http.ResponseWriter,r *http.Request) {
	//First check the path
	rURL:=r.URL.Path[1:]	//Strip the leading slash

	if rURL=="" {
		//Load the main template
		t,err:=template.ParseFiles("templates/main.tpl")
		if err!=nil { log.Fatalln(err.Error()) }
		t.Execute(w,PageTags{Title:"Address Book",})

	}

	if r.Method=="GET" {
		key:="q"
		val:=r.URL.Query().Get(key)
		io.WriteString(w,"<br><br>")
		io.WriteString(w,"You sent a GET request<br>")
		io.WriteString(w,"You searched for '" + val + "'")
	} else {
		io.WriteString(w,"You are posting data\n")
	}
}

func main() {
	http.HandleFunc("/",pageHandler)
	http.ListenAndServe(":8080",nil)
}
