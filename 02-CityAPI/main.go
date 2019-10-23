package main

import(
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)
type City struct{
	Name string
	Area int
}
func mainLogic(w http.ResponseWriter,r *http.Request){
	if r.Method == "POST"{
		var city City
		decoder := json.NewDecoder(r.Body)
		err:= decoder.Decode(&city)
		if err != nil{
            panic(err)
		}
		defer r.Body.Close()
		fmt.Printf("Got %s city with area of %d sq miles!\n",city.Name, city.Area)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - Created"))
	}else{
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
	}
}



func main(){
	   http.HandleFunc("/city",mainLogic)
	   log.Fatalln(http.ListenAndServe(":8080",nil))
}