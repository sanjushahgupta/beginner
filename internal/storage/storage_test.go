package storage

import (
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
}
