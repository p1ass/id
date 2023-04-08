package internal_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/p1ass/id/backend/oidc/internal"
)

func TestAuthenticatedClient_IdenticalRedirectURI(t *testing.T) {
	t.Parallel()

	type fields struct {
		redirectURIs []url.URL
	}
	type args struct {
		redirectURI url.URL
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "when redirectURI is registered, return nil",
			fields: fields{
				redirectURIs: []url.URL{
					mustURLParse("https://web.test/callback"),
					mustURLParse("https://web.test/callback2"),
				},
			},
			args: args{
				redirectURI: mustURLParse("https://web.test/callback"),
			},
			wantErr: false,
		},
		{
			name: "when redirectURI is not registered, return err",
			fields: fields{
				redirectURIs: []url.URL{
					mustURLParse("https://web.test/callback"),
					mustURLParse("https://web.test/callback2"),
				},
			},
			args: args{
				redirectURI: mustURLParse("https://web.test/not-registered"),
			},
			wantErr: true,
		},
		{
			name: "when redirectURI matches partial, return err",
			fields: fields{
				redirectURIs: []url.URL{
					mustURLParse("https://web.test"),
					mustURLParse("https://web.test"),
				},
			},
			args: args{
				redirectURI: mustURLParse("https://web.test/partial-match"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c, err := internal.NewClient(
				"ID", internal.ClientTypeConfidential, &internal.HashedPassword{}, tt.fields.redirectURIs)
			if err != nil {
				t.Fatal(err)
			}
			if err := c.IdenticalRedirectURI(tt.args.redirectURI); (err != nil) != tt.wantErr {
				t.Errorf("IdenticalRedirectURI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func mustURLParse(rawURL string) url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return *u
}

func mustClient(c *internal.Client, err error) *internal.Client {
	if err != nil {
		panic(err)
	}
	return c
}

func basicAuthHeader(t *testing.T, username, password string) http.Header {
	t.Helper()
	req := http.Request{
		Header: map[string][]string{},
	}
	req.SetBasicAuth(username, password)
	return req.Header
}
