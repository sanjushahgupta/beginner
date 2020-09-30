package main

import (
	"fmt"
	"net/http"

	"html/template"

	b "github.com/sanjushahgupta/beginner/basic"
)

func show(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/show", show)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
		b.Cal(3, 6)
		//	b.add(2, 5)
	})

	http.ListenAndServe(":8080", nil)

}
