package main

import (
	"io/ioutil"
	"net/http"
	//"strconv"
	"log"
)

type MyHandler struct {
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Println(path)

	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		w.Write(data)
	} else {
		data, err := ioutil.ReadFile(string("/error/404.html"))
		if err==nil {
			w.Write(data)
		} else {
			w.Write([]byte("Error page not found"))
		}
		
	}
}

func main() {
	http.Handle("/", new(MyHandler))
	http.ListenAndServe(":8080", nil)
}
