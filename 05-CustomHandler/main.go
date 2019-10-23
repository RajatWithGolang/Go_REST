package main

import (
	"log"
	"net/http"
	"io"
)
type customMux struct{
	message string
}

func (c *customMux) ServeHTTP(w http.ResponseWriter,r *http.Request){
         io.WriteString(w,c.message)
}

func main(){
	    mux := http.NewServeMux()
		c1 :=    &customMux{"Hello From Custom Handler"}
		c2 :=    &customMux{"Bye From Custom Handler"}
	    mux.Handle("/hello",c1)
	    mux.Handle("/bye",c2)
	    log.Fatalln(http.ListenAndServe(":8080",mux))
}