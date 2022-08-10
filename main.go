package main

import (
	"fmt"
	arkham_game "github.com/HaBaLeS/arkham-go/modules/arkham-game"
	"github.com/HaBaLeS/arkham-go/runtime"
	"log"
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

	session := runtime.PlaySession{}

	d1, err := runtime.LoadDeckFromFile("data/deck1.txt", app.db)
	if err != nil {
		panic(err)
	}
	d2, err := runtime.LoadDeckFromFile("data/deck2.txt", app.db)
	if err != nil {
		panic(err)
	}

	session.AddPlayer("falko", d1)
	session.AddPlayer("zwerg", d2)

	session.Init(nil)

	crd := app.db.FindCardByName("Dr. Milan Christopher")

	if crd == nil {
		panic("Card not found")
	}
	log.Printf("Found Card: %s", crd.CardCode())

	scn := runtime.GetFirstScenarioData(app.db)

	arkham := arkham_game.BuildArkhamGame(scn)
	arkham.Start()

}
