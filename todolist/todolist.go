package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"golang-learn/todolist/model"
	"log"
	"net/http"
	"sort"
	"strconv"
)

var rd *render.Render
var todoMap map[int]model.Todo
var lastID = 0

func main() {
	rd = render.New()
	m := MakeWebHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}

}

func MakeWebHandler() http.Handler {
	todoMap = make(map[int]model.Todo)
	todoMux := mux.NewRouter()
	todoMux.Handle("/", http.FileServer(http.Dir("public")))
	todoMux.HandleFunc("/todos", GetTodoListHandler).Methods("GET")
	todoMux.HandleFunc("/todos", PostTodoHandler).Methods("POST")
	todoMux.HandleFunc("/todos/{id:[0-9]+}", RemoveTodoHandler).Methods("DELETE")
	todoMux.HandleFunc("/todos/{id:[0-9]+}", UpdateTodoHandler).Methods("PUT")
	return todoMux
}

func GetTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(model.Todos, 0)
	for _, todo := range todoMap {
		list = append(list, todo)
	}
	sort.Sort(list)
	err := rd.JSON(w, http.StatusOK, list)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func PostTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lastID++
	todo.ID = lastID
	todoMap[lastID] = todo
	err = rd.JSON(w, http.StatusCreated, todo)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func RemoveTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		err := rd.JSON(w, http.StatusOK, model.Success{true})
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		err := rd.JSON(w, http.StatusNotFound, model.Success{false})
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var updateTodo model.Todo
	err := json.NewDecoder(r.Body).Decode(&updateTodo)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if todo, ok := todoMap[id]; ok {
		todo.Name = updateTodo.Name
		todo.Completed = updateTodo.Completed
		rd.JSON(w, http.StatusOK, model.Success{true})
	} else {
		rd.JSON(w, http.StatusBadRequest, model.Success{false})
	}
}
