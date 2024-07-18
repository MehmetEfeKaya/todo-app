package main

import (
	"log"
	"net/http"
	"yeni/backend/pkg/router"
)

func main() {
	r := router.SetupRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
