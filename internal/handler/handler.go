package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sanjushahgupta/beginner/internal/storage"
)

type profile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()
	b, _ := ioutil.ReadAll(r.Body)

	var p storage.Profile
	json.Unmarshal(b, &p)

	fn, err := storage.CreateUser(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(fmt.Sprintf("%s.json created", fn)))
}
