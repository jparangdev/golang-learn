package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func BarHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values.Get("name")
	if name == "" {
		name = "world"
	}
	id, _ := strconv.Atoi(values.Get("id"))
	fmt.Fprintf(w, "Hello %s, your ID is %d", name, id)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Jparangdev")
	})
	mux.HandleFunc("/bar", BarHandler)

	http.ListenAndServe(":3000", mux)
}
