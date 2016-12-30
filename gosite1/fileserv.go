package main

import (
	"net/http"
	"io/ioutil"
	//"strconv"
	"log"
)

type MyHandler struct {
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path:=r.URL.Path[1:]
	log.Println(path)

	data,err:=ioutil.ReadFile(string(path))

	if err==nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 " + http.StatusText(404)))
	}
}

func main() {
	http.Handle("/",new(MyHandler))
	http.ListenAndServe(":8080",nil)
}

