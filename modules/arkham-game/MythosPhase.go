package arkham_game

import (
	"log"
	"time"
)

func CreateMythosPhase() *ArkhamPhase {
	return &ArkhamPhase{
		name:     "Mythos",
		execfunc: DoMythos,
	}
}

func DoMythos() {
	log.Printf("\t Place a Doom Marker")
	log.Printf("\t\t Advance Aganda Deck if necessary")
	time.Sleep(500 * time.Millisecond)
}
