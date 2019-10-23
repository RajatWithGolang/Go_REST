package main

import (
	 "github.com/julienschmidt/httprouter"
	 "net/http"
	 "fmt"
	 "log"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", p.ByName("user"))
}
func main(){
	router := httprouter.New();
	router.GET("/",Index)
	router.GET("/hello/:user",Hello)
	log.Fatalln(http.ListenAndServe(":8080",router))
}