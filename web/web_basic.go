package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	var student = Student{"jude", 35, 95}
	data, _ := json.Marshal(student)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(data))
}

func BarHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values.Get("name")
	if name == "" {
		name = "world"
	}
	id, _ := strconv.Atoi(values.Get("id"))
	fmt.Fprintf(w, "Hello %s, your ID is %d", name, id)
}

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/bar", BarHandler)
	mux.HandleFunc("/student", StudentHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	return mux
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
