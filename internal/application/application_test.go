package application

import (
	"testing"

	"github.com/sanjushahgupta/beginner/internal/storage"
)

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name    string
		file    string
		profile storage.Profile
	}{
		{
			name: "asd_qwe",
			profile: storage.Profile{
				FirstName: "asd",
				LastName:  "qwe",
				Email:     "asd@qwe.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn, err := CreateUser(tt.profile)
			if err != nil {
				t.Error(err)
				return
			}
			if fn != tt.file {
				t.Errorf("wanted %s, got %s", tt.file, fn)
			}
		})
	}
}
