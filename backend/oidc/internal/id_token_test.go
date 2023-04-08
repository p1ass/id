package internal_test

import (
	"testing"
	"time"

	"github.com/Songmu/flextime"
	"github.com/google/go-cmp/cmp"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/p1ass/id/backend/oidc/internal"
)

func TestNewSignedIDToken_IssShouldBeIDP1assCom(t *testing.T) {
	t.Parallel()

	idToken, err := internal.NewSignedIDToken("dummy_sub", buildDummyAuthenticatedClient(t, "dummy_client_id"))
	if err != nil {
		t.Errorf("NewSignedIDToken should not be error, but got error: %s", err)
		return
	}

	wantIssuer := "https://api.dev.id.p1ass.com"
	parsed := mustParseIDToken(t, idToken)
	if parsed.Issuer() != wantIssuer {
		t.Errorf("NewSignedIDToken issu should be %s, but got %s", wantIssuer, parsed.Issuer())
	}
}

func TestNewSignedIDToken_EqualSub(t *testing.T) {
	t.Parallel()

	idToken, err := internal.NewSignedIDToken("dummy_sub", buildDummyAuthenticatedClient(t, "dummy_client_id"))
	if err != nil {
		t.Errorf("NewSignedIDToken should not be error, but got error: %s", err)
		return
	}

	wantSub := "dummy_sub"
	parsed := mustParseIDToken(t, idToken)
	if parsed.Subject() != wantSub {
		t.Errorf("NewSignedIDToken sub should be %s, but got %s", wantSub, parsed.Subject())
	}
}

func TestNewSignedIDToken_EqualAudience(t *testing.T) {
	t.Parallel()

	idToken, err := internal.NewSignedIDToken("dummy_sub", buildDummyAuthenticatedClient(t, "dummy_client_id"))
	if err != nil {
		t.Errorf("NewSignedIDToken should not be error, but got error: %s", err)
		return
	}

	wantAudiences := []string{"dummy_client_id"}
	parsed := mustParseIDToken(t, idToken)
	if !cmp.Equal(parsed.Audience(), wantAudiences) {
		t.Errorf("NewSignedIDToken aud diff = %s", cmp.Diff(parsed.Audience(), wantAudiences))
	}
}

func TestNewSignedIDToken_IssuedAtShouldBeNowSecond(t *testing.T) {
	t.Parallel()

	// iat resolution is seconds
	now := flextime.Now().UTC().Truncate(time.Second)
	defer flextime.Fix(now)()

	idToken, err := internal.NewSignedIDToken("dummy_sub", buildDummyAuthenticatedClient(t, "dummy_client_id"))
	if err != nil {
		t.Errorf("NewSignedIDToken should not be error, but got error: %s", err)
		return
	}

	parsed := mustParseIDToken(t, idToken)
	if !parsed.IssuedAt().Equal(now) {
		t.Errorf("NewSignedIDToken iat should be %s, but got %s", now, parsed.IssuedAt())
	}
}

func TestNewSignedIDToken_ExpirationShouldBeAfter10Minutes(t *testing.T) {
	t.Parallel()

	now := flextime.Now().UTC()
	defer flextime.Fix(now)()

	idToken, err := internal.NewSignedIDToken("dummy_sub", buildDummyAuthenticatedClient(t, "dummy_client_id"))
	if err != nil {
		t.Errorf("NewSignedIDToken should not be error, but got error: %s", err)
		return
	}

	parsed := mustParseIDToken(t, idToken)
	wantExpiration := now.Add(10 * time.Minute).Truncate(time.Second)
	if !parsed.Expiration().Equal(wantExpiration) {
		t.Errorf("NewSignedIDToken exp should be %s, but got %s", now, parsed.Expiration())
	}
}

func mustParseIDToken(t *testing.T, idToken *internal.SignedIDToken) jwt.Token {
	t.Helper()

	parsed, err := jwt.Parse([]byte(idToken.Token()), jwt.WithKey(jwa.RS256, mustGetPublicKey(t)))
	if err != nil {
		t.Fatalf("jwt.Parse should not be error, but got error: %s", err)
	}
	return parsed
}

func mustGetPublicKey(t *testing.T) jwk.Key {
	t.Helper()

	key, err := jwk.ParseKey([]byte(internal.PrivateKeySrc))
	if err != nil {
		t.Fatalf("jwk.ParseKey failed: %s", err)
	}

	pubKey, err := jwk.PublicKeyOf(key)
	if err != nil {
		t.Fatalf("jwk.PublicKeyOf failed: %s", err)
	}
	return pubKey
}
