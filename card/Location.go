package card

import (
	"fmt"
	"log"
)

type Location struct {
	Common
	Clues int `json:"clues"`
}

func (l *Location) ActivateLocation(inv ...*Investigator) {
	log.Printf("Aktivate Location:(%s) %s for %v", l.CCode, l.Name, inv)
	l.Flipped = true
}

func AcAsLocation(ac ArkhamCard) *Location {
	c, ok := ac.(*Location)
	if !ok {
		panic(fmt.Errorf("could not convert %s to Location", ac.CardCode()))
	}
	return c
}
