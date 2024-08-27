package main

import (
	"log"
	"net/http"

	gene "github.com/iWorld-y/Gene/src"
)

func main() {
	engine := gene.Engine{}
	log.Fatal(http.ListenAndServe(":12345", &engine))
}
