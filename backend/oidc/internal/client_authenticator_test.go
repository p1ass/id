package internal_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/p1ass/id/backend/oidc/internal"
)

func TestBasicClientAuthenticator_Authenticate(t *testing.T) {
	t.Parallel()

	type savedClient struct {
		ID             string
		hashedPassword *internal.HashedPassword
	}
	type args struct {
		header http.Header
	}

	secret := internal.NewHashedPassword("verySecureSecret1")
	tests := []struct {
		name        string
		savedClient savedClient
		args        args
		want        *internal.AuthenticatedClient
		wantErr     error
	}{
		{
			name: "when client id and secret match, return authenticated client",
			savedClient: savedClient{
				ID:             "clientID1",
				hashedPassword: secret,
			},
			args: args{
				header: basicAuthHeader(t, "clientID1", "verySecureSecret1"),
			},
			want: &internal.AuthenticatedClient{
				Client: mustClient(internal.NewClient("clientID1", internal.ClientTypeConfidential, secret, nil)),
			},
			wantErr: nil,
		},
		{
			name: "when client id not match, return error",
			savedClient: savedClient{
				ID:             "unmatchedClientID",
				hashedPassword: secret,
			},
			args: args{
				header: basicAuthHeader(t, "clientID1", "verySecureSecret1"),
			},
			want:    nil,
			wantErr: internal.ErrClientNotAuthenticated,
		},
		{
			name: "when client secret not match, return error",
			savedClient: savedClient{
				ID:             "clientID1",
				hashedPassword: secret,
			},
			args: args{
				header: basicAuthHeader(t, "clientID1", "unmatchedSecret"),
			},
			want:    nil,
			wantErr: internal.ErrClientNotAuthenticated,
		},
		{
			name: "when basic auth header not found, return error",
			savedClient: savedClient{
				ID:             "clientID1",
				hashedPassword: secret,
			},
			args: args{
				header: http.Header{},
			},
			want:    nil,
			wantErr: internal.ErrClientNotAuthenticated,
		},
		{
			name: "when basic auth header is invalid, return error",
			savedClient: savedClient{
				ID:             "clientID1",
				hashedPassword: secret,
			},
			args: args{
				header: map[string][]string{"Authorization": {"Basic invalidHeaderValue"}},
			},
			want:    nil,
			wantErr: internal.ErrClientNotAuthenticated,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c, err := internal.NewClient(
				tt.savedClient.ID, internal.ClientTypeConfidential, tt.savedClient.hashedPassword, nil)
			if err != nil {
				t.Fatal(err)
			}

			datastore := internal.NewInMemoryClientDatastore()
			if err := datastore.SaveClient(c); err != nil {
				t.Fatal(err)
			}

			authenticator := internal.NewBasicClientAuthenticator(datastore)

			got, err := authenticator.Authenticate(context.Background(), tt.args.header)
			if !errors.Is(err, tt.wantErr) {
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

func TestBasicClientAuthenticator_Authenticate_ClientType(t *testing.T) {
	t.Parallel()

	type savedClient struct {
		ID         string
		clientType internal.ClientType
	}
	type args struct {
		header http.Header
	}

	secret := internal.NewHashedPassword("verySecureSecret1")
	tests := []struct {
		name        string
		savedClient savedClient
		args        args
		want        *internal.AuthenticatedClient
		wantErr     error
	}{
		{
			name: "when client id is confidential, return authenticated client",
			savedClient: savedClient{
				ID:         "clientID2",
				clientType: internal.ClientTypeConfidential,
			},
			args: args{
				header: basicAuthHeader(t, "clientID2", "verySecureSecret1"),
			},
			want: &internal.AuthenticatedClient{
				Client: mustClient(internal.NewClient("clientID2", internal.ClientTypeConfidential, secret, nil)),
			},
			wantErr: nil,
		},
		{
			name: "when client type is public, return ErrClientCredentialNotAllowed error",
			savedClient: savedClient{
				ID:         "clientID2",
				clientType: internal.ClientTypePublic,
			},
			args: args{
				header: basicAuthHeader(t, "clientID2", "verySecureSecret1"),
			},
			want:    nil,
			wantErr: internal.ErrClientCredentialNotAllowed,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c, err := internal.NewClient(
				tt.savedClient.ID, tt.savedClient.clientType, secret, nil)
			if err != nil {
				t.Fatal(err)
			}

			datastore := internal.NewInMemoryClientDatastore()
			if err := datastore.SaveClient(c); err != nil {
				t.Fatal(err)
			}

			authenticator := internal.NewBasicClientAuthenticator(datastore)

			got, err := authenticator.Authenticate(context.Background(), tt.args.header)
			if !errors.Is(err, tt.wantErr) {
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

func TestBasicClientAuthenticator_Authenticate_Datastore(t *testing.T) {
	t.Parallel()

	type savedClient struct {
		ID         string
		clientType internal.ClientType
	}
	type args struct {
		header http.Header
	}

	secret := internal.NewHashedPassword("verySecureSecret1")
	tests := []struct {
		name        string
		savedClient *savedClient
		args        args
		want        *internal.AuthenticatedClient
		wantErr     error
	}{
		{
			name: "when client is saved, return authenticated client",
			savedClient: &savedClient{
				ID:         "clientID3",
				clientType: internal.ClientTypeConfidential,
			},
			args: args{
				header: basicAuthHeader(t, "clientID3", "verySecureSecret1"),
			},
			want: &internal.AuthenticatedClient{
				Client: mustClient(internal.NewClient("clientID3", internal.ClientTypeConfidential, secret, nil)),
			},
			wantErr: nil,
		},
		{
			name:        "when client is not saved, return error",
			savedClient: nil,
			args: args{
				header: basicAuthHeader(t, "clientID3", "verySecureSecret1"),
			},
			want:    nil,
			wantErr: internal.ErrClientNotAuthenticated,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			datastore := internal.NewInMemoryClientDatastore()

			if tt.savedClient != nil {
				c, err := internal.NewClient(
					tt.savedClient.ID, tt.savedClient.clientType, secret, nil)
				if err != nil {
					t.Fatal(err)
				}
				if err := datastore.SaveClient(c); err != nil {
					t.Fatal(err)
				}
			}

			authenticator := internal.NewBasicClientAuthenticator(datastore)

			got, err := authenticator.Authenticate(context.Background(), tt.args.header)
			if !errors.Is(err, tt.wantErr) {
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
