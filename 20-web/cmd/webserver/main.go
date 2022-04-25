package main

import (
	"log"
	"net/http"
	"poker"
)

func main() {
	fsps, cleanUp, err := poker.FileSystemPlayerStoreFromFile()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanUp()

	server := poker.NewPlayersServer(fsps)
	// server := NewPlayersServer(NewInMemStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}
