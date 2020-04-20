package main

import (
	"fmt"
	"log"
	"net/http"
)

// http://localhost:8282

func HomePostRouterHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("username")
	age := r.FormValue("userage")
	fmt.Fprintf(w, "Server is listening...\n")
	fmt.Fprintf(w, "Имя: %s Возраст: %s", name, age)
}

func FileUserHtml(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "user.html")
}

func main() {
	http.HandleFunc("/", FileUserHtml)
	http.HandleFunc("/postform", HomePostRouterHandler)

	err := http.ListenAndServe(":8282", nil)
	if err != nil {
		log.Fatal("POST: ListenAndServe: ", err)
	}
}
