package application

import (
	"fmt"
	"os"

	"github.com/sanjushahgupta/beginner/internal/storage"
)

func CreateUser(p storage.Profile) (string, error) {
	fn := fmt.Sprintf("%s_%s.json", p.FirstName, p.LastName)
	_, err := os.Stat(fn)
	if err == nil {
		return fn, fmt.Errorf("%s exists", fn)
	}
	if !os.IsNotExist(err) {
		return fn, err
	}

	return storage.CreateUser(fn, p)
}
