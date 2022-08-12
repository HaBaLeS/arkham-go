package main

import (
	"github.com/HaBaLeS/arkham-go/command"
	"github.com/HaBaLeS/arkham-go/engine"
	"github.com/HaBaLeS/arkham-go/runtime"
	"log"
)

var guiChan chan command.GuiCommand
var running = true

func main() {

	//Load all Data
	guiChan = make(chan command.GuiCommand, 100)
	runtime.Init(guiChan)

	//Register listener for Gui Stuff
	go startChannelListener()

	//Start the Main loop
	gameEngine := engine.BuildArkhamGame()
	gameEngine.Start()
}

func startChannelListener() {
	for running {
		cmd := <-guiChan
		log.Printf("Received Command: %v", cmd)
	}
}
