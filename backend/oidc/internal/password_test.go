package internal_test

import (
	"reflect"
	"testing"

	"github.com/p1ass/id/backend/oidc/internal"
)

func TestNewHashedPasswordGeneratesDifferentValueEvenIfRawInputIsSame(t *testing.T) {
	t.Parallel()

	rawPassword := internal.RawPassword("rawPassword1234")

	hashed1 := internal.NewHashedPassword(rawPassword)
	hashed2 := internal.NewHashedPassword(rawPassword)

	if reflect.DeepEqual(hashed1, hashed2) {
		t.Errorf("NewHashedPassword(%v) hashed output is same %v", rawPassword, hashed1)
	}
}

func TestHashedPassword_ComparePassword(t *testing.T) {
	t.Parallel()

	type args struct {
		other internal.RawPassword
	}

	tests := []struct {
		name           string
		hashedPassword *internal.HashedPassword
		args           args
		wantErr        bool
	}{
		{
			name:           "should return nil when password is match",
			hashedPassword: internal.NewHashedPassword("rawPassword1"),
			args: args{
				other: "rawPassword1",
			},
			wantErr: false,
		},
		{
			name:           "should not return error when password is not match",
			hashedPassword: internal.NewHashedPassword("rawPassword1"),
			args: args{
				other: "rawPassword2",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if err := tt.hashedPassword.ComparePassword(tt.args.other); (err != nil) != tt.wantErr {
				t.Errorf("ComparePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
