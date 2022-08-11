package command

type GuiCommand interface {
	CommandType() string //fixme create type
}

type InfoCommand struct {
	Ctype string
}

func (i *InfoCommand) CommandType() string {
	return i.Ctype
}

type PlayCardCommand struct {
	Ctype      string
	CardToPlay string //card sprite must be loaded already somewhere!

	//fixme this is still unsolved wheere to put them properkly ... probably hardcoded surroundings to a location for the befinning
	NextTo string
}

func (i *PlayCardCommand) CommandType() string {
	return i.Ctype
}
