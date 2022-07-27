package main

import (
	"arkham-go/engine"
	"arkham-go/runtime"
	"fmt"
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
	fmt.Printf("Found Card: %s\n", crd.CardCode())

	println("Running the Game")
	game := engine.NewGame()
	game.Run()
}
