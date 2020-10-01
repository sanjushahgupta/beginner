package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/sanjushahgupta/beginner/basic"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to calculator!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/:x/:operator/:y", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		x, _ := strconv.Atoi(ps.ByName("x"))

		y, _ := strconv.Atoi(ps.ByName("y"))
		switch op := ps.ByName("operator"); op {
		case "plus":
			w.Write([]byte(fmt.Sprintf("%d + %d = %d", x, y, basic.Add(x, y))))
		case "minus":
			w.Write([]byte(fmt.Sprintf("%d - %d = %d", x, y, basic.Sub(x, y))))
		case "multiply":
			w.Write([]byte(fmt.Sprintf(" %d * %d = %d", x, y, basic.Mul(x, y))))
		case "divide":
			w.Write([]byte(fmt.Sprintf("%d / %d = %d", x, y, basic.Div(x, y))))
		default:
			w.Write([]byte(fmt.Sprintf("error: please use %s operators:\n- plus\n- minus\n- multiply\n- divide", op)))
		}

	})
	router.GET("/", Index)
	http.ListenAndServe(":8080", router)
}
