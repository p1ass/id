package internal_test

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
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
				"ID", &internal.HashedPassword{}, tt.fields.redirectURIs)
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

func TestClient_Authenticate(t *testing.T) {
	t.Parallel()

	type fields struct {
		ID             string
		hashedPassword *internal.HashedPassword
	}
	type args struct {
		header http.Header
	}

	secret := internal.NewHashedPassword(internal.RawPassword("verySecureSecret1"))
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *internal.AuthenticatedClient
		wantErr bool
	}{
		{
			name: "when client id and secret match, return authenticated client",
			fields: fields{
				ID:             "clientID1",
				hashedPassword: secret,
			},
			args: args{
				header: basicAuthHeader(t, "clientID1", "verySecureSecret1"),
			},
			want: &internal.AuthenticatedClient{
				Client: mustClient(internal.NewClient("clientID1", secret, nil)),
			},
			wantErr: false,
		},
		{
			name: "when client id not match, return error",
			fields: fields{
				ID:             "unmatchedClientID",
				hashedPassword: secret,
			},
			args: args{
				header: basicAuthHeader(t, "clientID1", "verySecureSecret1"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "when client secret not match, return error",
			fields: fields{
				ID:             "clientID1",
				hashedPassword: secret,
			},
			args: args{
				header: basicAuthHeader(t, "clientID1", "unmatchedSecret"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "when basic auth header not found, return error",
			fields: fields{
				ID:             "clientID1",
				hashedPassword: secret,
			},
			args: args{
				header: http.Header{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "when basic auth header is invalid, return error",
			fields: fields{
				ID:             "clientID1",
				hashedPassword: secret,
			},
			args: args{
				header: map[string][]string{"Authorization": {"Basic invalidHeaderValue"}},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c, err := internal.NewClient(
				tt.fields.ID, tt.fields.hashedPassword, nil)
			if err != nil {
				t.Fatal(err)
			}

			got, err := c.Authenticate(context.Background(), tt.args.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opt := cmp.AllowUnexported(internal.Client{}, internal.HashedPassword{})
			if !cmp.Equal(got, tt.want, opt) {
				t.Errorf("Authenticate() diff = %v", cmp.Diff(got, tt.want, opt))
			}
		})
	}
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
