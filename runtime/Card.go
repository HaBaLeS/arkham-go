package runtime

import "arkham-go/card"

type Card struct {
	Card   card.ArkhamCard
	Script *Script
}

func (crd *Card) Activate(ps *PlaySession) {
	crd.Script.CallEventIfExists(ps, NewGameEvent("Activate", map[string]interface{}{
		// Umm
	}))
}

func (crd *Card) HandleEvent(ps *PlaySession, event GameEvent) {
	//TODO implement me
	panic("implement me")
}

func (crd *Card) FetchCommands(ps *PlaySession) []GameCommand {
	//TODO implement me
	panic("implement me")
}
