package main

import (
	"log"
	"net/http"

	"github.com/bikkusah/urlShortening/controller"
	"github.com/bikkusah/urlShortening/database"
)

func main() {
	database.ConnectDb()

	http.HandleFunc("/short", controller.ShortTheUrl)
	http.HandleFunc("/url/", controller.RedirectURL)

	log.Println("Starting server on :8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("could not start server: %s", err.Error())
	}
}
