package internal

import (
	"errors"
	"sync"
)

var ErrClientNotFound = errors.New("client not found")

type ClientDatastore interface {
	FetchClient(id string) (*Client, error)
	SaveClient(client *Client) error
}

type inMemoryClientDatastore struct {
	datastore sync.Map
}

func NewInMemoryClientDatastore() ClientDatastore {
	return &inMemoryClientDatastore{
		datastore: sync.Map{},
	}
}

func (d *inMemoryClientDatastore) FetchClient(id string) (*Client, error) {
	value, ok := d.datastore.Load(id)
	if !ok {
		return nil, ErrClientNotFound
	}
	return value.(*Client), nil
}

func (d *inMemoryClientDatastore) SaveClient(client *Client) error {
	d.datastore.Store(client.ID, client)
	return nil
}
