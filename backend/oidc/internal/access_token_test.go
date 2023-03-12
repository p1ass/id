package internal_test

import (
	"github.com/Songmu/flextime"
	"github.com/p1ass/id/backend/oidc/internal"
	"testing"
	"time"
)

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
		t.Run(tt.name, func(t1 *testing.T) {
			flextime.Fix(createdTime)

			at, err := internal.NewAccessToken("dummy_sub", &internal.Client{}, []internal.Scope{internal.ScopeOpenId})
			if err != nil {
				t.Fatal(err)
			}

			flextime.Fix(tt.now)

			if got := at.ExpiresInSec(); got != tt.want {
				t1.Errorf("ExpiresInSec() = %v, want %v", got, tt.want)
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
		t.Run(tt.name, func(t1 *testing.T) {
			flextime.Fix(createdTime)

			at, err := internal.NewAccessToken("dummy_sub", &internal.Client{}, []internal.Scope{internal.ScopeOpenId})
			if err != nil {
				t.Fatal(err)
			}

			flextime.Fix(tt.now)

			if got := at.Expired(); got != tt.want {
				t1.Errorf("Expired() = %v, want %v", got, tt.want)
			}
		})
	}
}
