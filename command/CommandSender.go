package command

import "log"

var (
	guiChan chan GuiCommand
)

func SetGuiChannel(gc chan GuiCommand) {
	guiChan = gc
}

func SendGuiCommand(cmd GuiCommand) {
	if guiChan == nil {
		panic(guiChan)
	}
	guiChan <- cmd
	log.Printf("Command Sent %s", cmd)
}
