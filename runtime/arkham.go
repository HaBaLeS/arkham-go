package runtime

import (
	"github.com/HaBaLeS/arkham-go/command"
	"log"
)

var (
	globals_initialized bool = false
	cardDbG             *CardDB
	scenarioG           *ScenarioData
	guiChan             chan command.GuiCommand
)

func Init(guiChan chan command.GuiCommand) {

	//--CARD Database --//
	cardDbG = NewCardDB()
	err := cardDbG.Init("../data/all_pretty.json")
	if err != nil {
		log.Panicf("could not init DB: %v", err)
	}
	log.Printf("CardDB:\n %s\n", cardDbG.Status())

	//-- SCENARIO --//
	scenarioG = GetFirstScenarioData(cardDbG)

	//-- Player --//
	d1, err := LoadPlayerDeckFromFile("../data/decks/deck1.txt", cardDbG)
	if err != nil {
		panic(err)
	}
	d2, err := LoadPlayerDeckFromFile("../data/decks/deck2.txt", cardDbG)
	if err != nil {
		panic(err)
	}
	scenarioG.AddPlayer(d1)
	scenarioG.AddPlayer(d2)
	scenarioG.CurrentPlayer = d1

	globals_initialized = true

	command.SetGuiChannel(guiChan)
	command.SendGuiCommand(&command.InfoCommand{
		Ctype: "Ready to Start!",
	})

}

func CardDBG() *CardDB {
	mustInit()
	return cardDbG
}

func ScenarioSession() *ScenarioData {
	mustInit()
	return scenarioG
}

func mustInit() {
	if !globals_initialized {
		log.Panicln("arkham_go Globals are not initialized")
	}
}
