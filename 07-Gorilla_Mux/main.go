package main

import (
	 "github.com/gorilla/mux"
	 "net/http"
	 "fmt"
	 "log"
)

func Index(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);
	fmt.Fprintf(w, "hello, %s!\n",vars["user"])
	fmt.Fprintf(w, "ID is , %s!\n",vars["id"])
}
func main(){
	router := mux.NewRouter()
	router.HandleFunc("/",Index).Methods("GET")
	router.HandleFunc("/hello/{user}/{id}",Hello).Methods("GET")
	log.Fatalln(http.ListenAndServe(":8080",router))
}