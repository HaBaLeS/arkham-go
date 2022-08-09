package arkham_game

import (
	"log"
	"time"
)

func CreateUpkeepPhase() *ArkhamPhase {
	return &ArkhamPhase{
		name:     "Upkeep",
		execfunc: DoUpKeep,
	}
}

func DoUpKeep() {
	log.Printf("\t Reactivate Investigator")
	log.Printf("\t Untap all Cards")
	log.Printf("\t\t Enemys on same locations as Investigator will engage with one of them")
	log.Printf("\t Pull a card")
	log.Printf("\t Get a Resource")
	time.Sleep(500 * time.Millisecond)
}
