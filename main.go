package main

import (
	"fmt"
)

func main() {
	fmt.Printf("go to hell\n")

	app := NewApp()
	app.InitApp()

	fmt.Printf("CardDB:\n %s\n", app.Runtime.CardDb.Status())

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
	crd := app.Runtime.CardDb.FindCardByName("Dr. Milan Christopher")
	if crd == nil {
		panic("Card not found")
	}
	fmt.Printf("Found Card: %s\n", crd.CardCode())

	println("Running the Game")
	app.Runtime.PlaySession.Run()
	//app.Web.Gin.Run("[::1]:8080")
}
