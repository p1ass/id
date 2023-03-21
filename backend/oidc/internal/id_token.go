package internal

import (
	"fmt"
	"time"

	"github.com/Songmu/flextime"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/lestrrat-go/jwx/v2/jwt/openid"
)

type SignedIDToken struct {
	serialized []byte
}

const (
	idTokenExpiration = 10 * time.Minute

	PrivateKeySrc = `{
      "d": "vt4a05RfGvWSeT3geaH2FZAvio8Fs4x_nvyATUxvpPeAHurb7Du_YHMeZu_lXnzS872RZmM9K5jP_jpRExzgpiZgWqIqK6F1TcmKLzAqcrgExjPNg5VcqEXJZFlp_QM5nMCY1RZ0P4tv6WyHl-XXj10TOfyFiwEpOAw-Qi8jkiEHc8qcR1Jyb47IIWbRclkGhoVG_9CyGrmOtetoUH94AOVaPjFJh6vzPFU2trHvIwG_uTCF3f-3hb6uVbpvE5Rcp-xfd6JxdmsiMhnm-Xk7OqHrAY9g5HZr4UsswS5_rTiVpXvATpimWG1IjWTRnpsJgMt2ndjSMBIb-b4KZQl1oQ",
      "dp": "q7u3SubblCagQHMBJYUPmIF_ClHSbpEj_SOGBtxEZVjTYO7wwcx57UFo25eHQaOZ-tCTULOhRUzXkGoJF8yP-fiOR6E5HNBMbTty-pAj6x6qJ6nGnX7jDJb2D0v-kQdMz6WmPOaXLtdaMDOFpQ5eLVMhpletZwvuF11Lif6f7yk",
      "dq": "UgVaYxJO5IvkA508xi0YGA5rNrCAzfnPIIM82Wxu9F9rT__T6Re56oQ9Hf95iX1fGXDImqOHRfcSdpqmMzpJPrWvTFcoMuEqzL1wRoq7v2hScF78ZtDtSRNQcv2-4XlUevZhOyI-X-qn-BXrkwqFsNa46B_GBaIilAJ7tRSRT3s",
      "e": "AQAB",
      "kty": "RSA",
      "n": "7Svbi9qehbxTSSYSbDaONuraOLR8TxvE_pfN9MgUMt0lKbB4LRSlcD33zB04scbeK_B0Mz1v_cbg1_Kjj6dAbMJ0v7f0L1AuDKI73F1vKB08p4SJdQUi8hnPm6zu9Jlzm6wxJfGUuI9cFMHO4aa04OPewaqDJdgBcU4NBR9pp4QBkRl3OaDWWO6dc_KomLKH7_xgZRscptGxd5qgCmZLPnJt_JUEkBDpGaLJ0J1WYFJX9PXNudjt1rlE374Z-AavLbr3_TVptyyAVHCsrRCf13JOCjgtxx_P4aNrrjY5w9uyaqF29qepugDHT-dvjGgyF3wGhYkwkYmFarbOqaL1aw",
      "p": "_bMkUVq7x8LbeQKrYKcNk16VNuaEBX26CgbFFDFkTsOLvp8JBg95n_ZXOOTyWQiLRPqX_iW3YH2MsxrTFi-VNnuqNERZ2RpZjn35DWNvgCOTweiA-sw_MkGIzfO-FTY8mLQfPryYKrGc2O_zfiOZmusrBzQc0oqUxdugpJbanHk",
      "q": "71JaEXI22svFgJEp1WrVCHOQCxM-ookrpRMbPHqBR67v6VT3s6HkAqs1pqCWAKYLxYGrlVISmFIRLwzggdijaAkvvpssEx4qZdQKNkOFtMEYYyESs3sPcZVQTJIjQUyC2NYjmbCTjeZPBCLTtXA9ZxKWI1E7XO0GiE3GBvSkIAM",
      "qi": "dPcbrKQoaIli_qbJ--jdr9LiGAiqRMO-6mPFyfkaIZSkT4FdMcrqlrwWiNbQ-ZIx8SnR9nZiO7_n99TECPQ7xAv_6UXjTfu47RCgLyd6t-6h8PRrA4y9XJzXDvLef8yopk4k2ls8T85x7SOWrrIjuYtUdq0ETm80JCaGHF3SNP0",
      "use": "sig",
      "alg": "RS256",
      "kid": "local_first_kid"
}`
)

func NewSignedIDToken(sub string, clientID string) (*SignedIDToken, error) {
	// iat resolution is seconds
	now := flextime.Now().UTC().Truncate(time.Second)

	idToken, err := openid.NewBuilder().
		// TODO: 固定値をやめる
		Issuer("https://api.dev.id.p1ass.com").
		Subject(sub).
		Audience([]string{clientID}).
		IssuedAt(now).
		Expiration(now.Add(idTokenExpiration)).
		// Claim("auth_time","TODO").
		// Claim("nonce","TODO").
		// Claim("acr","TODO").
		// Claim("amr","TODO").
		// Claim("azp","TODO").
		// Claim("at_hash","TODO").
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to create id token: %w", err)
	}

	key, err := jwk.ParseKey([]byte(PrivateKeySrc))
	if err != nil {
		return nil, fmt.Errorf("failed to parse jwk: %w", err)
	}

	serialized, err := jwt.Sign(idToken, jwt.WithKey(jwa.RS256, key))
	if err != nil {
		return nil, fmt.Errorf("failed to sign jwt: %w", err)
	}

	return &SignedIDToken{serialized: serialized}, nil
}

// Token returns encoded id token.
func (t *SignedIDToken) Token() string {
	return string(t.serialized)
}
