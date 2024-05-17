package util

import (
	"context"
	"encoding/json"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

func ServiceAccount(credentialFile string) (*oauth2.Token, error) {
	b, err := os.ReadFile(credentialFile)
	if err != nil {
		return nil, err
	}
	var c = struct {
		Email      string `json:"client_email"`
		PrivateKey string `json:"private_key"`
	}{}

	err = json.Unmarshal(b, &c)
	if err != nil {
		return nil, err
	}

	config := &jwt.Config{
		Email:      c.Email,
		PrivateKey: []byte(c.PrivateKey),
		Scopes: []string{
			"https://www.googleapis.com/auth/firebase.remoteconfig",
		},
		TokenURL: google.JWTTokenURL,
	}
	token, err := config.TokenSource(context.Background()).Token()
	if err != nil {
		return nil, err
	}
	return token, nil
}
