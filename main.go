package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sanjushahgupta/beginner/internal/handler"
)

func main() {
	router := httprouter.New()
	router.POST("/profile", handler.CreateUser)
	http.ListenAndServe(":8080", router)
}
