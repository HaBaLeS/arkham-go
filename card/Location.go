package card

import (
	"fmt"
	"github.com/HaBaLeS/arkham-go/command"
	"image"
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

	for i, v := range inv {
		command.SendGuiCommand(&command.PlayCardCommand{
			Ctype:      "play_card",
			CardToPlay: v.CCode,
			NextTo:     l.CCode,
			Scale:      0.8,
			X:          1920/2 - 330 + float64(i*495),
			Y:          float64(400),
			SubImage:   image.Rect(3, 7, 215, 175),
		})

		//Note down where the play is
		v.CurrentLocation = l.CCode
	}

	command.SendGuiCommand(&command.DisableCommand{
		What: "testButton",
	})

	command.SendGuiCommand(&command.EnableCommand{
		What: "investigator_gui",
	})
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
