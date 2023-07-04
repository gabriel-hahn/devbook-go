package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gabriel-hahn/devbook/config"
	"github.com/gabriel-hahn/devbook/router"
)

func main() {
	config.Initialize()

	r := router.Initialize()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
