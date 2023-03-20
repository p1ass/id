package internal

import (
	"errors"
	"fmt"
	"net/url"
	"sync"
)

var ErrCodeNotFound = errors.New("code not found")

type CodeDatastore interface {
	// Fetch fetches AuthorizationCode witch is not expired.
	// When code is expired, returns ErrCodeIsExpired error.
	Fetch(code, clientID string, redirectURI url.URL) (*AuthorizationCode, error)
	Save(code *AuthorizationCode) error
}

type inMemoryCodeDatastore struct {
	datastore sync.Map
}

func NewInMemoryCodeDatastore() CodeDatastore {
	return &inMemoryCodeDatastore{
		datastore: sync.Map{},
	}
}

func (d *inMemoryCodeDatastore) Fetch(code, clientID string, redirectURI url.URL) (*AuthorizationCode, error) {
	value, ok := d.datastore.Load(d.key(code, clientID, redirectURI))
	if !ok {
		return nil, ErrCodeNotFound
	}
	gotCode := value.(*AuthorizationCode)
	if gotCode.Expired() {
		return nil, ErrCodeIsExpired
	}
	return gotCode, nil
}

func (d *inMemoryCodeDatastore) Save(code *AuthorizationCode) error {
	d.datastore.Store(d.key(code.Code, code.clientID, code.redirectURI), code)
	return nil
}

// key for only in-memory datastore.
func (d *inMemoryCodeDatastore) key(code, clientID string, redirectURI url.URL) string {
	return fmt.Sprintf("%s-%s-%s", clientID, code, redirectURI.String())
}
