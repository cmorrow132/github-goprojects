package main

import (
	"fmt"
	"text/template"
	"net/http"
	//"io"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"strings"
)

type PageTags struct {
	Title string	`json:"title"`
	RecordCount int	`json:"recordcount"`
	FirstName	string	`json:"firstname"`
	LastName	string	`json:"lastname"`
	Phone		string	`json:"phone"`
	DebugData	string
}

var (
	dbFirstName string
	dbLastName string
	dbPhone string
	resultCount int
	dbResults string
	pageDebugData string
)

func pageHandler(w http.ResponseWriter,r *http.Request) {

	db, err := sql.Open("mysql", "goservices:password@/address_book")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err=db.Ping()
	if err != nil {
		panic(err.Error())
	}

	//First check the path
	rURL:=r.URL.Path[1:]	//Strip the leading slash

	if r.Method=="GET" {
		key:="q"
		val:=r.URL.Query().Get(key)
		
		pageDebugData="You sent a GET request<br>You searched for '" + key + "=" + val +"'"
	} else {
		pageDebugData="You are posting data<br>"
		r.ParseForm()
		for k,v:=range r.Form {
			pageDebugData+= k+" = "+strings.Join(v,"")+"<br>"
		}
	}

	if rURL=="" {
		//Load the main template
		tpl:=template.New("main.tpl")
		tpl=tpl.Funcs(template.FuncMap{
			"populateData": func() string {
				rows,err := db.Query("select firstname, lastname, phone from records")
				defer rows.Close()

				dbResults=""
				resultCount=0
				for rows.Next() {
					err=rows.Scan(&dbFirstName,&dbLastName,&dbPhone)
					if err!=nil {
						log.Fatalln(err)
					}
					dbResults+="<tr><td>"+dbFirstName+"</td><td>"+dbLastName+"</td><td>"+dbPhone+"</td></tr>"
					resultCount=resultCount+1
				}

				return dbResults
			},
		})

		tpl,err=tpl.ParseFiles("templates/main.tpl")
		if err!=nil { log.Fatalln(err.Error()) }
		err = tpl.Execute(w,PageTags{Title:"Address Book",RecordCount:resultCount,DebugData:pageDebugData,})
		if err!=nil {
			log.Fatalln(err)
		}
	}

}

func main() {
	http.HandleFunc("/",pageHandler)
	fmt.Println("Listening and ready on port :8080")
	http.ListenAndServe(":8080",nil)
}
