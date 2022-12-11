package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

const baseUrl string = "https://us.api.blizzard.com"
const proudmooreRealmId int = 5

var auctionHouseUrl string = fmt.Sprintf("/data/wow/connected-realm/%d/auctions", proudmooreRealmId)

const namespace string = "dynamic-us"
const locale string = "en_US"

var urlQueries string = fmt.Sprintf("?namespace=%s&locale=%s", namespace, locale)

func check(err error, message string) {
	if err != nil {
		formatted := fmt.Sprintf("%s\n%s", message, err)
		log.Fatal(formatted)
	}
}

func oauthToken(codeChan chan string) *http.Client {
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
	check(err, "Token couldn't be exchanged!")

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
	client := oauthToken(codeChan)
	client.Timeout = 60 * time.Second // Auction House results take a while to come back
	close(codeChan)

	reqUrl := baseUrl + auctionHouseUrl + urlQueries
	fmt.Println(reqUrl)
	rsp, err := client.Get(reqUrl)
	check(err, "No response received!")

	defer rsp.Body.Close()
	body, err := io.ReadAll(rsp.Body)
	check(err, "No response body could be read!")

	fmt.Println(string(body))
}
