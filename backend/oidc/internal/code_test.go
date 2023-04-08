package internal_test

import (
	"net/url"
	"testing"
	"time"

	"github.com/Songmu/flextime"
	"github.com/p1ass/id/backend/oidc/internal"
)

func TestNewAuthorizationCode_NewAuthorizationCode_GenerateDifferentCode(t *testing.T) {
	t.Parallel()
	client, err := internal.NewClient("clientID",
		internal.ClientTypeConfidential,
		internal.NewHashedPassword("verySecurePassword"),
		[]url.URL{mustURLParse("https://web.test/callback")})
	if err != nil {
		t.Fatal(err)
	}
	code1 := internal.NewAuthorizationCode(client, mustURLParse("https://web.test/callback"))
	code2 := internal.NewAuthorizationCode(client, mustURLParse("https://web.test/callback"))

	if code1.Code == code2.Code {
		t.Errorf("Code is same: %s", code1.Code)
	}
}

func TestAuthorizationCode_Expired(t *testing.T) {
	t.Parallel()

	issued := flextime.Now().UTC()

	tests := []struct {
		name string
		now  time.Time
		want bool
	}{
		{
			name: "When now process 30 seconds, return false",
			now:  issued.Add(30 * time.Second),
			want: false,
		},
		{
			name: "When now process 31 seconds, return true",
			now:  issued.Add(31 * time.Second),
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			client, err := internal.NewClient("clientID",
				internal.ClientTypeConfidential,
				internal.NewHashedPassword("verySecurePassword"),
				[]url.URL{mustURLParse("https://web.test/callback")})
			if err != nil {
				t.Fatal(err)
			}

			flextime.Fix(issued)
			c := internal.NewAuthorizationCode(client, mustURLParse("https://web.test/callback"))
			flextime.Fix(tt.now)

			if got := c.Expired(); got != tt.want {
				t.Errorf("Expired() = %v, want %v", got, tt.want)
			}
		})
	}
}
