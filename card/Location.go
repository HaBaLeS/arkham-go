package card

import (
	"fmt"
	"log"
)

type Location struct {
	Common
	Clues int `json:"clues"`

	//-- Runtime info --//
	ActiveClueTokens int
}

func (l *Location) ActivateLocation(inv ...*Investigator) {
	log.Printf("Aktivate Location:(%s) %s for %v", l.CCode, l.Name, inv)
	l.PlaceClueTokensOnCard(l.Clues)
	l.Flipped = false
}

func (l *Location) PlaceClueTokensOnCard(clues int) {
	l.ActiveClueTokens = clues
}

func (l *Location) RemoveClueTokensOnCard(clues int) {
	l.ActiveClueTokens--
}

func AcAsLocation(ac ArkhamCard) *Location {
	c, ok := ac.(*Location)
	if !ok {
		panic(fmt.Errorf("could not convert %s to Location", ac.CardCode()))
	}
	return c
}
