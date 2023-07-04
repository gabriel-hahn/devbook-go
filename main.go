package main

import (
	"log"
	"net/http"

	"github.com/gabriel-hahn/devbook/router"
)

func main() {
	r := router.Initialize()

	log.Fatal(http.ListenAndServe(":5001", r))
}
