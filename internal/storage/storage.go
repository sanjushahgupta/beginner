package storage

import (
	"encoding/json"
	"io/ioutil"
)

type Profile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func CreateUser(fn string, p Profile) (string, error) {
	b, _ := json.MarshalIndent(p, "", "  ")
	return fn, ioutil.WriteFile(fn, b, 0644)
}
