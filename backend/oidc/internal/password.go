package internal

import (
	"crypto/rand"
	"encoding/base64"
	"io"

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
	HashedPassword string
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

// NewHashedPassword generates hashed password and salt.
func NewHashedPassword(rawPassword RawPassword) (hashedPassword string, salt string) {
	salt = mustGenerateSalt(argon2SaltByte)
	hashedPassword = base64.RawURLEncoding.EncodeToString(argon2.IDKey(
		[]byte(rawPassword),
		[]byte(salt),
		argon2idTime,
		argon2idMemoryKB,
		argon2idThreads,
		argon2idKeLenByte),
	)
	return hashedPassword, salt
}

func (p RawPassword) String() string {
	return "[masked]"
}

func (p RawPassword) GoString() string {
	return "[masked]"
}

func (p HashedPassword) String() string {
	return "[masked]"
}

func (p HashedPassword) GoString() string {
	return "[masked]"
}

func mustGenerateSalt(saltByte uint16) string {
	unencodedSalt := make([]byte, saltByte)
	_, err := io.ReadFull(rand.Reader, unencodedSalt)
	if err != nil {
		panic(err)
	}

	return base64.RawURLEncoding.EncodeToString(unencodedSalt)
}
