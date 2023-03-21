package internal

import "net/url"

// NewClientFixture creates client fixture for local development.
func NewClientFixture() []*Client {
	redirectURIForTestSuite, err := url.Parse("https://localhost:8443/test/a/local/callback")
	if err != nil {
		panic(err)
	}
	redirectURIForLocalUI, err := url.Parse("http://local.p1ass.com:3000/oauth2/callback")
	if err != nil {
		panic(err)
	}
	redirectURIForProduction, err := url.Parse("https://id.p1ass.com/oauth2/callback")
	if err != nil {
		panic(err)
	}
	client, err := NewClient("dummy_client_id", NewHashedPassword("dummy_password"), []url.URL{
		*redirectURIForTestSuite,
		*redirectURIForLocalUI,
		*redirectURIForProduction,
	})
	if err != nil {
		panic(err)
	}

	return []*Client{client}
}
