package main

import (
	"log"
	"net/http"
	"os"
	"example.com/build_an_application"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	webserver := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":8080", webserver); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
