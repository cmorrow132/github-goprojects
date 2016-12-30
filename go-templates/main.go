//
package main

import (
	//"log"
	//"os"
	//"text/template"
	"net/http"
	"fmt"
)

type Page struct {
	Title string
	Body string
}

func mainPage(w http.ResponseWriter, r* http.Request) {
	fmt.Println(r.RequestURI)

	/*rURL :=r.URL.Path[1:]
	fmt.Println(rURL)

	if rURL=="" { rURL="main" }

	tpl,err:=template.ParseGlob("templates/"+rURL+".tpl")
	if err != nil { 
		fmt.Println(err)
	}
	err=tpl.Execute(w,Page{Title:"My Title",Body:"This is a test message",})*/
}

func main() {
	http.HandleFunc("/",mainPage)
	http.ListenAndServe(":8080",nil)
}
