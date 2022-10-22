package internal

// Client represents OAuth 2.0 client.
type Client struct {
	// ID is a unique string  and is exposed to public.
	ID string
	// hashedPassword is used for HTTP Basic Authentication Scheme [RFC2617].
	// [RFC2617]: https://www.rfc-editor.org/rfc/rfc2617.html
	hashedPassword HashedPassword
}
