package main

import (
	
	"net/http"
//	"github.com/gorilla/mux"
	"github.com/RajatWithGolang/12-GO_REST/09-TaskAPI/controllers"
//	"gopkg.in/mgo.v2"
	
 )

func main() {
	taskRouter  := mux.NewRouter()
	tc := controllers.NewTaskController(getSession())
	taskRouter.HandleFunc("/tasks", tc.CreateTask).Methods("POST")
	taskRouter.HandleFunc("/tasks/{id}", tc.UpdateTask).Methods("PUT")
	taskRouter.HandleFunc("/tasks", tc.GetTasks).Methods("GET")
	taskRouter.HandleFunc("/tasks/{id}", tc.GetTaskByID).Methods("GET")
	taskRouter.HandleFunc("/tasks/{id}", tc.DeleteTask).Methods("DELETE")	
	http.ListenAndServe(":8080",taskRouter)
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}