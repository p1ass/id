package main

import (
	"fmt"

	"github.com/p1ass/id/backend/gen/oidc/v1/oidcv1connect"

	"golang.org/x/oauth2"
)

type foo struct {
}

var _ oidcv1connect.AuthenticationServiceHandler = &foo{}

func main() {
	fmt.Println(oauth2.Config{})

}
