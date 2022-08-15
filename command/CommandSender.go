package command

import "log"

var (
	guiChan    chan GuiCommand
	eningeChan chan EngineCommand
)

func SetGuiChannel(gc chan GuiCommand) {
	guiChan = gc
}
func SetEngineChannel(gc chan EngineCommand) {
	eningeChan = gc
}

func SendGuiCommand(cmd GuiCommand) {
	if guiChan == nil {
		panic("GuiChan not open")
	}
	guiChan <- cmd
	log.Printf("Gui Command Sent %s", cmd)
}

func SendEngineCommand(cmd EngineCommand) {
	if eningeChan == nil {
		panic("EningeChan not open")
	}
	eningeChan <- cmd
	log.Printf("Engine Command Sent %s", cmd)
}
