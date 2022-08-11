package runtime

import (
	"bufio"
	"fmt"
	"github.com/HaBaLeS/arkham-go/card"
	"os"
	"strconv"
	"strings"
)

type PlayerDeck struct {
	Title        string
	Investigator *card.Investigator
	Cards        []card.ArkhamCard
	Hand         string
	GraveYard    string
	Reserve      string
	HandLimit    int
}

//shuffle

//draw

//addCards

//init

//build(description)

func LoadPlayerDeckFromFile(file string, db *CardDB) (*PlayerDeck, error) {

	retVal := new(PlayerDeck)

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		} else if strings.HasPrefix(line, "Packs:") {
			continue //ignore which packs are needed
		} else if line == "Assets" || line == "Skills" || line == "Events" || line == "Treacheries" {
			continue
		}

		if retVal.Title == "" {
			retVal.Title = line
		} else if retVal.Investigator == nil {
			crd := db.FindCardByName(line)
			if crd == nil {
				return nil, fmt.Errorf("could not find card for '%s'", line)
			}
			retVal.Investigator = card.AcAsInvestigator(crd)
		} else {

			line = strings.Split(line, "(")[0]
			line = strings.TrimSpace(line)
			amount, name, _ := strings.Cut(line, " ")
			name = strings.Split(name, ":")[0] //some cards
			crd := db.FindCardByName(name)
			if crd == nil {
				return nil, fmt.Errorf("could not find card for '%s'", name)
			}
			amount = strings.Trim(amount, "x")
			cnt, err := strconv.Atoi(amount)
			if err != nil {
				return nil, err
			}
			for i := 0; i < cnt; i++ {
				retVal.Cards = append(retVal.Cards, crd)
			}

		}
	}
	fmt.Printf("Created Deck '%s' for %s with %d cards\n", retVal.Title, retVal.Investigator.Name, len(retVal.Cards))
	return retVal, nil
}
