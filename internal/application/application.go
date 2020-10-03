package application

import (
	"fmt"
	"os"

	"github.com/sanjushahgupta/beginner/internal/storage"
)

// CreateUser returns error if Profile is already created
func CreateUser(p storage.Profile) (string, error) {
	fn := fmt.Sprintf("%s_%s.json", p.FirstName, p.LastName)
	_, err := os.Stat(fn)
	if err == nil {
		return fn, fmt.Errorf("%s exists", fn)
	}
	if !os.IsNotExist(err) {
		return fn, err
	}

	return fn, storage.CreateUser(fn, p)
}
