package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Profile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func CreateUser(p Profile) (string, error) {
	fn := fmt.Sprintf("%s_%s.json", p.FirstName, p.LastName)
	_, err := os.Stat(fn)
	if err == nil {
		return fn, fmt.Errorf("%s exists", fn)
	}
	if !os.IsNotExist(err) {
		return fn, err
	}

	b, _ := json.MarshalIndent(p, "", "  ")
	return fn, ioutil.WriteFile(fn, b, 0644)
}
