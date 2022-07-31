package main

import (
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

	/*
		d1, err := runtime.LoadDeckFromFile("data/deck1.txt", app.db)
		if err != nil {
			panic(err)
		}
		d2, err := runtime.LoadDeckFromFile("data/deck2.txt", app.db)
		if err != nil {
			panic(err)
		}
	*/
	crd := app.db.FindCardByName("Dr. Milan Christopher")
	if crd == nil {
		panic("Card not found")
	}
	fmt.Printf("Found Card: %s\n", crd.CardCode())

	println("Running the Game")
	game := runtime.NewGame()
	game.Run()
}
