package main

import (
	"fmt"
	arkham_game "github.com/HaBaLeS/arkham-go/modules/arkham-game"
	"github.com/HaBaLeS/arkham-go/runtime"
)

type App struct {
	db *runtime.CardDB
}

func main() {
	fmt.Printf("go to hell\n")

	app := App{}
	app.db = runtime.NewCardDB()

	err := app.db.Init("all_pretty.json")
	if err != nil {
		panic(err)
	}

	fmt.Printf("CardDB:\n %s\n", app.db.Status())

	d1, err := runtime.LoadPlayerDeckFromFile("../data/decks/deck1.txt", app.db)
	if err != nil {
		panic(err)
	}
	d2, err := runtime.LoadPlayerDeckFromFile("../data/decks/deck2.txt", app.db)
	if err != nil {
		panic(err)
	}

	scn := runtime.GetFirstScenarioData(app.db)

	arkham := arkham_game.BuildArkhamGame(scn, nil)
	arkham.AddPlayer(d1)
	arkham.AddPlayer(d2)

	arkham.Start()

}
