package main

import (
	//"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	//"encoding/json"
)



func main() {
	type Person struct {
		Id		int
		First_Name	string	`json:"first_name"`
		Last_Name	string	`json:"last_name"`
	}

	db,err:=sql.Open("mysql","goservices:password@tcp(127.0.0.1:3306)/restap1")

	if err!=nil {
		fmt.Print(err.Error())
	}

	defer db.Close()

	err=db.Ping()
	if err!=nil {
		fmt.Print(err.Error())
	}

	router:=gin.Default()

	router.GET("/person/:id", func(c *gin.Context) {
		var (
			person Person
			result gin.H
		)
		id:=c.Param("id")
		row:=db.QueryRow("select id, first_name, last_name from person where id = ?;",id)
		err=row.Scan(&person.Id,&person.First_Name,&person.Last_Name)

		if err!=nil {
			result=gin.H{
				"result": nil,
				"count": 0,
			}
		} else {
			result=gin.H{
				"result": person,
				"count": 1,
			}
		}

		c.JSON(http.StatusOK,result)
	})

	router.GET("/persons", func(c *gin.Context) {
		var (
			person Person
			persons []Person
		)

		rows,err:=db.Query("select id, first_name, last_name from person;")
		if err!=nil {
			fmt.Print(err.Error())
		}

		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": persons,
			"count": len(persons),
		})
	})

	router.Run(":3000")
}
