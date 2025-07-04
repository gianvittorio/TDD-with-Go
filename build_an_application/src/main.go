package main

import (
	"log"
	"net/http"

	"example.com/build_an_application/src/server"
	"example.com/build_an_application/src/store"
)

func main() {
	webserver := server.NewPlayerServer(store.NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":8080", webserver))
}
