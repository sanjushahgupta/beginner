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

func Sum(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	x, _ := strconv.Atoi(ps.ByName("x"))

	y, _ := strconv.Atoi(ps.ByName("y"))

	plus := strconv.Itoa(basic.Add(x, y))
	w.Write([]byte(plus))
}
func Subt(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	x, _ := strconv.Atoi(ps.ByName("x"))

	y, _ := strconv.Atoi(ps.ByName("y"))
	minus := strconv.Itoa(basic.Sub(x, y))
	w.Write([]byte(minus))
}

func Multi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	x, _ := strconv.Atoi(ps.ByName("x"))

	y, _ := strconv.Atoi(ps.ByName("y"))

	multiply := strconv.Itoa(basic.Mul(x, y))

	w.Write([]byte(multiply))
}
func Divd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	x, _ := strconv.Atoi(ps.ByName("x"))

	y, _ := strconv.Atoi(ps.ByName("y"))
	divide := strconv.Itoa(basic.Div(x, y))

	w.Write([]byte(divide))
}

func main() {
	router := httprouter.New()

	router.GET("/:x/plus/:y", Sum)

	router.GET("/:x/minus/:y", Subt)
	router.GET("/:x/multiply/:y", Multi)
	router.GET("/:x/divide/:y", Divd)

	router.GET("/", Index)

	http.ListenAndServe(":8080", router)
}
