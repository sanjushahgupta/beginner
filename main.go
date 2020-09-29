package main

import (
	"fmt"
	"net/http"

	b "github.com/sanjushahgupta/beginner/basic"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
		b.Cal(3, 6)
	})

	http.ListenAndServe(":8080", nil)

}
