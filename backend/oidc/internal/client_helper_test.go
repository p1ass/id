package internal_test

import (
	"net/url"
	"testing"

	"github.com/p1ass/id/backend/oidc/internal"
)

// buildDummyAuthenticatedClient builds  a dummy authenticated client for testing.
func buildDummyAuthenticatedClient(t *testing.T, id string) *internal.AuthenticatedClient {
	t.Helper()

	return &internal.AuthenticatedClient{Client: buildDummyClient(t, id)}
}

// buildDummyClient builds a dummy client for testing.
func buildDummyClient(t *testing.T, id string) *internal.Client {
	t.Helper()

	redirectURI, err := url.Parse("http://localhost:3000/oauth2/callback")
	if err != nil {
		t.Fatal(err)
	}

	client, err := internal.NewClient(id, internal.ClientTypeConfidential, internal.NewHashedPassword("dummyPassword"), []url.URL{*redirectURI})
	if err != nil {
		t.Fatal(err)
	}
	return client
}
