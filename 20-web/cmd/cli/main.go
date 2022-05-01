package main

import (
	"fmt"
	"log"
	"os"
	"poker"
)

const dbFilename = "game.db.json"

func main() {
	fmt.Println("play poker")
	fmt.Println("input `<name> wins` and press enter to record a win")

	fsps, cleanUp, err := poker.FileSystemPlayerStoreFromFile()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanUp()

	alerter := poker.BlindAlerterFunc(poker.StdOutAlerter)
	game := poker.NewCLI(fsps, os.Stdin, alerter)
	game.PlayPoker()
}
