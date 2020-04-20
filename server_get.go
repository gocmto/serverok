package main

import (
	"fmt"
	"log"
	"net/http"
)

// http://localhost:8181?name=Sam&age=21

func HomeGetRouterHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	fmt.Fprintf(w, "Server is listening...\n")
	fmt.Fprintf(w, "Имя: %s Возраст: %s", name, age)
}

func main() {
	http.HandleFunc("/", HomeGetRouterHandler)

	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		log.Fatal("GET: ListenAndServe: ", err)
	}
}
