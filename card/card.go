package card

import "fmt"

type CType struct {
	CardType string `json:"type_code"`
}

const (
	InvestigatorType string = "investigator"
	TreacheryType    string = "treachery"
)

type ArkhamCard interface {
	Base() *Common
	CardType() string
	CardCode() string
}

type Common struct {
	CCode     string `json:"code"`
	Pack      string `json:"pack_code"`
	CType     string `json:"type_code"`
	Image     string `json:"imagesrc"`
	BackImage string `json:"backimagesrc"`
	Name      string `json:"name"`

	Flipped bool
	Tapped  bool
}

func (c *Common) Base() *Common {
	return c
}

func (c *Common) CardType() string {
	return c.CType
}

func (c *Common) CardCode() string {
	return c.CCode
}

type Investigator struct {
	Common
	//Name     string `json:"name"` Do not duplicate common here!! this will cause inconsistency
	RealName string `json:"real_name"`
	SubName  string `json:"subname"`
	Sanity   int    `json:"sanity"`
	Health   int    `json:"health"`
}

type Treachery struct {
	Common
}

type Asset struct {
	Common
}

type Event struct {
	Common
}
type Enemy struct {
	Common
}
type Skill struct {
	Common
}
type Scenario struct {
	Common
}
type Agenda struct {
	Common
}
type Act struct {
	Common
}

type Story struct {
	Common
}

func AcAsInvestigator(ac ArkhamCard) *Investigator {
	c, ok := ac.(*Investigator)
	if !ok {
		panic(fmt.Errorf("could not convert %s to Investigator", ac.CardCode()))
	}
	return c
}

func AcAsAct(ac ArkhamCard) *Act {
	c, ok := ac.(*Act)
	if !ok {
		panic(fmt.Errorf("could not convert %s to Act", ac.CardCode()))
	}
	return c
}

func AcAsAgenda(ac ArkhamCard) *Agenda {
	c, ok := ac.(*Agenda)
	if !ok {
		panic(fmt.Errorf("could not convert %s to Agenda", ac.CardCode()))
	}
	return c
}
