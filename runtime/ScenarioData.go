package runtime

import "github.com/HaBaLeS/arkham-go/card"

type ScenarioData struct {
	StartLocation *card.Location
	CurrentAct    *card.Act
	CurrentAgenda *card.Agenda

	locations   []*card.Location
	agendaCards []*card.Agenda
	actCards    []*card.Act

	Player        []*PlayerDeck
	CurrentPlayer *PlayerDeck
}

func GetFirstScenarioData(db *CardDB) *ScenarioData {
	scn := &ScenarioData{}
	crd := db.GetCard("01111")
	scn.StartLocation = card.AcAsLocation(crd)

	c1105 := db.GetCard("01105")
	//c1106 := db.GetCard("01106")
	//c1107 := db.GetCard("01107")
	scn.CurrentAgenda = card.AcAsAgenda(c1105)

	c1108 := db.GetCard("01108")
	//c1109 := db.GetCard("01109")
	//c1110 := db.GetCard("01110")
	scn.CurrentAct = card.AcAsAct(c1108)

	scn.locations = make([]*card.Location, 0)
	scn.agendaCards = make([]*card.Agenda, 0)
	scn.actCards = make([]*card.Act, 0)
	scn.Player = make([]*PlayerDeck, 0)

	scn.locations = append(scn.locations, scn.StartLocation)

	return scn
}

func (s *ScenarioData) AddPlayer(p *PlayerDeck) {
	s.Player = append(s.Player, p)
}

func (s *ScenarioData) GetActiveLocations() []*card.Location {
	return s.locations
}
