package main

import (
	"fmt"
	"text/template"
	"net/http"
	//"io"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"database/sql"
	//"strings"
	"strconv"
)

type PageTags struct {
	Title string	`json:"title"`
	RecordCount int	`json:"recordcount"`
	FirstName	string	`json:"firstname"`
	LastName	string	`json:"lastname"`
	Phone		string	`json:"phone"`
	DebugData	string
	ErrorMsg	string
	ProxyURL	string
}

var (
	dbFirstName string
	dbLastName string
	dbPhone string
	resultCount int
	dbResults string
	pageDebugData string
	dbUsername string
	dbPassword string
	dbLoginString string
	dbQuery string
	proxyURL string
)

func setVars() (int) {
	dbUsername="goservices"
	dbPassword="C7163mwx!"
	dbLoginString=dbUsername+":"+dbPassword
	proxyURL="address-book"				//URL used by nginx to proxy
	return 8890
}

func errorPage(w http.ResponseWriter, r *http.Request) {
	tpl:=template.New("error.tpl")
	tpl,err:=tpl.ParseFiles("templates/error.tpl")
	if err!=nil { log.Fatalln(err.Error()) }
	err = tpl.Execute(w,PageTags{Title:"Whoops",ErrorMsg:"That page wasn't found",})
	if err!=nil {
		log.Fatalln(err)
	}
}

func pageHandler(w http.ResponseWriter,r *http.Request, params  httprouter.Params) {

	pageDebugData=""

	if r.Method=="POST" {
		pageDebugData=r.FormValue("cmd")
		switch r.FormValue("cmd") {
		case "filter":
			dbQuery = "select firstname, lastname, phone from records where LOWER(firstname)=LOWER('" + r.FormValue("filter") + "') or LOWER(lastname)=LOWER('" + r.FormValue("filter") + "') or phone='" + r.FormValue("filter") + "'"
			pageDebugData+="<br>" + dbQuery
		case "add":
			dbQuery = "insert into records values('" + r.FormValue("firstname") + "','" + r.FormValue("lastname") + "','" + r.FormValue("phone") + "')"
			pageDebugData+="<br>" + dbQuery
		}
	} else {
		dbQuery = "select firstname, lastname, phone from records"
		pageDebugData+="<br>" + dbQuery
	}

	db, err := sql.Open("mysql", dbLoginString+"@/address_book")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err=db.Ping()
	if err != nil {
		panic(err.Error())
	}


	//Load the main template
	tpl:=template.New("main.tpl")
	tpl=tpl.Funcs(template.FuncMap{
		"populateData": func() string {
			rows,err := db.Query(dbQuery)
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
	err = tpl.Execute(w,PageTags{Title:"Address Book",RecordCount:resultCount,DebugData:pageDebugData,ProxyURL:proxyURL,})
	if err!=nil {
		log.Fatalln(err)
	}
}

func main() {
	port:=strconv.Itoa(setVars())
	router:=httprouter.New()
	router.NotFound=http.HandlerFunc(errorPage)

	router.GET("/",pageHandler)
	router.POST("/",pageHandler)
	fmt.Println("Listening and ready on port: " +port)
	http.ListenAndServe(":"+port,router)
}
