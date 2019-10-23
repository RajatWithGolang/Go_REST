package main

import (
	"log"
	"net/http"
	"io"
)

func sayHello(w http.ResponseWriter,r *http.Request){
         io.WriteString(w,"Hello World!")
}
func main(){
	   http.HandleFunc("/hello",sayHello)
	   log.Fatalln(http.ListenAndServe(":8080",nil))
}