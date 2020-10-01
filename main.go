package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/sanjushahgupta/beginner/basic"
)

func show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/:x/:operator/:y", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		x, err := strconv.Atoi(ps.ByName("x"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		y, err := strconv.Atoi(ps.ByName("y"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		switch op := ps.ByName("operator"); op {
		case "plus":
			w.Write([]byte(fmt.Sprintf("%d + %d = %d", x, y, basic.Add(x, y))))
		case "minus":
			w.Write([]byte(fmt.Sprintf("%d - %d = %d", x, y, basic.Subtract(x, y))))
		case "multiply":
			w.Write([]byte(fmt.Sprintf("%d * %d = %d", x, y, basic.Multiply(x, y))))
		case "divide":
			w.Write([]byte(fmt.Sprintf("%d / %d = %d", x, y, basic.Divide(x, y))))
		default:
			w.Write([]byte(fmt.Sprintf("error: operator %s not implemented\navailable operators:\n- plus\n- minus\n- multiply\n- divide", op)))
		}
	})

	http.ListenAndServe(":8080", router)
}
