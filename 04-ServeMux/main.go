package main

import (
	"log"
	"net/http"
	"io"
)

func sayHello(w http.ResponseWriter,r *http.Request){
         io.WriteString(w,"Hello from Servemux!")
}
func sayBye(w http.ResponseWriter,r *http.Request){
         io.WriteString(w,"Bye from Servemux!")
}
func main(){
		mux := http.NewServeMux()
	    mux.HandleFunc("/hello",sayHello)
	    mux.HandleFunc("/bye",sayBye)
	   log.Fatalln(http.ListenAndServe(":8080",mux))
}