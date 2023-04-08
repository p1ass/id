package internal

import (
	"errors"
	"sync"
)

var ErrAccessTokenNotFound = errors.New("access token not found")

type AccessTokenDatastore interface {
	// Fetch fetches AccessToken.
	// If access token is not found, return ErrAccessTokenNotFound error.
	Fetch(token string) (*AccessToken, error)
	// Save saves a access token to datastore.
	Save(accessToken *AccessToken) error
}

type inMemoryAccessTokenDatastore struct {
	datastore sync.Map
}

func NewInMemoryAccessTokenDatastore() AccessTokenDatastore {
	return &inMemoryAccessTokenDatastore{
		datastore: sync.Map{},
	}
}

func (d *inMemoryAccessTokenDatastore) Fetch(token string) (*AccessToken, error) {
	value, ok := d.datastore.Load(token)
	if !ok {
		return nil, ErrAccessTokenNotFound
	}
	gotAccessToken, ok := value.(*AccessToken)
	if !ok {
		panic("type assertion error")
	}

	return gotAccessToken, nil
}

func (d *inMemoryAccessTokenDatastore) Save(accessToken *AccessToken) error {
	d.datastore.Store(accessToken.Token, accessToken)
	return nil
}
