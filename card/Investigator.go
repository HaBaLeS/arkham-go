package card

import "fmt"

type Investigator struct {
	Common
	//Name     string `json:"name"` Do not duplicate common here!! this will cause inconsistency
	RealName string `json:"real_name"`
	SubName  string `json:"subname"`
	Sanity   int    `json:"sanity"`
	Health   int    `json:"health"`

	//-- Runtime info --//
	CurrentLocation string
}

func AcAsInvestigator(ac ArkhamCard) *Investigator {
	c, ok := ac.(*Investigator)
	if !ok {
		panic(fmt.Errorf("could not convert %s to Investigator", ac.CardCode()))
	}
	return c
}
