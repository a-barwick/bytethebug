package main

import (
	"bytethebug/internal/client"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

const (
	testdata = "testdata"
)

func main() {
	homePage := client.BuildHomePage()

	http.Handle("/", templ.Handler(homePage))

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
