package command

import "image"

type GuiCommand interface {
	//empty interface .. sent anything
}

type InfoCommand struct {
	Ctype string
}

type PlayCardCommand struct {
	Ctype      string
	CardToPlay string //card sprite must be loaded already somewhere!

	//fixme this is still unsolved where to put them properly ... probably hardcoded surroundings to a location for the befinning
	NextTo   string
	Scale    float64
	X        float64
	Y        float64
	SubImage image.Rectangle
}

type EnableCommand struct {
	What string
}

type DisableCommand struct {
	What string
}

type EngineCommand interface {
	//empty interface .. sent anything
}

type DoInvestigate struct {
	Investigator string
	Location     string
}
