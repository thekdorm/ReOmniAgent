package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

func ouathToken(codeChan chan string) *http.Client {
	// From: https://pkg.go.dev/golang.org/x/oauth2#example-Config

	id := Id
	secret := Secret

	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     id,
		ClientSecret: secret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://oauth.battle.net/authorize",
			TokenURL: "https://oauth.battle.net/token",
		},
		RedirectURL: "https://localhost:8000/blizzard/wow/api",
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	log.Printf("Visit the URL for the auth dialog: %v", url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	code := <-codeChan
	log.Println("Code: " + code)

	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Token: " + tok.AccessToken)

	return conf.Client(ctx, tok)
}

func server(codeChan chan<- string) {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})
	http.HandleFunc("/blizzard/wow/api", func(res http.ResponseWriter, req *http.Request) {
		codeChan <- req.URL.Query().Get("code")
	})

	log.Fatal(http.ListenAndServeTLS(":8000", "certs/localhost.crt", "certs/localhost.key", nil))
}

func main() {
	codeChan := make(chan string)

	go server(codeChan)
	ouathToken(codeChan)
}
