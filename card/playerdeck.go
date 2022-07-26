package card

type PlayerDeck struct {
	Hand         string
	GraveYard    string
	Reserve      string
	Investigator Investigator
	HandLimit    int
}

func (pd *PlayerDeck) validate() error {
	//is deck valid
	return nil
}
