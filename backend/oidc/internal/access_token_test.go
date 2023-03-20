package internal_test

import (
	"testing"
	"time"

	"github.com/Songmu/flextime"
	"github.com/p1ass/id/backend/oidc/internal"
	"github.com/p1ass/id/backend/pkg/ascii"
)

func TestNewAccessToken_TokenTypeMustBeBearer(t *testing.T) {
	t.Parallel()

	got, err := internal.NewAccessToken("dummy_sub", &internal.Client{}, []internal.Scope{internal.ScopeOpenID})
	if err != nil {
		t.Fatal(err)
	}

	if got.TokenType != internal.AccessTokenTypeBearer {
		t.Errorf("token type must be bearer, but got is %s", got.TokenType)
	}
}

func TestNewAccessToken_TokenShouldBeASCII(t *testing.T) {
	t.Parallel()

	got, err := internal.NewAccessToken("dummy_sub", &internal.Client{}, []internal.Scope{internal.ScopeOpenID})
	if err != nil {
		t.Fatal(err)
	}

	if !ascii.IsASCII(got.Token) {
		t.Errorf("token should be ascii, but got %s is not ascii", got.Token)
	}
}

func TestNewAccessToken_TokenShouldNotBeEmpty(t *testing.T) {
	t.Parallel()

	got, err := internal.NewAccessToken("dummy_sub", &internal.Client{}, []internal.Scope{internal.ScopeOpenID})
	if err != nil {
		t.Fatal(err)
	}

	if got.Token == "" {
		t.Errorf("token should not be empty, but got is empty")
	}
}

func TestAccessToken_ExpiresInSec(t *testing.T) {
	t.Parallel()

	createdTime := flextime.Now().UTC()

	tests := []struct {
		name string
		now  time.Time
		want uint32
	}{
		{
			name: "When expiry is 30 seconds later, return 30",
			now:  createdTime.Add(14 * time.Minute).Add(30 * time.Second),
			want: 30,
		},
		{
			name: "When token is expired, return 0",
			now:  createdTime.Add(15 * time.Minute).Add(1 * time.Second),
			want: 0,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			flextime.Fix(createdTime)

			at, err := internal.NewAccessToken("dummy_sub", &internal.Client{}, []internal.Scope{internal.ScopeOpenID})
			if err != nil {
				t.Fatal(err)
			}

			flextime.Fix(tt.now)

			if got := at.ExpiresInSec(); got != tt.want {
				t.Errorf("ExpiresInSec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessToken_Expired(t *testing.T) {
	t.Parallel()

	createdTime := flextime.Now().UTC()

	tests := []struct {
		name string
		now  time.Time
		want bool
	}{
		{
			name: "When expiry is 30 seconds later, return false",
			now:  createdTime.Add(14 * time.Minute).Add(30 * time.Second),
			want: false,
		},
		{
			name: "When token is expired, return true",
			now:  createdTime.Add(15 * time.Minute).Add(1 * time.Second),
			want: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			flextime.Fix(createdTime)

			at, err := internal.NewAccessToken("dummy_sub", &internal.Client{}, []internal.Scope{internal.ScopeOpenID})
			if err != nil {
				t.Fatal(err)
			}

			flextime.Fix(tt.now)

			if got := at.Expired(); got != tt.want {
				t.Errorf("Expired() = %v, want %v", got, tt.want)
			}
		})
	}
}
