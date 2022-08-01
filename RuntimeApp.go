package main

import "arkham-go/runtime"

type RuntimeApp struct {
	CardDb      *runtime.CardDB
	PlaySession *runtime.PlaySession
}

func NewRuntime(app *App) *RuntimeApp {
	cardDB := runtime.NewCardDB()
	return &RuntimeApp{
		CardDb:      cardDB,
		PlaySession: runtime.NewGame(cardDB),
	}
}

func (rt *RuntimeApp) Init(app *App) {
	err := rt.CardDb.Init("all_pretty.json")
	if err != nil {
		panic(err)
	}
}
