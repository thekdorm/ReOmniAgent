package main

import (
	"fmt"
	"log"
	"net/http"
)

func server(codeChan chan<- string) {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})
	http.HandleFunc("/blizzard/wow/api", func(res http.ResponseWriter, req *http.Request) {
		codeChan <- req.URL.Query().Get("code")
	})

	log.Fatal(http.ListenAndServeTLS(":8000", "certs/localhost.crt", "certs/localhost.key", nil))
}
