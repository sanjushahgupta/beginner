package storage

import (
	"os"
	"testing"
)

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name    string
		file    string
		profile Profile
	}{
		{
			name: "asd_qwe",
			file: "asd_qwe.json",
			profile: Profile{
				FirstName: "asd",
				LastName:  "qwe",
				Email:     "asd@qwe.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateUser(tt.file, tt.profile)
			if err != nil {
				t.Error(err)
			}
		})
	}
	// cleanup so tests are idempotent
	for _, tt := range tests {
		err := os.Remove(tt.file)
		if err != nil {
			t.Fatal(err)
		}
	}
}
