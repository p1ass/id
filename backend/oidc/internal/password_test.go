package internal_test

import (
	"testing"

	"github.com/p1ass/id/backend/oidc/internal"
)

func TestNewHashedPasswordGeneratesDifferentValueEvenIfRawInputIsSame(t *testing.T) {
	t.Parallel()

	rawPassword := internal.RawPassword("rawPassword1234")

	hashed1, salt1 := internal.NewHashedPassword(rawPassword)
	hashed2, salt2 := internal.NewHashedPassword(rawPassword)

	if hashed1 == hashed2 {
		t.Errorf("NewHashedPassword(%v) hashed output is same %v", rawPassword, hashed1)
	}

	if salt1 == salt2 {
		t.Errorf("NewHashedPassword(%v) salt is same %v", rawPassword, salt1)
	}

}
