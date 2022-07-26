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
type Location struct {
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
