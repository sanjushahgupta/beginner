package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type profile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()
	b, _ := ioutil.ReadAll(r.Body)

	var p profile
	json.Unmarshal(b, &p)

	fn := fmt.Sprintf("%s_%s.json", p.FirstName, p.LastName)
	_, err := os.Stat(fn)
	if err == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%s.json exists", fn)))
		return
	}
	if !os.IsNotExist(err) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	b, _ = json.MarshalIndent(p, "", "  ")
	err = ioutil.WriteFile(fn, b, 0644)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(fmt.Sprintf("%s.json created", fn)))
}
