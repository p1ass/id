package internal

import (
	"encoding/base64"
	"errors"

	"github.com/p1ass/id/backend/pkg/randgenerator"
	"golang.org/x/crypto/argon2"
)

type (
	// RawPassword is a non hashed password.
	// RawPassword implements [fmt.Stringer] and [fmt.GoStringer], so raw password is not exposed.
	RawPassword string

	// HashedPassword is a base64 raw url encoded hashed password using [Argon2id] algorithm.
	// [Argon2id] is the winner of the 2015 Password Hashing Competition and
	// is recommended by OWASP [Password Storage Cheat Sheet].
	//
	// [Argon2id]: https://en.wikipedia.org/wiki/Argon2
	// [Password Storage Cheat Sheet]: https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html
	HashedPassword struct {
		hashedPassword string
		salt           string
	}
)

// According to OWASP [Password Storage Cheat Sheet], argon2id parameters should use following parameters.
//
// [Password Storage Cheat Sheet]: https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#password-hashing-algorithms
const (
	argon2idMemoryKB  = 37 * 1024
	argon2idTime      = 1
	argon2idThreads   = 1
	argon2idKeLenByte = 32
	argon2SaltByte    = 16
)

// ErrMismatchedHashAndPassword is returned from ComparePassword when a password and hash do not match.
var ErrMismatchedHashAndPassword = errors.New("password is not the hash of the given password")

// NewHashedPassword generates hashed password and salt.
func NewHashedPassword(rawPassword RawPassword) *HashedPassword {
	salt := randgenerator.MustGenerateToString(argon2SaltByte)
	hashedPassword := base64.RawURLEncoding.EncodeToString(argon2.IDKey(
		[]byte(rawPassword),
		[]byte(salt),
		argon2idTime,
		argon2idMemoryKB,
		argon2idThreads,
		argon2idKeLenByte),
	)
	return &HashedPassword{
		hashedPassword: hashedPassword,
		salt:           salt,
	}
}

const masked = "[masked]"

func (p RawPassword) String() string {
	return masked
}

func (p RawPassword) GoString() string {
	return masked
}

func (p HashedPassword) String() string {
	return masked
}

func (p HashedPassword) GoString() string {
	return masked
}

// ComparePassword compares the given raw password with it.
// It returns nil on success, or an error on failure.
func (p HashedPassword) ComparePassword(other RawPassword) error {
	otherHash := base64.RawURLEncoding.EncodeToString(argon2.IDKey(
		[]byte(other),
		[]byte(p.salt),
		argon2idTime,
		argon2idMemoryKB,
		argon2idThreads,
		argon2idKeLenByte),
	)
	if p.hashedPassword != otherHash {
		return ErrMismatchedHashAndPassword
	}
	return nil
}
