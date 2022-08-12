package command

import "image"

type GuiCommand interface {
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
