package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RajatWithGolang/12-GO_REST/09-TaskAPI/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TaskController struct {
	session *mgo.Session
}

func NewTaskController(s *mgo.Session) *TaskController {
	return &TaskController{s}
}

func (tc *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	task := models.Task{}
	json.NewDecoder(r.Body).Decode(&task)
	task.Id = bson.NewObjectId()
	if err := tc.session.DB("TaskDB").C("tasks").Insert(task); err != nil {
		w.WriteHeader(404)
		return
	}
	tj, err := json.Marshal(task)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", tj)

}
func (tc *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {
	task := models.Task{}
	err := tc.session.DB("TaskDB").C("tasks").Find(nil).All(&task)
	if err != nil {
		w.WriteHeader(404)
		fmt.Println("Results All: ", task)
		return
	}
	tj, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", tj)
}

func (tc *TaskController) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := bson.ObjectIdHex(params["id"])
	task := models.Task{}
	err := tc.session.DB("TaskDB").C("tasks").FindId(id).One(&task)
	if err != nil {
		w.WriteHeader(404)
		fmt.Println("Results All: ", task)
		return
	}
	tj, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", tj)
}

func (tc *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := bson.ObjectIdHex(params["id"])
	task := models.Task{}
	err := tc.session.DB("TaskDB").C("tasks").Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"description": "New Description"}})
	if err != nil {
		fmt.Println("error during update: ", err)
		return
	}
	tj, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", tj)
}

func (tc *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := bson.ObjectIdHex(params["id"])
	err := tc.session.DB("TaskDB").C("tasks").RemoveId(id)
	if err != nil {
		fmt.Println("error during update: ", err)
		return
	}
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted task", id, "\n")
}
