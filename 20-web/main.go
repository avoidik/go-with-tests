package main

import (
	"log"
	"net/http"
	"os"
)

const dbFilename = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFilename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("unable to open %q due to error %v", dbFilename, err)
	}
	defer db.Close()

	fsps, err := NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}

	server := NewPlayersServer(fsps)
	// server := NewPlayersServer(NewInMemStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}
