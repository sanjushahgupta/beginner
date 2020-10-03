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

// CreateUser writes Profile to given filename
func CreateUser(fn string, p Profile) error {
	b, _ := json.MarshalIndent(p, "", "  ")
	return ioutil.WriteFile(fn, b, 0644)
}
